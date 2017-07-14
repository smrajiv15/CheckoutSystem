package checkout

import (
	"checkoutsystem/basket"
	"math"
	"testing"
)

//TestApploffer1 - Buy less than 3 apple. No offer
func TestApploffer1(t *testing.T) {
	clearBasket()

	basket.Basket.AddItem("AP1")
	basket.Basket.AddItem("AP1")

	var totalItem, appleCount int64

	Cos.PrintBills()

	for _, be := range Cos.bill {
		totalItem++

		if be.prodCode == "AP1" {
			appleCount++
		}
	}

	if totalItem != appleCount {
		t.Errorf("ordered 2 apple. But found: %d", appleCount)
	}
}

func TestApploffer2(t *testing.T) {
	clearBasket()
	basket.Basket.AddItem("AP1")
	basket.Basket.AddItem("AP1")
	basket.Basket.AddItem("AP1")
	basket.Basket.AddItem("AP1")

	Cos.PrintBills()

	var appleCount int16
	var profitEarned float64

	for _, be := range Cos.bill {

		if be.prodCode == "AP1" {
			appleCount++
		}

		if be.discountType == "APPL" {
			profitEarned += math.Abs(be.price)
		}
	}

	if appleCount != 4 {
		t.Errorf("ordered 4 apple. But found: %d apples", appleCount)
	}

	if profitEarned != 6.0 {
		t.Errorf("profit Mismatch. Expected 6.0, But found: %f$", profitEarned)
	}

}
