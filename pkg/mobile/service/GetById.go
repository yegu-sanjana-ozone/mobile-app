package service

import (

	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/model"
)

func (s *service) GetByID (id int) Model.Mobile {
	mobile := s.store.GetByID(id)
	return mobile
}