package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/zikrykr/library-management/services/book/config"
	"github.com/zikrykr/library-management/services/book/config/db"
	"github.com/zikrykr/library-management/services/book/internal/books/port"
	"gorm.io/gorm"

	bookHandler "github.com/zikrykr/library-management/services/book/internal/books/handler"
	bookRepo "github.com/zikrykr/library-management/services/book/internal/books/repository"
	bookService "github.com/zikrykr/library-management/services/book/internal/books/service"
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

	BookRepo port.IBookRepo
}

// Services
type initServicesApp struct {
	BookService port.IBookService
}

// Handler
type InitHandlerApp struct {
	BookHandler port.IBookHandler
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

	initializeApp.Repositories.BookRepo = bookRepo.NewRepository(gormDB)
}

func initAppService(initializeApp *InternalAppStruct) {
	initializeApp.Services.BookService = bookService.NewBookService(initializeApp.Repositories.BookRepo)
}

func initAppHandler(initializeApp *InternalAppStruct) {
	initializeApp.Handler.BookHandler = bookHandler.NewBookHandler(initializeApp.Services.BookService)
}
