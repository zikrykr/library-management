package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	appSetup "github.com/zikrykr/library-management/services/category/cmd/setup"
	"github.com/zikrykr/library-management/services/category/config"
	categoryRoutes "github.com/zikrykr/library-management/services/category/internal/categories/routes"
	"github.com/zikrykr/library-management/shared/constants"
	"github.com/zikrykr/library-management/shared/middleware"
)

// BaseURL base url of api
const (
	BaseURL      = "/api/v1/categories"
	BaseURLAdmin = "/api/v1/admin/categories"
)

func StartServer(setupData appSetup.SetupData) {
	conf := config.GetConfig()
	appName := conf.App.Name
	if conf.App.Env == constants.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	// GIN Init
	router := gin.Default()
	router.UseRawPath = true

	router.GET("/app", func(c *gin.Context) {
		c.JSON(http.StatusOK, fmt.Sprintf("%s is running", appName))
	})

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.JwtAuthMiddleware(conf.App.JWTSecret))

	//Init Main APP and Route
	initRoute(router, setupData.InternalApp)

	router.Use(middleware.CheckAdminRole())

	//Init Admin Route
	initAdminRoute(router, setupData.InternalApp)

	port := config.GetConfig().Http.Port
	httpServer := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := httpServer.ListenAndServe(); err != nil {
			logrus.Error(fmt.Printf("listen: %s\n", err))
		}
	}()
	logrus.Info("webserver started")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit

	logrus.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logrus.Panic("Server Shutdown:", err)
	}

	_ = appSetup.CloseDB()

	logrus.Info("Server exiting")
}

func initAdminRoute(router *gin.Engine, internalAppStruct appSetup.InternalAppStruct) {
	r := router.Group(BaseURLAdmin)
	categoryRoutes.AdminRoutes.NewAdminRoutes(r, internalAppStruct.Handler.CategoryHandler)
}

func initRoute(router *gin.Engine, internalAppStruct appSetup.InternalAppStruct) {
	r := router.Group(BaseURL)
	categoryRoutes.Routes.NewRoutes(r, internalAppStruct.Handler.CategoryHandler)
}
