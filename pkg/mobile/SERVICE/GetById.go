package service

import (
	"github.com/gocql/gocql"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	store "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/STORE"
)

func GetByID (id int , cassandraSession *gocql.Session) Model.Mobile {
	mobile := store.GetByID(cassandraSession,id)
	return mobile
}