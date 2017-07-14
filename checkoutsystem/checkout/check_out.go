//Package checkout implements the checkout system for bill generation.
//Additionally, it also provides interface to add and remove offers to the checkout system.
package checkout

import (
	"checkoutsystem/basket"
	"fmt"
)

//BillFormat skelton of the product entry to put in to the bill.
type BillFormat struct {
	prodCode     string  //product code
	discountType string  //discount type availed
	price        float64 //price tag
}

//checkOutSystem checkOutSystem skeleton.
type checkOutSystem struct {
	bill            []*BillFormat                 // final bill
	availableOffers map[string]OfferSpecification //list of available offers in the market
}

//depCheckApomAppl function used when items in the basket eligible for APOM and APPL offer.
func depCheckApomAppl(sys *checkOutSystem, appleProfit, apomProfit float64, appleBill, apomBill *[]*BillFormat) {
	if appleProfit > 0 && apomProfit > 0 {
		if appleProfit > apomProfit {
			sys.bill = append(sys.bill, (*appleBill)...)
			prodCount := basket.Basket.GetBasket()
			oatCount := (*prodCount)["OM1"]

			for oatCount > 0 {
				newItem := new(BillFormat)
				formBillentry(newItem, "OM1", dspace, 3.69)
				sys.bill = append(sys.bill, newItem)
				oatCount--
			}
		} else if apomProfit > appleProfit || appleProfit == apomProfit {
			sys.bill = append(sys.bill, (*apomBill)...)
		}
		return
	}

	if appleProfit > 0 {
		sys.bill = append(sys.bill, (*appleBill)...)
		//added this because apple offer depends only on apple
		sys.bill = append(sys.bill, (*apomBill)...)
		return
	} else if apomProfit > 0 {
		sys.bill = append(sys.bill, (*apomBill)...)
		return
	}

	if apomProfit == 0 {
		sys.bill = append(sys.bill, (*apomBill)...)
	}
	if appleProfit == 0 {
		sys.bill = append(sys.bill, (*appleBill)...)
	}
}

//processOffers function is used to generate the bill as per the selected items.
func processOffers(sys *checkOutSystem) {
	appleBill := []*BillFormat{}
	apomBill := []*BillFormat{}

	appleProfit := 0.00
	apomProfit := 0.00

	for _, v := range sys.availableOffers {
		switch v.(type) {
		case *BogoOffer:
			sys.bill = append(sys.bill, v.CheckOfferCriteria()...)
		case *ApplOffer:
			appOff := v.(*ApplOffer)
			appleBill = v.CheckOfferCriteria()
			appleProfit = appOff.profitEarned
		case *ApomOffer:
			apomOff := v.(*ApomOffer)
			apomBill = v.CheckOfferCriteria()
			apomProfit = apomOff.profitEarned
		case *ChmkOffer:
			sys.bill = append(sys.bill, v.CheckOfferCriteria()...)
		}
	}
	depCheckApomAppl(sys, appleProfit, apomProfit, &appleBill, &apomBill)
}

//PrintBills helper function to print bill to the user.
func (s *checkOutSystem) PrintBills() {
	s.bill = []*BillFormat{}
	processOffers(s)
	totolPrice := 0.0

	fmt.Println("Farmers Market Bill")
	fmt.Println("*******************")

	fmt.Println("Code\t Discount\tPrice")

	for _, prodEntry := range s.bill {
		fmt.Printf("%s \t%s    \t%f\n", prodEntry.prodCode, prodEntry.discountType, prodEntry.price)
		totolPrice += prodEntry.price
	}
	fmt.Printf("Total\t%s\t%f\n", "       ", totolPrice)
	fmt.Println()
}

//addOffer Funtion is used to add new offer to the checkout system.
func (s *checkOutSystem) addOffer(name string, off OfferSpecification) {
	s.availableOffers[name] = off
}

//deleteOffer Funtion is used to delete the offer from the checkout System.
func (s *checkOutSystem) deleteOffer(name string) {
	delete(s.availableOffers, name)
}

//Printoffers helper function to display the avaiable offers in the checkout system.
func (s *checkOutSystem) Printoffers() {
	for k := range s.availableOffers {
		fmt.Printf("%s\n", k)
	}
}

//Cos Global handler for the checkOutSystem
var Cos checkOutSystem

func init() {
	Cos.availableOffers = make(map[string]OfferSpecification)
	Cos.addOffer("BOGO", new(BogoOffer))
	Cos.addOffer("APPL", new(ApplOffer))
	Cos.addOffer("CHMK", new(ChmkOffer))
	Cos.addOffer("APOM", new(ApomOffer))
}
