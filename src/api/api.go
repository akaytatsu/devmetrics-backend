package api

import (
	"log"

	"app/api/handlers"
	"app/config"
	infra_mongodb "app/infrastructure/mongodb"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupDatabase() *mongo.Database {
	conn := infra_mongodb.Connect()
	return conn
}

func setupRouter(conn *mongo.Database) *gin.Engine {
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	handlers.MountSamplesHandlers(r)
	handlers.MountUsersHandlers(r, conn)

	return r
}

func SetupRouters() *gin.Engine {
	conn := setupDatabase()
	return setupRouter(conn)
}

func StartWebServer() {
	config.ReadEnvironmentVars()

	r := SetupRouters()

	// Bind to a port and pass our router in
	log.Fatal(r.Run())
}
