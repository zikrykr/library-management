package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/auth/internal/auth/payload"
	"github.com/zikrykr/library-management/services/auth/internal/auth/port"
	"github.com/zikrykr/library-management/shared/constants"
	"github.com/zikrykr/library-management/shared/pkg"
)

type SignUpHandler struct {
	signUpService port.ISignUpService
}

func NewSignUpHandler(service port.ISignUpService) port.ISignUpHandler {
	return SignUpHandler{
		signUpService: service,
	}
}

func (h SignUpHandler) SignUp(c *gin.Context) {
	var data payload.SignUpReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	data.Role = constants.ROLE_USER

	if err := h.signUpService.SignUp(ctx, data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, pkg.HTTPResponse{
		Success: true,
		Message: "User successfully registered",
	})
}

func (h SignUpHandler) SignUpAdmin(c *gin.Context) {
	var data payload.SignUpReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	data.Role = constants.ROLE_ADMIN

	if err := h.signUpService.SignUp(ctx, data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, pkg.HTTPResponse{
		Success: true,
		Message: "User successfully registered",
	})
}
