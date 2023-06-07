package account

import (
	"crmservice/entities"
	"crmservice/middleware"
	"crmservice/repositories"
	"time"
)

type ActorInterfaceUseCase interface {
	GetActorByUsername(username string) (entities.Actor, error)
	CreateActor(actor ActorParam) (entities.Actor, error)
	GetActorById(id uint) (entities.Actor, error)
	UpdateActor(actor ActorParam, id uint) (any, error)
	DeleteActor(id uint) (any, error)
}

type ActorStructUseCase struct {
	actorRepo repositories.ActorInterfaceRepo
}

func (au ActorStructUseCase) GetActorByUsername(username string) (entities.Actor, error) {
	var actor entities.Actor
	actor, err := au.actorRepo.GetActorByUsername(username)
	return actor, err
}

func (au ActorStructUseCase) CreateActor(actor ActorParam) (entities.Actor, error) {
	var newActor *entities.Actor
	newActor = &entities.Actor{
		Username:   actor.Username,
		Password:   middleware.CreateHashPassword(actor.Password),
		RoleId:     2,
		IsVerified: actor.IsVerified,
		IsActive:   actor.IsActive,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	_, err := au.actorRepo.CreateActor(newActor)
	if err != nil {
		return *newActor, err
	}
	return *newActor, nil
}

func (au ActorStructUseCase) GetActorById(id uint) (entities.Actor, error) {
	var actor entities.Actor
	actor, err := au.actorRepo.GetActorById(id)
	return actor, err
}

func (au ActorStructUseCase) UpdateActor(actor ActorParam, id uint) (any, error) {
	var editActor *entities.Actor
	editActor = &entities.Actor{
		ID:         id,
		Username:   actor.Username,
		Password:   middleware.CreateHashPassword(actor.Password),
		RoleId:     2,
		IsVerified: actor.IsVerified,
		IsActive:   actor.IsActive,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	_, err := au.actorRepo.UpdateActor(editActor)
	if err != nil {
		return *editActor, err
	}
	return *editActor, nil
}

func (au ActorStructUseCase) DeleteActor(id uint) (any, error) {
	_, err := au.actorRepo.DeleteActor(id)
	return nil, err
}
