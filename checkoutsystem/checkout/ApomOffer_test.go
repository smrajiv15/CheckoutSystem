package checkout

import (
	"checkoutsystem/basket"
	"math"
	"testing"
)

//TestAppomoffer1 - order only oatmeal. No offer eligible
func TestApomoffer1(t *testing.T) {
	clearBasket()
	basket.Basket.AddItem("OM1")
	basket.Basket.AddItem("OM1")
	basket.Basket.AddItem("OM1")
	basket.Basket.AddItem("OM1")

	Cos.PrintBills()

	var totalItem, oatCount int64

	for _, be := range Cos.bill {
		if be.prodCode == "OM1" {
			oatCount++
		}
		totalItem++
	}

	if totalItem != oatCount {
		t.Errorf("ordered 4 oatmeal. But found: %d oatmeal", oatCount)
	}
}

func TestApomoffer2(t *testing.T) {
	clearBasket()
	basket.Basket.AddItem("OM1")
	basket.Basket.AddItem("OM1")
	basket.Basket.AddItem("OM1")
	basket.Basket.AddItem("OM1")

	basket.Basket.AddItem("AP1")
	basket.Basket.AddItem("AP1")

	Cos.PrintBills()

	var oatCount, appleCount int64
	var profitEarned float64

	for _, be := range Cos.bill {
		if be.prodCode == "OM1" {
			oatCount++
		}

		if be.prodCode == "AP1" {
			appleCount++
		}

		if be.discountType == "APOM" {
			profitEarned += math.Abs(be.price)
		}
	}

	if appleCount != 2 {
		t.Errorf("ordered 2 apple. But found: %d apples", appleCount)
	}

	if oatCount != 4 {
		t.Errorf("ordered 4 oatmeal. But found: %d oatmeal", oatCount)
	}
}
