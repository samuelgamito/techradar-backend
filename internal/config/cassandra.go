package config

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"log"
	"time"
)

func CreateConfig() *gocqlx.Session {

	cluster := gocql.NewCluster("localhost:9042")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "", Password: ""}
	cluster.Keyspace = "tech_radar"
	var lastError error
	for retryCount := 0; retryCount < 3; retryCount++ {
		session, err := gocqlx.WrapSession(cluster.CreateSession())
		log.Printf("Cassandra connection attempt: %d\n", retryCount+1)
		if err == nil {
			return &session
		}
		lastError = err
	}

	log.Fatal(lastError)
	return nil
}
