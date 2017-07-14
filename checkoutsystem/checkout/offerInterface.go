package checkout

import "checkoutsystem/basket"

const dspace = "        "
const pspace = "    "

//formBillentry helper function for form product entry to be added in to the bill.
func formBillentry(b *BillFormat, code string, dtype string, price float64) {
	b.prodCode = code
	b.discountType = dtype
	b.price = price
}

//OfferSpecification interface to extend different offers.
type OfferSpecification interface {
	CheckOfferCriteria() []*BillFormat //method to be implemented by the offers
}

//BogoOffer house keeping for Buy one get One offer.
type BogoOffer struct {
}

//CheckOfferCriteria checks for Bogo offer in the basket.
func (b *BogoOffer) CheckOfferCriteria() []*BillFormat {
	var coffeeList []*BillFormat
	prodList := basket.Basket.GetBasket()
	coffeeCount := (*prodList)["CF1"]

	if coffeeCount == 0 {
		return nil
	}

	if oeFlag := coffeeCount % 2; oeFlag == 1 {
		coffeeCount++
	}

	for coffeeCount > 0 {

		billEntry := new(BillFormat)
		formBillentry(billEntry, "CF1", dspace, 11.23)
		coffeeList = append(coffeeList, billEntry)

		if coffeeCount%2 == 1 {
			billEntry = new(BillFormat)
			formBillentry(billEntry, pspace, "BOGO", -11.23)
			coffeeList = append(coffeeList, billEntry)
		}
		coffeeCount--
	}

	return coffeeList
}

//ApplOffer house keeping for APPL offer.
type ApplOffer struct {
	profitEarned float64
}

//CheckOfferCriteria checks for APPL offer in the basket.
func (a *ApplOffer) CheckOfferCriteria() []*BillFormat {
	a.profitEarned = 0
	var appleItem []*BillFormat
	prodList := basket.Basket.GetBasket()
	appleCount := (*prodList)["AP1"]
	discNeeded := false

	if appleCount == 0 {
		return nil
	}

	if appleCount >= 3 {
		discNeeded = true
	}

	for appleCount > 0 {
		newItem := new(BillFormat)
		formBillentry(newItem, "AP1", dspace, 6.00)
		appleItem = append(appleItem, newItem)

		if discNeeded {
			newItem = new(BillFormat)
			formBillentry(newItem, pspace, "APPL", -1.50)
			a.profitEarned += 1.50
			appleItem = append(appleItem, newItem)
		}
		appleCount--
	}
	return appleItem
}

//ChmkOffer house keeping for CHMK in the basket.
type ChmkOffer struct {
}

//CheckOfferCriteria checks for CHMK offer in the basket.
func (c *ChmkOffer) CheckOfferCriteria() []*BillFormat {
	var itemList []*BillFormat
	prodList := basket.Basket.GetBasket()

	chaiCount := (*prodList)["CH1"]
	milkCount := (*prodList)["MK1"]

	for chaiCount > 0 {
		chaiItem := new(BillFormat)
		formBillentry(chaiItem, "CH1", dspace, 3.11)
		itemList = append(itemList, chaiItem)
		chaiCount--
	}
	chaiCount = (*prodList)["CH1"]

	for milkCount > 0 {
		milkItem := new(BillFormat)
		formBillentry(milkItem, "MK1", dspace, 4.75)
		itemList = append(itemList, milkItem)

		if milkCount == 1 && chaiCount > 0 {
			milkItem = new(BillFormat)
			formBillentry(milkItem, pspace, "CHMK", -4.75)
			itemList = append(itemList, milkItem)
		}
		milkCount--
	}

	milkCount = (*prodList)["MK1"]
	if milkCount > 0 {
		return itemList
	}

	if milkCount == 0 && chaiCount > 0 {
		milkItem := new(BillFormat)
		formBillentry(milkItem, "MK1", dspace, 4.75)
		itemList = append(itemList, milkItem)

		milkItem = new(BillFormat)
		formBillentry(milkItem, pspace, "CHMK", -4.75)
		itemList = append(itemList, milkItem)
	}

	return itemList
}

//ApomOffer house keeping for APOM offer in the Basket.
type ApomOffer struct {
	profitEarned float64
}

//CheckOfferCriteria checks for APOM offer in the basket.
func (a *ApomOffer) CheckOfferCriteria() []*BillFormat {
	a.profitEarned = 0
	var itemList []*BillFormat
	prodList := basket.Basket.GetBasket()

	oatCount := (*prodList)["OM1"]

	if oatCount == 0 {
		return nil
	}

	for oatCount > 0 {
		newItem := new(BillFormat)
		formBillentry(newItem, "OM1", dspace, 3.69)
		itemList = append(itemList, newItem)
		oatCount--
	}

	oatCount = (*prodList)["OM1"]
	appleCount := (*prodList)["AP1"]

	if appleCount != 0 {
		for oatCount > 0 && appleCount > 0 {
			newItem := new(BillFormat)
			formBillentry(newItem, "AP1", dspace, 6.00)
			itemList = append(itemList, newItem)

			newItem = new(BillFormat)
			formBillentry(newItem, pspace, "APOM", -3.00)
			a.profitEarned += 3
			itemList = append(itemList, newItem)
			oatCount--
			appleCount--
		}
	}

	if appleCount > 0 {
		for appleCount > 0 {
			newItem := new(BillFormat)
			formBillentry(newItem, "AP1", dspace, 6.00)
			itemList = append(itemList, newItem)
			appleCount--
		}
	}

	return itemList
}
