package service

import (
	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/DATABASE"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	store "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/STORE"
)

func CreateMobile (mobile Model.Mobile) error {
 
	cassandraSession := db.Session
	err := store.CreateMobile(cassandraSession ,mobile)
	if err != nil{
		return err
	}
return nil
}