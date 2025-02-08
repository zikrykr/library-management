package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/zikrykr/library-management/services/auth/config"
	"github.com/zikrykr/library-management/services/auth/config/db"
	"gorm.io/gorm"

	authHandler "github.com/zikrykr/library-management/services/auth/internal/auth/handler"
	authPorts "github.com/zikrykr/library-management/services/auth/internal/auth/port"
	authrepo "github.com/zikrykr/library-management/services/auth/internal/auth/repository"
	authService "github.com/zikrykr/library-management/services/auth/internal/auth/service"
)

type SetupData struct {
	ConfigData  config.Config
	InternalApp InternalAppStruct
}

type InternalAppStruct struct {
	Repositories initRepositoriesApp
	Services     initServicesApp
	Handler      InitHandlerApp
}

// Repositories
type initRepositoriesApp struct {
	dbInstance *gorm.DB
	AuthRepo   authPorts.IAuthRepo
}

// Services
type initServicesApp struct {
	SignUpService  authPorts.ISignUpService
	LoginService   authPorts.ILoginService
	ProfileService authPorts.IProfileService
}

// Handler
type InitHandlerApp struct {
	SignUpHandler  authPorts.ISignUpHandler
	LoginHandler   authPorts.ILoginHandler
	ProfileHandler authPorts.IProfileHandler
}

// CloseDB close connection to db
var CloseDB func() error

func InitSetup() SetupData {
	configData := config.GetConfig()

	//DB INIT
	dbConn, err := db.Init()
	if err != nil {
		logrus.Fatal("database error", err)
	}

	CloseDB = func() error {
		if err := dbConn.CloseConnection(); err != nil {
			return err
		}

		return nil
	}

	internalAppVar := initInternalApp(dbConn.GormDB)

	return SetupData{
		ConfigData:  configData,
		InternalApp: internalAppVar,
	}
}

func initInternalApp(gormDB *db.GormDB) InternalAppStruct {
	var internalAppVar InternalAppStruct

	initAppRepo(gormDB, &internalAppVar)
	initAppService(&internalAppVar)
	initAppHandler(&internalAppVar)

	return internalAppVar
}

func initAppRepo(gormDB *db.GormDB, initializeApp *InternalAppStruct) {
	// Get Gorm instance
	initializeApp.Repositories.dbInstance = gormDB.DB

	initializeApp.Repositories.AuthRepo = authrepo.NewRepository(gormDB)
}

func initAppService(initializeApp *InternalAppStruct) {
	initializeApp.Services.SignUpService = authService.NewSignUpService(initializeApp.Repositories.AuthRepo)
	initializeApp.Services.LoginService = authService.NewLoginService(initializeApp.Repositories.AuthRepo)
	initializeApp.Services.ProfileService = authService.NewProfileService(initializeApp.Repositories.AuthRepo)
}

func initAppHandler(initializeApp *InternalAppStruct) {
	initializeApp.Handler.SignUpHandler = authHandler.NewSignUpHandler(initializeApp.Services.SignUpService)
	initializeApp.Handler.LoginHandler = authHandler.NewLoginHandler(initializeApp.Services.LoginService)
	initializeApp.Handler.ProfileHandler = authHandler.NewProfileHandler(initializeApp.Services.ProfileService)
}
