package service

import (
	"github.com/gocql/gocql"
	store "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/STORE"
)

func DeleteMobile(cassandraSession *gocql.Session,id int) error {
	err := store.DeleteByID(cassandraSession,id)
	if err!= nil {
        return err
    }
	return err
}