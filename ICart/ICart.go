package main

import "fmt"

type product struct {
	id   string
	name string
}

type productList []product //products list

type icart interface {
	get(id string) (product, int)
	getAll()
	add(item product)
	update(updatedEntry product, index int)
	remove(index int)
}

func print(item product) {
	fmt.Printf("ID:%v	Name:%v\n", item.id, item.name)
}

func (products productList) get(id string) (product, int) {
	for index, item := range products {
		if item.id == id {
			return item, index
		}
	}
	return product{}, -1
}

func (products productList) getAll() {

	if len(products) == 0 {
		fmt.Println("No Product Available")
	} else {

		//	productList := icart.list() // make(productList, len(icart.list()))
		for _, item := range products {

			print(item)
		}
		//	}

		//return products
	}
}

func (products *productList) add(item product) {
	*products = append(*products, item)
}

func (products *productList) update(updatedEntry product, index int) {

	(*products)[index] = updatedEntry

}

func (products *productList) remove(index int) {
	length := len(*products)
	(*products)[index] = (*products)[length-1]
	*products = (*products)[:length-1]

}

func main() {

	var cart icart = &productList{}

	fmt.Println("Press 1: to Add an product to the list")
	fmt.Println("Press 2: to Update an product to the list")
	fmt.Println("Press 3: to Delete an product to the list")
	fmt.Println("Press 4: to Get an product to the list")
	fmt.Println("Press 5: to Get Whole product list")
	fmt.Println("Press 0: to exit")
	var input int
	for {
		fmt.Println("Enter ur choice")
		fmt.Scanf("%v", &input)
		if input == 0 {
			break
		}
		switch input {
		case 1:
			fmt.Println("enter id of new product")
			var ID, Name string
			fmt.Scanf("%v", &ID)
			fmt.Println("enter name of new product")
			fmt.Scanf("%v", &Name)

			newProduct := product{
				id: ID, name: Name,
			}
			cart.add(newProduct)
		case 2:
			fmt.Println("enter id of the product to be updated")
			var ID, Name string
			fmt.Scanf("%v", &ID)

			_, index := cart.get(ID)
			if index != -1 {
				fmt.Println("enter new name for the product")
				fmt.Scanf("%v", &Name)
				newProduct := product{
					id: ID, name: Name,
				}
				cart.update(newProduct, index)
			} else {
				fmt.Println("Product not found")
			}
		case 3:
			fmt.Println("enter id of the product to be removed")
			var ID string
			fmt.Scanf("%v", &ID)
			_, index := cart.get(ID)
			if index != -1 {
				cart.remove(index)
			} else {
				fmt.Println("Product not found")
			}
		case 4:
			fmt.Println("enter id of the product to be displayed")
			var ID string
			fmt.Scanf("%v", &ID)
			product, index := cart.get(ID)
			if index != -1 {
				print(product)
			} else {
				fmt.Println("Product not found")
			}
		case 5:
			cart.getAll()
		default:
			fmt.Println("Entered Wrong Choice")
		}
	}
}
