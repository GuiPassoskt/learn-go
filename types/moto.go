package types

import "fmt"

type Moto struct {
	Fabricante string
	Ano        int
}

func (m Moto) buzinar() {
	fmt.Println(m, "buzinou")
}
