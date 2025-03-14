package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Users []User

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	r.Use(cors.New(config))

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Users{
			{
				Name: "Endra",
				Email: "endra@mail.com",
			},
			{
				Name: "Nanda",
				Email: "nanda@mail.com",
			},
		})
	})

	r.Run(":8888")
}
