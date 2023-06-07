package account

import (
	"crmservice/dto"
	"crmservice/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ActorStructRequestHandler struct {
	ctr ActorStructController
}

func NewActorRequestHandler(dbCrud *gorm.DB) ActorStructRequestHandler {
	return ActorStructRequestHandler{
		ctr: ActorStructController{
			actorUseCase: ActorStructUseCase{
				actorRepo: repositories.NewActor(dbCrud),
			},
		},
	}
}
func (ah ActorStructRequestHandler) LoginActor(c *gin.Context) {
	request := LoginActorParam{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage("MSG LOGIN"))
		return
	}
	result, err := ah.ctr.LoginActor(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage("MSG LOGIN 2"))
		return
	}
	c.JSON(http.StatusOK, result)
}

func (ah ActorStructRequestHandler) CreateActor(c *gin.Context) {
	request := ActorParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := ah.ctr.CreateActor(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ah ActorStructRequestHandler) GetActorById(c *gin.Context) {
	request := ActorParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}
	res, err := ah.ctr.GetActorById(uint(actorId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ah ActorStructRequestHandler) UpdateActor(c *gin.Context) {
	request := ActorParam{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := ah.ctr.UpdateActor(request, uint(actorId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ah ActorStructRequestHandler) DeleteActor(c *gin.Context) {
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := ah.ctr.DeleteActor(uint(actorId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}
