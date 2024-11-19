package main

import (
	"fmt"
	"strings"
)

type FacturationResponse struct {
	Discount      string
	Price         int
	DiscountPrice int
}

func facturation(price int) FacturationResponse {
	facturationResponse := FacturationResponse{Discount: "", Price: 0, DiscountPrice: 0}
	if price < 25000 {
		facturationResponse.Price = price
		facturationResponse.DiscountPrice = price
		facturationResponse.Discount = "0%"
		return facturationResponse
	}
	if price == 25000 || price <= 50000 {
		facturationResponse.Price = price
		facturationResponse.DiscountPrice = (price * 5) / 100
		facturationResponse.Discount = "5%"
		return facturationResponse
	}
	if price > 50000 && price <= 100000 {
		facturationResponse.Price = price
		facturationResponse.DiscountPrice = (price * 10) / 100
		facturationResponse.Discount = "10%"
		return facturationResponse
	}
	facturationResponse.Price = price
	facturationResponse.DiscountPrice = (price * 15) / 100
	facturationResponse.Discount = "15%"
	return facturationResponse
}

func display(information FacturationResponse) {
	fmt.Printf("Montant initial : %v CFA\n", information.Price)
	fmt.Printf("Remise : %v\n", information.Discount)
	fmt.Printf("Montant final : %v CFA\n", information.DiscountPrice)
}

func designSeparater() {
	fmt.Println("")
	fmt.Println("**************************************************")
	fmt.Println("")
}

func main() {
	fmt.Println("############ Bienvenu sur l'application de facturation ###############")

	for true {
		var price int
		var commandOption string
		fmt.Println("")
		fmt.Println("pour une nouvelle facturation taper F")
		fmt.Println("pour arrêter l'application taper A")
		fmt.Print("Choisissez une option: ")
		fmt.Scanln(&commandOption)
		if strings.ToUpper(commandOption) == "A" {
			break
		}
		if strings.ToUpper(commandOption) == "F" {
			fmt.Print("Saisissez le montant total des achats: ")
			fmt.Scanln(&price)
			resultat := facturation(price)
			designSeparater()
			display(resultat)
			designSeparater()
		} else {
			fmt.Println()
			fmt.Println("commande inconnue, ressayer !")
			continue
		}

	}

}
