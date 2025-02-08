package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/auth/internal/auth/payload"
	"github.com/zikrykr/library-management/services/auth/internal/auth/port"
	"github.com/zikrykr/library-management/shared/pkg"
)

type LoginHandler struct {
	loginService port.ILoginService
}

func NewLoginHandler(service port.ILoginService) port.ILoginHandler {
	return LoginHandler{
		loginService: service,
	}
}

func (h LoginHandler) Login(c *gin.Context) {
	var data payload.LoginReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	resp, err := h.loginService.Login(ctx, data)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Login successful",
		Data:    resp,
	})
}
