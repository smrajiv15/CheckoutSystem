//Package basket implements basket for storing items selected by users
package basket

//Container Basket containing the items selected by the customer
type Container struct {
	prodBasket map[string]int64
}

//Basket Global handle for basket containing list of products selected by user
var Basket Container

//AddItem adds item in to the basket as per the user selection
func (b *Container) AddItem(code string) {
	b.prodBasket[code]++
}

//GetBasket returns the basket handle containing the selected products
func (b *Container) GetBasket() *map[string]int64 {
	return &b.prodBasket
}

func init() {
	Basket.prodBasket = make(map[string]int64)
}
