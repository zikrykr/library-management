package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/auth/internal/auth/port"
	"github.com/zikrykr/library-management/shared/constants"
	"github.com/zikrykr/library-management/shared/pkg"
)

type ProfileHandler struct {
	profileService port.IProfileService
}

func NewProfileHandler(service port.IProfileService) port.IProfileHandler {
	return ProfileHandler{
		profileService: service,
	}
}

func (h ProfileHandler) GetProfile(c *gin.Context) {
	email, exists := c.Get(constants.CONTEXT_CLAIM_USER_EMAIL)
	if !exists {
		pkg.ResponseError(c, http.StatusUnauthorized, errors.New("email not found in context"))
		return
	}

	// Convert email to string
	emailStr, ok := email.(string)
	if !ok {
		pkg.ResponseError(c, http.StatusUnauthorized, errors.New("email type assertion failed"))
		return
	}

	ctx := c.Request.Context()

	resp, err := h.profileService.GetProfile(ctx, emailStr)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Get profile successful",
		Data:    resp,
	})
}
