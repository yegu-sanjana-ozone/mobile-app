package service

import (

	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"

)


func (s *service) GetMobile()  ([]Model.Mobile, error){

	err ,mobile:= s.store.GetAllMobiles()
	if err!= nil {
			return err,nil
		}
		return nil,mobile
}