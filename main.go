package main

import (
	"github.com/GuiPassoskt/learn-go.git/types"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	carro := types.Carro{
		Fabricante: "Chevrolet",
		Modelo:     "Blazer",
		Ano:        2005,
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": carro,
		})
	})
	r.Run()
}
