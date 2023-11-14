package app

import (
	"github.com/jeffleon1/ionix/docs"
	httpHandler "github.com/jeffleon1/ionix/internal/handler/http"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *App) RegisterRoutes(h *httpHandler.Handler) {

	a.Router.POST("/login", h.SignIn)
	a.Router.POST("/signup", h.SignUp)
	a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName(docs.SwaggerInfo.InstanceName())))

	// Drugs Routes
	private := a.Router.Group("/private")
	private.Use(JWTMiddleware(a.config.Auth.Secret))
	private.POST("/drugs", h.CreateDrug)
	private.PUT("/drugs/:id", h.UpdateDrug)
	private.GET("/drugs", h.GetDrugs)
	private.DELETE("/drugs/:id", h.DeleteDrug)
	// Vaccination Routes
	private.POST("/vaccination", h.CreateVaccination)
	private.PUT("/vaccination/:id", h.UpdateVaccination)
	private.GET("/vaccination", h.GetVaccinations)
	private.DELETE("/vaccination/:id", h.DeleteVaccination)
}
