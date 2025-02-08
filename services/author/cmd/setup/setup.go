package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/zikrykr/library-management/services/author/config"
	"github.com/zikrykr/library-management/services/author/config/db"
	authorHandler "github.com/zikrykr/library-management/services/author/internal/authors/handler"
	"github.com/zikrykr/library-management/services/author/internal/authors/port"
	authorRepo "github.com/zikrykr/library-management/services/author/internal/authors/repository"
	authorService "github.com/zikrykr/library-management/services/author/internal/authors/service"
	"gorm.io/gorm"
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

	AuthorRepo port.IAuthorRepo
}

// Services
type initServicesApp struct {
	AuthorService port.IAuthorService
}

// Handler
type InitHandlerApp struct {
	AuthorHandler port.IAuthorHandler
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

	initializeApp.Repositories.AuthorRepo = authorRepo.NewRepository(gormDB)
}

func initAppService(initializeApp *InternalAppStruct) {
	initializeApp.Services.AuthorService = authorService.NewAuthorService(initializeApp.Repositories.AuthorRepo)
}

func initAppHandler(initializeApp *InternalAppStruct) {
	initializeApp.Handler.AuthorHandler = authorHandler.NewAuthorHandler(initializeApp.Services.AuthorService)
}
