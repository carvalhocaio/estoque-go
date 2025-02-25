package main

import (
	"estoque/internal/models"
	"fmt"
)

func main() {
	fmt.Println("Sistema de Estoque")

	item1 := models.Item{
		ID:       1,
		Name:     "Primeiro Produto",
		Quantity: 10,
		Price:    19.99,
	}

	item2 := models.Item{
		ID:       2,
		Name:     "Segundo Produto",
		Quantity: 2,
		Price:    9.99,
	}

	fmt.Println(item1.Info())
	fmt.Println(item2.Info())
}
