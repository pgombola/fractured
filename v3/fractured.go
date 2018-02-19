package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

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

	var config Config
	switch {
	case os.IsNotExist(err):
		log.Println("Config file missing; using defaults")
		config = Config{
			DataDir:  "/var/lib/data",
			Host:     "127.0.0.1",
			Port:     "26257",
			Database: "",
		}
	case err == nil:
		if err := json.Unmarshal(data, &config); err != nil {
			log.Fatal(err)
		}
	default:
		log.Println(err)
	}

	log.Println("Overriding defaults from environment")
	if os.Getenv("APP_DATADIR") != "" {
		config.DataDir = os.Getenv("APP_DATADIR")
	}
	if os.Getenv("APP_HOST") != "" {
		config.Host = os.Getenv("APP_HOST")
	}
	if os.Getenv("APP_PORT") != "" {
		config.Port = os.Getenv("APP_PORT")
	}
	if os.Getenv("APP_USERNAME") != "" {
		config.Username = os.Getenv("APP_USERNAME")
	}
	if os.Getenv("APP_PASSWORD") != "" {
		config.Password = os.Getenv("APP_PASSWORD")
	}
	if os.Getenv("APP_DATABASE") != "" {
		config.Database = os.Getenv("APP_DATABASE")
	}

	// Check for data dir
	_, err = os.Stat(config.DataDir)
	if os.IsNotExist(err) {
		log.Println("Creating data dir ", config.DataDir)
		err = os.MkdirAll(config.DataDir, 0755)
	}
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

	var dbErr error
	maxAttempts := 20
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		dbErr = db.Ping()
		if dbErr == nil {
			break
		}
		log.Println(dbErr)
		time.Sleep(time.Duration(attempts) * time.Second)
	}
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	log.Println("App started successfully")
}
