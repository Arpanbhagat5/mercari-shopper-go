package main

import (
	"fmt"
	"mercari-shopper-go/internal/interface-adaptors/api"
	"mercari-shopper-go/internal/usecase"
)


func main() {
	fmt.Println("Mercari Shopper CLI Application Started!!")

	// --- Dummy api search with dummy api client---
	searchItemClient := api.DummyMercariAPIClient{}
	searchItemUseCase := usecase.NewRealSearchItemsUseCase(searchItemClient)


	api_query := "dummy api object"
	items, err := searchItemUseCase.SearchItems(api_query)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("\nSearch results (via API Adapter - but dummy API for now) for:", api_query)
	if len(items) == 0 {
		fmt.Println("No items found.")
		return
	}
	for _, item := range items {
		fmt.Println(item.Summarize())
	}

	// --- end of dummy api search with dummy api client ---

	// ------------------------------------------------------------

	// --- Dummy api search with dummy hardcoded response ---
	searchItemsUseCase := usecase.DummySearchItemUsecase{}

	// create a dummy query
	dummy_query := "dummy object"
	// search for dummy object
	dummy_items, err := searchItemsUseCase.SearchItems(dummy_query)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// print the search results
	fmt.Println("Search Results (via hardcoded data) for:", dummy_query)
	if len(dummy_items) == 0 {
		fmt.Println("No dummy_items found.")
		return
	}
	for _, item := range dummy_items {
		fmt.Println(item.Summarize())
	}
	// --- end of dummy api search with dummy hardcoded response ---

}

