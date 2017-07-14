//Package main contains user interface logic for pruchasing the products from the market.
package main

import (
	"bufio"
	basket "checkoutsystem/basket"
	co "checkoutsystem/checkout"
	product "checkoutsystem/product"
	"fmt"
	"os"
	"strings"
)

//clearBasket refresh basket to start including new items
func clearBasket() {
	for k := range *basket.Basket.GetBasket() {
		delete(*basket.Basket.GetBasket(), k)
	}
}

func insertProduct() {
	var code, name string
	var price float64

newcode:
	fmt.Printf("New Product Insert\n")
	fmt.Println("==================")
	fmt.Println()
	fmt.Printf("Enter the product code: ")
	fmt.Scanf("%s", &code)

	if product.CheckProdExist(strings.ToUpper(code)) == true {
		fmt.Println("Error: Existing product code")
		goto newcode
	}

	fmt.Printf("Enter the product name: ")
	fmt.Scanf("%s", &name)

	fmt.Printf("Enter the product price: ")
	fmt.Scanf("%f", &price)
	product.AddProduct(code, name, price)
}

func removeProduct() {
	code := ""
	fmt.Printf("Enter the product Code to delete: ")
	fmt.Scanf("%s", &code)
	product.DeleteProduct(strings.ToUpper(code))
}

func dbConfigureMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Println("Product Configuration Menu")
		fmt.Println("==========================")
		fmt.Println("1. Add New Product")
		fmt.Println("2. Remove Existing Product")
		fmt.Println("3. Exit Configuration")

		choice := 0
		fmt.Printf("\nEnter your choice: ")
		_, err := fmt.Scanf("%d", &choice)
		fmt.Println()

		if err != nil {
			fmt.Println("Error: Invalid Input. Expected Integer")
			reader.ReadString('\n')
			fmt.Println()
		}

		switch choice {
		case 1:
			insertProduct()
		case 2:
			removeProduct()
		default:
			return
		}

	}

}

//displayMenu helper funtion for showing the menu for purchasing the product.
func displayMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Println("The Farmer's Market")
		fmt.Println("-------------------")
		fmt.Println("1. Insert new item into the basket")
		fmt.Println("2. Show register")
		fmt.Println("3. Clear Basket")
		fmt.Println("4. Product DB configuration")
		fmt.Println("5. Exit")

		fmt.Println()
		fmt.Println("Enter the Choice: ")
		choice := 0
		_, err := fmt.Scanf("%d", &choice)
		fmt.Println()

		if err != nil {
			fmt.Println("Error: Invalid Input. Expected Integer")
			reader.ReadString('\n')
			fmt.Println()
		}

		switch choice {
		case 1:
			product.ShowItems()
			fmt.Println("\nEnter the Product Code to add it to Basket: ")
			itemID, _ := reader.ReadString('\n')
			itemID = strings.ToUpper(strings.TrimRight(itemID, "\n"))

			if product.CheckProdExist(itemID) != true {
				fmt.Println("Error: Invalid Product Code. Enter again")
				fmt.Println()
				continue
			}
			basket.Basket.AddItem(strings.ToUpper(itemID))
		case 2:
			co.Cos.PrintBills()
		case 3:
			clearBasket()
		case 4:
			dbConfigureMenu()
		default:
			product.Prodcleanup()
			os.Exit(0)
		}
	}
}

func main() {
	displayMenu()
}
