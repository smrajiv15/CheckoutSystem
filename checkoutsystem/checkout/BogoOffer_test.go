package checkout

import (
	"checkoutsystem/basket"
	"math"
	"strconv"
	"testing"
)

func clearBasket() {
	for k := range *basket.Basket.GetBasket() {
		delete(*basket.Basket.GetBasket(), k)
	}
}

//TestBogoOffer1 - adds 3 coffee. add a free coffee and provide the bill for two.
func TestBogoOffer1(t *testing.T) {
	clearBasket()
	basket.Basket.AddItem("CF1")
	basket.Basket.AddItem("CF1")
	basket.Basket.AddItem("CF1")

	Cos.PrintBills()
	var count int64
	profitEarned := 0.0

	for _, be := range Cos.bill {
		if be.prodCode == "CF1" {
			count++
		}
		if be.discountType == "BOGO" {
			profitEarned += math.Abs(be.price)
		}
	}

	if count != 4 {
		t.Error("Totally 4 coffee, in addition to user ordered coffee. But found:  " + strconv.FormatInt(count, 10))
	}

	if profitEarned != 22.46 {
		t.Error("Profit earned mismatched: Expected 11.23. Got: " + strconv.FormatFloat(profitEarned, 'f', -1, 64))
	}
}

//TestBogoOffer2 - user adds 4 coffee. make the 2 coffee free
//Additionally it covers for all Even cases which makes the bill half the amount
func TestBogoOffer2(t *testing.T) {
	clearBasket()
	//user adding 2 coffee
	basket.Basket.AddItem("CF1")
	basket.Basket.AddItem("CF1")
	basket.Basket.AddItem("CF1")
	basket.Basket.AddItem("CF1")

	Cos.PrintBills()
	var count int64
	profitEarned := 0.0

	for _, be := range Cos.bill {
		if be.prodCode == "CF1" {
			count++
		}
		if be.discountType == "BOGO" {
			profitEarned += math.Abs(be.price)
		}
	}

	if count != 4 {
		t.Error("User Ordered 4 cofee entry. But found : " + strconv.FormatInt(count, 10))
	}

	if profitEarned != 22.46 {
		t.Error("Profit earned mismatched: Expected 11.23. Got: " + strconv.FormatFloat(profitEarned, 'f', -1, 64))
	}
}
