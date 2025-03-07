package main

import (
	"estoque/internal/models"
	"estoque/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Sistema de Estoque")

	estoque := services.NewEstoque()
	itens := []models.Item{
		{ID: 1, Name: "Fone", Quantity: 10, Price: 100},
		{ID: 2, Name: "Camiseta", Quantity: 1, Price: 50},
		{ID: 3, Name: "Mouse", Quantity: 2, Price: 12.99},
	}

	for _, item := range itens {
		err := estoque.AddItem(item, "Monkey D. Luffy")
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	for _, item := range estoque.ListItems() {
		fmt.Printf(
			"\nID: %d | Item: %s | Quantidade: %d | Preço: %2.f",
			item.ID, item.Name, item.Quantity, item.Price,
		)
	}
	fmt.Println("\nValor total do estoque R$:", estoque.CalculateTotalCost())

	/* logs := estoque.ViewAuditLog()
	for _, log := range logs {
		fmt.Printf(
			"\n[%s] Ação: %s - Usuário: %s - Item ID: %d - Quantidade: %d - Motivo: %s",
			log.Timestamp.Format("01/02 15:04:05"), log.Action, log.User, log.ItemID, log.Quantity, log.Reason,
		)
	} */

	itemBySearch, err := services.FindBy(itens, func(item models.Item) bool {
		return item.Price > 40
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("itens encontrados: ", itemBySearch)
}
