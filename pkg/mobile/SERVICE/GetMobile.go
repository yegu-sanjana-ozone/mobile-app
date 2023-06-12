package service

import (
	"github.com/gocql/gocql"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	store "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/STORE"
)


func GetMobile (cassandraSession *gocql.Session)  ([]Model.Mobile, error) {

	err ,mobile:= store.GetAllMobiles(cassandraSession)
	if err!= nil {
			return err,nil
		}
		return nil,mobile
}