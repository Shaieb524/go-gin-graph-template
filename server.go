package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joud-cxu/hackernews/hhttp"

	mssql "github.com/joud-cxu/hackernews/internal/pkg/db/migrations/mssql"
	database "github.com/joud-cxu/hackernews/internal/pkg/db/migrations/mysql"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.InitDB()
	database.Migrate()
	mssql.ConnectoMSSQL()

	server := gin.Default()
	server.GET("/", hhttp.PlaygroundHandler())
	server.POST("/query", hhttp.GraphQlHandler())
	server.Run(defaultPort)
}
