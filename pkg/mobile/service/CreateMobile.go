package service

import (

	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/model"

)

func (s *service) CreateMobile (mobile Model.Mobile) error {
 

	err := s.store.CreateMobile(mobile)
	if err != nil{
		return err
	}
return nil
}