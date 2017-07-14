//Package product helps to add and delete products available in the market
package product

import (
	"fmt"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getDbSession() *mgo.Session {
	sess, err := mgo.Dial("mongod:27017")
	if err != nil {
		panic(err)
	}
	return sess
}

var sess *mgo.Session
var productCollection *mgo.Collection

//PSchema provides Layout for a product entry to be added in the inventry
type PSchema struct {
	Code  string  //code for a product
	Name  string  //name of the product
	Price float64 //price of the product
}

//Prodcleanup - clean need for DB
func Prodcleanup() {
	sess.Close()
}

//ShowItems helper function for showing the available product in the market.
func ShowItems() {

	var item []PSchema
	productCollection.Find(nil).All(&item)
	fmt.Println("\nFarmers Market Items")

	fmt.Printf("Product Code\tName   \tPrice\n")
	for _, item := range item {
		fmt.Printf("%-12s\t%4s\t%f\n", item.Code, item.Name, item.Price)
	}
}

//AddProduct adds the new product in the markets inventry
func AddProduct(code, name string, price float64) {
	item := new(PSchema)
	item.Code = code
	item.Name = name
	item.Price = price

	err := productCollection.Insert(item)
	if err != nil {
		panic(err)
	}
}

//CheckProdExist - find it is a valid product ID
func CheckProdExist(code string) bool {
	n, _ := productCollection.Find(bson.M{"code": code}).Count()

	if n > 0 {
		return true
	}
	return false
}

//DeleteProduct deletes the product from the inventory.
func DeleteProduct(code string) {
	if CheckProdExist(code) != true {
		fmt.Println("Error: Invalid Product Code. Enter again")
		fmt.Println()
		return
	}

	err := productCollection.Remove(bson.M{"code": code})
	if err != nil {
		panic(err)
	}
}

func init() {
	sess = getDbSession()
	dbs, _ := sess.DatabaseNames()

	for _, db := range dbs {
		if strings.Compare("Market", db) == 0 {
			productCollection = sess.DB("Market").C("items")
			return
		}
	}

	productCollection = sess.DB("Market").C("items")

	AddProduct("CH1", "Chai", 3.11)
	AddProduct("AP1", "Apples", 6.00)
	AddProduct("CF1", "Cofee", 11.23)
	AddProduct("MK1", "Milk", 4.75)
	AddProduct("OM1", "Oatmeal", 3.69)
}
