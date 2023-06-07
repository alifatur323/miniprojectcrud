package account

import (
	"crmservice/dto"
	"crmservice/entities"
)

type LoginActorParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ActorParam struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RoleId     uint   `json:"role_id"`
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data ActorParam `json:"data"`
}

type FindActor struct {
	dto.ResponseMeta
	Data entities.Actor `json:"data"`
}
