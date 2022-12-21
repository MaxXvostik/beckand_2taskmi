package main

import (
	"exemple.com/beckand-2taskmi/database"
	"exemple.com/beckand-2taskmi/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	database.Init()

	router := gin.Default()
	router.GET("/hello", handlers.Hello)
	router.POST("/registration", handlers.Registracion)

	router.Run()

}
