package main

import "fmt"

// The Car class represents a specific car in the store.
type Car struct {
	Make  string
	Model string
	Year  int
	Product
}

// The Product class represents a product in the store, including cars.
// It has attributes of a product, such as its name, quantity in stock, and price.
type Product struct {
	Name     string
	Quantity int
	Price    float64
}

// The ProductInterface interface defines the methods that a Product must implement.
type ProductInterface interface {
	DisplayProduct()
	DisplayStatus()
	UpdateQuantity(int) error
}

// DisplayProduct is a method of the Product class that displays the product's information.
func (p Product) DisplayProduct() {
	fmt.Printf("Product: %s\n", p.Name)
	fmt.Printf("Quantity: %d\n", p.Quantity)
	fmt.Printf("Price: $%.2f\n", p.Price)
}

// DisplayStatus is a method of the Product class that displays the product's status (in stock or out of stock).
func (p Product) DisplayStatus() {
	if p.Quantity > 0 {
		fmt.Println("Status: In stock")
	} else {
		fmt.Println("Status: Out of stock")
	}
}

// UpdateQuantity is a method of the Product class that updates the product's quantity in stock.
func (p Product) UpdateQuantity(quantity int) error {
	if p.Quantity+quantity < 0 {
		return fmt.Errorf("cannot set quantity to a negative value")
	}
	p.Quantity += quantity
	return nil
}

// The Store class represents a store that sells products, including cars.
// It has a list of products in stock and a list of sold products.
type Store struct {
	Products     []ProductInterface
	SoldProducts []ProductInterface
}

// The StoreInterface interface defines the methods that a Store must implement.
type StoreInterface interface {
	AddProduct(ProductInterface)
	ListProducts()
	SellProduct(string) error
	ListSoldProducts()
	SearchProduct(string) ProductInterface
}

// AddProduct is a method of the Store class that adds a product to the store's list of products in stock.
func (s *Store) AddProduct(p ProductInterface) {
	s.Products = append(s.Products, p)
}

// ListProducts is a method of the Store class that lists all  the products in the store's list of products in stock.
func (s *Store) ListProducts() {
	for _, p := range s.Products {
		p.DisplayProduct()
		p.DisplayStatus()
		fmt.Println()
	}
}

// SellProduct is a method of the Store class that sells a product from the store's list of products in stock and adds it to the store's list of sold products.
func (s *Store) SellProduct(name string) error {

	// Search for the product in the store's list of products in stock.
	product := s.SearchProduct(name)
	if product == nil {
		return fmt.Errorf("product not found")
	}

	//Asser the product is a Product type.
	p, ok := product.(*Product)
	if !ok {
		return fmt.Errorf("product is not a Product type")
	}
	// Check if there is sufficient quantity in stock to sell the product.
	if p.Quantity < 1 {
		return fmt.Errorf("product is out of stock")
	}

	// Remove the product from the store's list of products in stock and add it to the store's list of sold products.
	for i, p := range s.Products {

		// Assert that the p variable is of type Product.
		p, ok := p.(*Product)
		if !ok {
			return fmt.Errorf("product has wrong type")
		}

		if p.Name == name {
			s.SoldProducts = append(s.SoldProducts, p)
			s.Products = append(s.Products[:i], s.Products[i+1:]...)
			break
		}
	}

	// Update the product's quantity in stock.
	err := product.UpdateQuantity(-1)
	if err != nil {
		return err
	}

	return nil
}

// ListSoldProducts is a method of the Store class that lists all the products in the store's list of sold products.
func (s *Store) ListSoldProducts() {
	for _, p := range s.SoldProducts {
		p.DisplayProduct()
		fmt.Println()
	}
}

// SearchProduct is a method of the Store class that searches for a product with a given name in the store's list of products in stock.
// It returns the product if found, or nil if not found.
func (s *Store) SearchProduct(name string) ProductInterface {
	for _, p := range s.Products {

		// Assert that the p variable is of type Product.
		p, ok := p.(*Product)
		if !ok {
			return nil
		}
		if p.Name == name {
			return p
		}
	}
	return nil
}

func main() {
	// Create a new store.
	store := &Store{}

	// Add some cars to the store.
	store.AddProduct(&Car{Make: "Toyota", Model: "Camry", Year: 2020, Product: Product{Name: "Toyota Camry", Quantity: 3, Price: 30000}})
	store.AddProduct(&Car{Make: "Honda", Model: "Accord", Year: 2021, Product: Product{Name: "Honda Accord", Quantity: 5, Price: 35000}})
	store.AddProduct(&Car{Make: "Ford", Model: "Mustang", Year: 2019, Product: Product{Name: "Ford Mustang", Quantity: 2, Price: 40000}})

	// List the products in the store.
	fmt.Println("Products in stock:")
	store.ListProducts()
	fmt.Println()

	// Sell a car from the store.
	err := store.SellProduct("Toyota Camry")
	if err != nil {
		fmt.Println(err)
	}

	// List the products in the store again.
	fmt.Println("\nProducts in stock:")
	store.ListProducts()
	fmt.Println()

	// List the sold products.
	fmt.Println("\nSold products:")
	store.ListSoldProducts()
}
