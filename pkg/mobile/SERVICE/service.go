package service

import (
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	store "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/STORE"
)

type Service interface {
	CreateMobile(mobile Model.Mobile) error
	DeleteMobile( id int ) error
	GetByID(id int ) Model.Mobile
	GetMobile()  ([]Model.Mobile, error)
	UpdateByID(id int ,brand string) error

   }

type service struct {
	store store.Store
}

func New() Service{
	return &service{
		store : store.GetStore(),
	}
}