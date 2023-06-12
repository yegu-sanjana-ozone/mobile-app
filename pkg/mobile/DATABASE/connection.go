package db

import (
	"log"

	"github.com/gocql/gocql"
)

type DBConnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var Session *gocql.Session

var connection DBConnection

func init() {
	connection.cluster = gocql.NewCluster("127.0.0.1:9042")
	connection.cluster.Consistency = gocql.One
	connection.cluster.Keyspace = "mobile_app"
	connection.cluster.Port = 9042
    connection.cluster.Authenticator = gocql.PasswordAuthenticator{Username: "test", Password: "testpwd"}
	var err error
	connection.session, err = connection.cluster.CreateSession()
	if err != nil {
		log.Printf("CreateSession: %v", err)
		return
	}
	Session = connection.session
} 