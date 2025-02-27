package main

import (
	"estoque/internal/models"
	"estoque/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Sistema de Estoque")

	estoque := services.NewEstoque()
	items := []models.Item{
		{ID: 1, Name: "Fone", Quantity: 10, Price: 100},
		{ID: 2, Name: "Camiseta", Quantity: 1, Price: 55.99},
		{ID: 3, Name: "Mouse", Quantity: 2, Price: 12.99},
	}

	for _, item := range items {
		err := estoque.AddItem(item)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	for _, item := range estoque.ListItems() {
		fmt.Printf("\nID: %d | Item: %s | Quantidade: %d | Preço: %2.f", item.ID, item.Name, item.Quantity, item.Price)
	}
}
