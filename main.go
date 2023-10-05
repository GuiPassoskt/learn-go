package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Veiculo interface {
	buzinar()
}

type Carro struct {
	Fabricante string
	Modelo     string
	Ano        int
}

func (c Carro) buzinar() {
	fmt.Println(c, "buzinou")
}

type Moto struct {
	Fabricante string
	Ano        int
}

func (m Moto) buzinar() {
	fmt.Println(m, "buzinou")
}

type Pessoa struct {
	Nome    string
	Veiculo Veiculo
}

func main() {
	r := gin.Default()

	carro := Carro{
		Fabricante: "Chevrolet",
		Modelo:     "Blazer",
		Ano:        2005,
	}

	joao := Pessoa{
		Nome:    "Joao",
		Veiculo: carro,
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": joao,
		})
	})
	r.Run()
}
