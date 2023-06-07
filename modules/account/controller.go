package account

import (
	"crmservice/dto"
	"crmservice/middleware"
	"errors"
)

type ActorInterfaceController interface {
	LoginActor(req LoginActorParam) (any, error)
	CreateActor(req ActorParam) (any, error)
	GetActorById(id uint) (FindActor, error)
	UpdateActor(req ActorParam, id uint) (any, error)
	DeleteActor(id uint) (any, error)
}

type ActorStructController struct {
	actorUseCase ActorInterfaceUseCase
}

func (ac ActorStructController) LoginActor(req LoginActorParam) (any, error) {
	var res dto.ResponseMeta
	login, err := ac.actorUseCase.GetActorByUsername(req.Username)
	if err != nil {
		return res, errors.New("Declined")
	}
	cek := middleware.VerifyPassword(req.Password, login.Password)
	if cek != true {
		return res, errors.New("Declined")
	}
	if login.IsVerified != "true" {
		return res, errors.New("Declined")
	}
	if login.IsActive != "true" {
		return res, errors.New("Declined")
	}
	token := middleware.CreateToken(login.ID, login.Username)
	res = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Token",
		Message:      token,
		ResponseTime: "",
	}
	return res, nil
}

func (ac ActorStructController) CreateActor(req ActorParam) (any, error) {
	actor, err := ac.actorUseCase.CreateActor(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create account",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: ActorParam{
			Username:   actor.Username,
			Password:   actor.Password,
			RoleId:     actor.RoleId,
			IsVerified: actor.IsVerified,
			IsActive:   actor.IsActive,
		},
	}
	return res, nil
}

func (ac ActorStructController) GetActorById(id uint) (FindActor, error) {
	var res FindActor
	actor, err := ac.actorUseCase.GetActorById(id)
	if err != nil {
		return FindActor{}, err
	}
	res.Data = actor
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get customer",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}

func (ac ActorStructController) UpdateActor(req ActorParam, id uint) (any, error) {
	var res dto.ResponseMeta
	_, err := ac.actorUseCase.UpdateActor(req, id)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success update"
	res.MessageTitle = "update"
	return res, nil
}

func (ac ActorStructController) DeleteActor(id uint) (any, error) {
	var res dto.ResponseMeta
	_, err := ac.actorUseCase.DeleteActor(id)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success Delete"
	res.MessageTitle = "Delete"
	return res, nil
}
