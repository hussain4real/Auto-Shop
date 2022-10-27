package main

import "fmt"

// The Product class should have attributes of a product i.e (the product, quantity of the product in stock, price of the product)
type Product struct {
	Name     string
	Quantity int
	Price    float64
}

// A car is a product of the store, but there can be other products so the attribute of the car can be promoted to the Product.
type Car struct {
	Product
}

// The Product class should have methods to display a product, and a method to display the status of a product if it is still in stock or not.
type ProductInterface interface {
	DisplayProduct()
	DisplayStatus()
}

// The Store class should have

// function to display the product
func (p Product) DisplayProduct() {
	fmt.Printf("Product: %s", p.Name)

}

// function to display the status of the product
func (p Product) DisplayStatus() {
	if p.Quantity > 0 {
		fmt.Println("In stock")
	} else {
		fmt.Println("Out of stock")
	}
}

// The Store class should have attributes like Number of products in the store that are still up for sale,Adding an Item to the store,Listing all product items in the store,Sell an item,Show a list of sold items and the total price
type Store struct {
	Product     []ProductInterface
	soldProduct []ProductInterface
}

// The Store class should have methods to add a product, list all products, sell a product, and show a list of sold products.
type StoreInterface interface {
	AddProduct(ProductInterface)
	ListProducts()
	SellProduct(string)
	ListSoldProducts()
}

// The Store class should have methods to add a product, list all products, sell a product, and show a list of sold products.
func (s *Store) AddProduct(p ProductInterface) {
	s.Product = append(s.Product, p)
}

// The Store class should have methods to add a product, list all products, sell a product, and show a list of sold products.
func (s *Store) ListProducts() {
	for _, p := range s.Product {
		p.DisplayProduct()
	}
}

// The Store class should have methods to add a product, list all products, sell a product, and show a list of sold products.
func (s *Store) SellProduct(name string) {
	//	Loop through the products in the store
	for i, p := range s.Product {
		//		If the product is found, then remove it from the store and add it to the sold products slice
		if p.(Product).Name == name {
			s.soldProduct = append(s.soldProduct, p)
			s.Product = append(s.Product[:i], s.Product[i+1:]...)
		}
	}
}
