package main

import (
	"fmt"
	"log"
	"time"
)

type Cart struct {
	ID          string
	CreatedDate time.Time
	Items       []Item
}

type Item struct {
	SKU      string
	Quantity int
}

func main() {
	var myPointerVar *int
	// print zero of pointer value type -> nil
	fmt.Println(myPointerVar)

	// map
	cities := make(map[string]string)
	addElement(cities)
	log.Println(cities)

	// slices: structure has three fields:
	// - A length
	// - A capacity
	// - A pointer to an internal array
	EUountries := []string{"Austria", "Belgium", "Bulgaria"}
	log.Println(EUountries)

	cart := Cart{
		ID:          "1234",
		CreatedDate: time.Now(),
	}
	cartPtr := &cart
	log.Println(cartPtr)

	// automatic dereference
	cartPtr.Items = []Item{
		{SKU: "154550", Quantity: 12},
		{SKU: "DTY8755", Quantity: 1},
	}
	log.Println(cart)
}

func addElement(cities map[string]string) {
	cities["France"] = "Paris"
}
