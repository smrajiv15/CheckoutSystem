package checkout

import (
	"checkoutsystem/basket"
	"math"
	"strconv"
	"testing"
)

//TestChmkoffer1 - order single chai. get 1 milk free
func TestChmkoffer1(t *testing.T) {
	clearBasket()
	basket.Basket.AddItem("CH1")
	Cos.PrintBills()

	var numItems, chaiItems, milkItems int64
	var discount float64

	for _, be := range Cos.bill {
		numItems++
		if be.prodCode == "CH1" {
			chaiItems++
		}
		if be.prodCode == "MK1" {
			milkItems++
		}
		if be.discountType == "CHMK" {
			discount += math.Abs(be.price)
		}
	}

	if chaiItems != 1 {
		t.Error("User Ordered 1 Chai. But found: " + strconv.FormatInt(chaiItems, 10) + " Chai")
	}
	if milkItems != 1 {
		t.Error("1 Milk free for 1 ordered chai. But found: " + strconv.FormatInt(milkItems, 10) + "Milk items")
	}

	if discount != 4.75 {
		t.Error("profit Mismatch. Expected 4.75 But found: " + strconv.FormatFloat(discount, 'f', -1, 64) + "$")
	}
}

//TestChmkoffer2 - order multiple chai. get only a free milk (limit 1)
func TestChmkoffer2(t *testing.T) {
	clearBasket()
	basket.Basket.AddItem("CH1")
	basket.Basket.AddItem("CH1")
	basket.Basket.AddItem("CH1")

	Cos.PrintBills()

	var chaiCount, milkCount int64
	var profitEarned float64

	for _, be := range Cos.bill {
		if be.prodCode == "CH1" {
			chaiCount++
		}

		if be.prodCode == "MK1" {
			milkCount++
		}

		if be.discountType == "CHMK" {
			profitEarned += math.Abs(be.price)
		}
	}

	if chaiCount != 3 {
		t.Error("User Ordered 3 Chai. But found: " + strconv.FormatInt(chaiCount, 10) + " Chai")
	}

	if milkCount != 1 {
		t.Error("For any number of chai, only 1 milk free. But found: " + strconv.FormatInt(milkCount, 10))
	}

	if profitEarned != 4.75 {
		t.Error("profit Mismatch. Expected 4.75, But found: " + strconv.FormatFloat(profitEarned, 'f', -1, 64) + "$")
	}
}

//TestChmkoffer3 - order only milk. No offer
func TestChmkoffer3(t *testing.T) {
	clearBasket()

	basket.Basket.AddItem("MK1")
	basket.Basket.AddItem("MK1")
	basket.Basket.AddItem("MK1")

	Cos.PrintBills()

	var milkCount int64
	var totalPrice float64

	for _, be := range Cos.bill {
		if be.prodCode == "MK1" {
			milkCount++
		}
		totalPrice += be.price
	}

	if milkCount != 3 {
		t.Error("Ordered 3 milk. But got: " + strconv.FormatInt(milkCount, 10) + "Milk")
	}

	if totalPrice != 14.25 {
		t.Error("Total Mismatch. Expected 14.25, But found: " + strconv.FormatFloat(totalPrice, 'f', -1, 64) + "$")
	}
}

//TestChmkoffer4 - order multiple chai and multiple milk
func TestChmkoffer4(t *testing.T) {
	clearBasket()
	basket.Basket.AddItem("CH1")
	basket.Basket.AddItem("CH1")

	basket.Basket.AddItem("MK1")
	basket.Basket.AddItem("MK1")
	basket.Basket.AddItem("MK1")

	Cos.PrintBills()

	var chaiCount, milkCount int64
	var profitEarned float64

	for _, be := range Cos.bill {
		if be.prodCode == "CH1" {
			chaiCount++
		}

		if be.prodCode == "MK1" {
			milkCount++
		}

		if be.discountType == "CHMK" {
			profitEarned += math.Abs(be.price)
		}
	}

	if chaiCount != 2 {
		t.Error("User Ordered 2 Chai. But found: " + strconv.FormatInt(chaiCount, 10) + " Chai")
	}

	if milkCount != 3 {
		t.Error("User Ordered 3 Milk. But found: " + strconv.FormatInt(milkCount, 10) + " Chai")
	}

	if profitEarned != 4.75 {
		t.Error("profit Mismatch. Expected 4.75, But found: " + strconv.FormatFloat(profitEarned, 'f', -1, 64) + "$")
	}
}
