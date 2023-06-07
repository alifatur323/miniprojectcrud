package repositories

import (
	"crmservice/entities"
	"gorm.io/gorm"
)

type Actor struct {
	db *gorm.DB
}

func NewActor(dbCrud *gorm.DB) Actor {
	return Actor{db: dbCrud}
}

type ActorInterfaceRepo interface {
	GetActorByUsername(username string) (entities.Actor, error)
	CreateActor(actor *entities.Actor) (*entities.Actor, error)
	GetActorById(id uint) (entities.Actor, error)
	UpdateActor(actor *entities.Actor) (any, error)
	DeleteActor(id uint) (any, error)
}

func (repo Actor) GetActorByUsername(username string) (entities.Actor, error) {
	var actor entities.Actor
	repo.db.First(&actor, "username = ?", username)
	return actor, nil
}

func (repo Actor) CreateActor(actor *entities.Actor) (*entities.Actor, error) {
	err := repo.db.Model(&entities.Actor{}).Create(actor).Error
	return actor, err
}

func (repo Actor) GetActorById(id uint) (entities.Actor, error) {
	var actor entities.Actor
	repo.db.First(&actor, "id = ?", id)
	return actor, nil
}

func (repo Actor) UpdateActor(actor *entities.Actor) (any, error) {
	err := repo.db.Save(actor).Error
	return nil, err
}

func (repo Actor) DeleteActor(id uint) (any, error) {
	err := repo.db.Model(&entities.Actor{}).
		Where("id = ?", id).
		Delete(&entities.Actor{}).
		Error
	return nil, err
}
