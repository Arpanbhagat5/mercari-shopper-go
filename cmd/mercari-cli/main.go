package main

import (
	"fmt"
	"mercari-shopper-go/internal/usecase"
)


func main() {
	fmt.Println("Mercari Shopper CLI Application Started!!")

	// create a new instance of the dummy search item usecase
	searchItemUsecase := usecase.DummySearchItemUsecase{}

	// create a dummy query
	query := "dummy object"

	// search for dummy object
	items, err := searchItemUsecase.SearchItems(query)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// print the search results
	fmt.Println("Search Results:")
	if len(items) == 0 {
		fmt.Println("No items found.")
		return
	}
	for _, item := range items {
		fmt.Println(item.Summarize())
	}
}

