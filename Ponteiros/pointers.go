package ponteiros

import "fmt"

type Carro struct {
	Marca string //Apontando para um lugar na memória
}

func (c *Carro) MudaMarca(marca string) {
	c.Marca = marca //Apontando para outro lugar na memória
	fmt.Println(c.Marca)
}

func main() {
	carro := Carro{Marca: "Fiat"}
	carro.MudaMarca("Ford")
	fmt.Println(carro.Marca)

}
