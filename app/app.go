package app

import (
	"github.com/sirupsen/logrus"
	"netsepio-gateway-v1.1/internal/caching"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/internal/server"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

// Initialize the app
func Init() {
	logrus.Println("Initializing the app...")
	// test db connection
	database.GetDb()
	// Migrate the database
	database.Migrate()
	// Initialize Redis
	caching.InitRedis()
	logwrapper.Init()
	// Initialize the server
	server.Init()

	logrus.Println("App initialized successfully.")

}
