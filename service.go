package go_auth_api_tests2

import "errors"

type Product struct {
	ID    int
	Name  string
	Price int
}

func GetProduct(id int) (Product, error) {
	if id <= 0 {
		return Product{}, errors.New("id must be a positive integer")
	}
	if id == 1 {
		return Product{ID: 1, Name: "Book", Price: 100}, nil
	}
	if id == 2 {
		return Product{2, "Pen", 20}, nil
	}
	return Product{}, errors.New("not found product")
}
