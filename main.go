package main

import (
	"github.com/3cognito/library/app/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Load()
}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	err := r.Run(":" + config.Configs.Port)
	if err != nil {
		panic(err)
	}
}
