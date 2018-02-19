package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
)

// Config holds configuration of the app
type Config struct {
	DataDir string

	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func main() {
	log.Println("starting app...")

	// Load config
	data, err := ioutil.ReadFile("/etc/config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	// Check for data dir
	_, err = os.Stat(config.DataDir)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to db
	hostPort := net.JoinHostPort(config.Host, config.Port)
	url := fmt.Sprintf("postgresql://%s@%s?sslmode=disable", config.Username, hostPort)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
