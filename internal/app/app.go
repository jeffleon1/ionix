package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jeffleon1/ionix/config"
	"github.com/jeffleon1/ionix/internal/controller/drug"
	"github.com/jeffleon1/ionix/internal/controller/user"
	"github.com/jeffleon1/ionix/internal/controller/vaccination"
	"github.com/jeffleon1/ionix/internal/db"
	httpHandler "github.com/jeffleon1/ionix/internal/handler/http"
	"github.com/jeffleon1/ionix/internal/repository/posgrest"
	"github.com/jeffleon1/ionix/pkg/entities"
	"gorm.io/gorm"
)

type App struct {
	db     *gorm.DB
	config *config.Config
	Router *gin.Engine
}

func (a *App) Initialize(cfg *config.Config) {
	a.config = cfg
	a.db = db.Connect(cfg)
	if err := db.MigrateDB(a.db); err != nil {
		panic(err)
	}
	drugRepository := posgrest.New[entities.Drug](a.db)
	vaccinationRepository := posgrest.New[entities.Vaccination](a.db)
	userRepository := posgrest.NewUserRepository[entities.User](a.db)
	drugCtrl := drug.New(drugRepository)
	vaccinationCtrl := vaccination.New(vaccinationRepository)
	userCtrl := user.New(userRepository, cfg.Auth.Secret, cfg.Auth.TokenDuration)
	handler := httpHandler.New(drugCtrl, vaccinationCtrl, userCtrl)
	a.Router = gin.Default()
	a.Router.Use(gin.Recovery())
	a.RegisterRoutes(handler)
}

func (a *App) Run() {
	err := a.Router.Run(fmt.Sprintf(":%s", a.config.APP.PORT))
	if err != nil {
		panic(err)
	}
}
