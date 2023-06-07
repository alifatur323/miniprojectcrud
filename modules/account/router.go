package account

import (
	"crmservice/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterActor struct {
	ActorRequestHandler ActorStructRequestHandler
}

func NewRouter(dbCrud *gorm.DB) RouterActor {
	return RouterActor{ActorRequestHandler: NewActorRequestHandler(dbCrud)}
}

func (r RouterActor) Handle(router *gin.Engine) {
	basepath := "/account"
	account := router.Group(basepath)
	account.POST("/login", r.ActorRequestHandler.LoginActor)

	account.Use(middleware.AuthMiddleware)
	{
		account.POST("/", r.ActorRequestHandler.CreateActor)
		account.GET("/:id", r.ActorRequestHandler.GetActorById)
		account.PUT("/:id", r.ActorRequestHandler.UpdateActor)
		account.DELETE("/:id", r.ActorRequestHandler.DeleteActor)
	}
}
