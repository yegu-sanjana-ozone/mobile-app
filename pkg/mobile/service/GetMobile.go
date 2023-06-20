package service

import (

	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/model"

)


func (s *service) GetMobile()  ([]Model.Mobile, error){

	err ,mobile:= s.store.GetAllMobiles()
	if err!= nil {
			return nil,nil
		}
		return nil,mobile
}