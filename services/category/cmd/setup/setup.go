package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/zikrykr/library-management/services/category/config"
	"github.com/zikrykr/library-management/services/category/config/db"
	"github.com/zikrykr/library-management/services/category/internal/categories/port"
	"gorm.io/gorm"

	categoryHandler "github.com/zikrykr/library-management/services/category/internal/categories/handler"
	categoryRepo "github.com/zikrykr/library-management/services/category/internal/categories/repository"
	categoryService "github.com/zikrykr/library-management/services/category/internal/categories/service"
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

	CategoryRepo port.ICategoryRepo
}

// Services
type initServicesApp struct {
	CategoryService port.ICategoryService
}

// Handler
type InitHandlerApp struct {
	CategoryHandler port.ICategoryHandler
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

	initializeApp.Repositories.CategoryRepo = categoryRepo.NewRepository(gormDB)
}

func initAppService(initializeApp *InternalAppStruct) {
	initializeApp.Services.CategoryService = categoryService.NewCategoryService(initializeApp.Repositories.CategoryRepo)
}

func initAppHandler(initializeApp *InternalAppStruct) {
	initializeApp.Handler.CategoryHandler = categoryHandler.NewCategoryHandler(initializeApp.Services.CategoryService)
}
