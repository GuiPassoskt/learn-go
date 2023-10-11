package client

import "fmt"

type Carro struct {
	Fabricante string
	Modelo     string
	Ano        int
}

func (c Carro) buzinar() {
	fmt.Println(c, "buzinou")
}
