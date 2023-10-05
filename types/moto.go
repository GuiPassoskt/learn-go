package types

type Moto struct {
	Fabricante string
	Ano        int
}

func (m Moto) buzinar() {
	print(m, "buzinou")
}
