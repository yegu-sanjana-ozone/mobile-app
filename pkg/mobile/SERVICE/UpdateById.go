package service

import (
	"github.com/gocql/gocql"
	store "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/STORE"
)

func UpdateByID(id int , cassandraSession *gocql.Session,brand string) error {
	err := store.UpdateByID(cassandraSession,id,brand)
	if err!= nil {
        return err
    }
	return err
}