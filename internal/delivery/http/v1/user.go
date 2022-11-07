package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuravlev-pe/course-watch/internal/delivery/http/v1/auth"
	"github.com/zhuravlev-pe/course-watch/internal/delivery/http/v1/utils"
	"github.com/zhuravlev-pe/course-watch/internal/repository"
	"github.com/zhuravlev-pe/course-watch/internal/service"
	"net/http"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	courses := api.Group("/user", h.bearer.Authenticate)
	{
		courses.GET("/", h.getUserInfo)
		courses.PUT("/", h.updateUserInfo)
	}
}

// @Summary Retrieve current user data
// @Tags User
// @Description returns info on the currently logged-in user. User_id is extracted from the bearer token
// @ModuleID getUserInfo
// @Accept  json
// @Produce  json
// @Success 200 {object} service.GetUserInfoOutput
// @Failure 400,401,404,500 {object} utils.Response
// @Failure default {object} utils.Response
// @Router /user/ [get]
func (h *Handler) getUserInfo(ctx *gin.Context) {
	up, err := auth.GetAuthenticatedUser(ctx)
	if err != nil {
		err = fmt.Errorf("authentication middleware failure: %w", err)
		utils.ErrorResponseMessageOverride(ctx, http.StatusInternalServerError, err, "user data processing failure")
		return
	}

	result, err := h.services.Users.GetUserInfo(ctx.Request.Context(), up.UserId)

	if err != nil {
		// TODO: discriminate between validation errors, logic errors and internal server errors
		if err == repository.ErrNotFound {
			utils.ErrorResponse(ctx, http.StatusNotFound, err)
			return
		}
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// @Summary Modify current user data
// @Tags User
// @Description modifies user info for the currently logged-in user. User_id is extracted from the bearer token
// @ModuleID updateUserInfo
// @Accept  json
// @Produce  json
// @Param input body service.UpdateUserInfoInput true "user info"
// @Success 200
// @Failure 400,401,404,500 {object} utils.Response
// @Failure default {object} utils.Response
// @Router /user/ [put]
func (h *Handler) updateUserInfo(ctx *gin.Context) {
	up, err := auth.GetAuthenticatedUser(ctx)
	if err != nil {
		err = fmt.Errorf("authentication middleware failure: %w", err)
		utils.ErrorResponseMessageOverride(ctx, http.StatusInternalServerError, err, "user data processing failure")
		return
	}
	var input service.UpdateUserInfoInput
	if err = ctx.BindJSON(&input); err != nil {
		utils.ErrorResponseString(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.services.Users.UpdateUserInfo(ctx.Request.Context(), up.UserId, &input)

	if err != nil {
		// TODO: discriminate between validation errors, logic errors and internal server errors
		if err == repository.ErrNotFound {
			utils.ErrorResponse(ctx, http.StatusNotFound, err)
			return
		}
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
