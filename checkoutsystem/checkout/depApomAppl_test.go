package checkout

import (
	"checkoutsystem/basket"
	"math"
	"testing"
)

//TestDepApplApom1 - oatmeal covering all the apple order. Ex: (10 om, 7 apple)
//this case will always fall in to APOM offer (eventough both have profits)
func TestDepApplApom1(t *testing.T) {
	clearBasket()
	for i := 0; i < 10; i++ {
		basket.Basket.AddItem("OM1")
	}

	for i := 0; i < 7; i++ {
		basket.Basket.AddItem("AP1")
	}

	Cos.PrintBills()

	var profitEarned float64

	for _, be := range Cos.bill {
		if be.discountType == "APOM" {
			profitEarned += math.Abs(be.price)
		}
	}

	if profitEarned != 21.0 {
		t.Errorf("profit Mismatch. Expected 21, But found: %f$", profitEarned)
	}
}

//TestDepApplApom2 - both APOM and APPL offer profit created.
//Make APPL profitable
func TestDepApplApom2(t *testing.T) {
	clearBasket()

	for i := 0; i < 3; i++ {
		basket.Basket.AddItem("AP1")
	}

	basket.Basket.AddItem("OM1")
	Cos.PrintBills()

	var profitEarned float64

	for _, be := range Cos.bill {

		if be.discountType == "APPL" {
			profitEarned += math.Abs(be.price)
		}
	}

	if profitEarned != 4.5 {
		t.Errorf("profit Mismatch. Expected 4.5, But found: %f $", profitEarned)
	}
}

//TestDepApplApom3 - Check for profit equality case. By default it chooses APOM
func TestDepApplApom3(t *testing.T) {
	clearBasket()

	for i := 0; i < 6; i++ {
		basket.Basket.AddItem("AP1")
	}

	for i := 0; i < 3; i++ {
		basket.Basket.AddItem("OM1")
	}

	Cos.PrintBills()
	var profitEarned float64

	for _, be := range Cos.bill {
		if be.discountType == "APOM" {
			profitEarned += math.Abs(be.price)
		}
	}

	if profitEarned != 9 {
		t.Errorf("profit Mismatch. Expected 4.5, But found: %f $", profitEarned)
	}
}
