# Farmer's Market Project

### Introduction
This project aims at developing a checkout system, which produces a final bill for the purchased items.Generation of total cost depends on the offers existing for that particular week. If any of the offers is applicable for the purchased items, a proper deduction should be made by the checkout system.

### Scope of the project
  - Create a checkout system.
  - Incorporate available offer intelligence to the checkout system.
  - Generate the bill for the purchased items along with applicable offer deduction.

### Assumptions
  - If the user select items which are presented in the offer and simultaneously it is also eligible for the offer containing the item, make the selected item deductible in the bill.
  - Example: user selects two items: CH1 and MK1. Make MK1 free as per the CHMK offer.


### Project design
This project is developed in Golang. Project structure mainly contains three packages:
  - Checkout - implements the checkout system logic.
  - Product - contains an inventory of available products.
  - Basket - maintains a bucket of purchased items at any point of time.

##### Checkout Package design
This package maintains the core logic of the system with two purposes:
 - Bill generation
 - Offer maintenance

###### **Offer maintenance design:**

```
type checkOutSystem struct {
	bill            []*BillFormat                 // final bill
	availableOffers map[string]OfferSpecification //list of available offers in the market
}
```
The checkout system has to include four offers as per the problem statement. An interface design created to make the design
pluggable.

```
type OfferSpecification interface {
	CheckOfferCriteria() []*BillFormat //method to be implemented by the offers
}
```
So all the four offers have to implement this interface function to make it plug into the checkout system. Example for such implementation:
```
func (a *ApplOffer) CheckOfferCriteria() []*BillFormat {
	a.profitEarned = 0
	var appleItem []*BillFormat
	prodList := basket.Basket.GetBasket()
	appleCount := (*prodList)["AP1"]
	discNeeded := false
```
The above code snippet shows APPL offer implementing the interface method.

During the package initialization, all the offers are added to the checkout system. If new offers are added to the checkout system call the below method:
```
func (s *checkOutSystem) addOffer(name string, off OfferSpecification) {
	s.availableOffers[name] = off
}
```
Designing with the help of interface can avoid the modification of existing code and make it pluggable in nature.

###### **Bill generation design:**
Bills are generated in the format as shown below:
```
type BillFormat struct {
	prodCode     string  //product code
	discountType string  //discount type availed
	price        float64 //price tag
}
```
When the user request to print the bill for display, all the offer specifications are called in a **polymorphic way** to see its availability depending on the items purchased. Finally, the bill gets generated with all the applicable deduction depending on the items purchased.

##### Product Package design
This package is used for maintaining the product inventory. In this code, the product specification are maintained in a **NoSQL database - Mongo DB**. The products are added to the database through the below function:
```
func addProduct(code, name string, price float64) {
  item := new(PSchema)
	item.Code = code
	item.Name = name
	item.Price = price

	err := productCollection.Insert(item)
	if err != nil {
		panic(err)
	}
}
```
We can also add new items in to the database and include the corresponding billing logic to show the appropriate values in the bill
generation process.

**Scalability:**
In the future, if the number of products gets increased, we can horizontally scale the database across the cluster. Additionally, we can also enhance the reliability by creating secondary databases.

##### Basket Package design
Basket package is used to keep track of purchased item at any given time. The format used to maintain is shown below:

```
type Container struct {
	prodBasket map[string]int64
}
```
The above structure is used to keep track of the current count of the purchased products. From this structure, the checkout system manages to generate the bill with corresponding offer deduction.

### Steps to run the project
 - Clone this repository in your local system.
 - Make sure the docker engine is installed before runnin the setup script.
 - Finally, run the setup shell script to create images and containers.
```sh
Run Setup Script
$ ./setup
Check the created images. Additionally, it will also starts Mongo server container.
$ docker images
Finally run the below to get in to the app container.
$ sudo docker run --link mongod  -it --name cos market /bin/bash
```
Inside the container, start running the project by typing the below command.
  - Type **market** command in the container terminal.
  - NOTE: Alias called **"market"** has been created for running the project while building the container. You can view the alias in bashrc file.
  - **"alias market="go run /home/Go_Project/src/checkoutsystem/main.go"**
```sh
To start the project type market command
$ market
1. Insert new item into the basket
2. Show register
3. Clear Basket
4. Product DB configuration
5. Exit
Enter the Choice:
<enter the required choice from the menu displayed above> On pressing 1

Farmers Market Items
Product Code	Name   	Price
CF1         	Cofee	11.230000
MK1         	Milk	4.750000
OM1         	Oatmeal	3.690000
CH1         	Chai	3.110000
AP1         	Apples	6.000000

Enter the Product Code to add it in Basket:
<enter the required item to put into the basket> ch1
1. Insert new item into the basket
2. Show register
3. Clear Basket
4. Product DB configuration
5. Exit
Enter the Choice:2 (to show the current register status)

Farmers Market Bill
*******************
Code	 Discount	Price
CH1 	            	3.110000
MK1 	            	4.750000
     	CHMK    	-4.750000
Total	       	3.110000

1. Insert new item into the basket
2. Show register
3. Clear Basket
4. Product DB configuration
5. Exit
```

Menu option provides:
  - adding several products into the basket at any time by choosing the option 1
  - checking the status of the register by choosing the option 2
  - clearing the products in the basket by pressing option 3
  - exiting the program by pressing 4

### Steps to run the test cases
Inside the container, there is also an additional alias called **marketTest***. On running this command, the list of test cases embedded into the project are started.
  - Type **marketTest** command in the container terminal.
  - NOTE: Information about the test cases ran can be found in the checkout package. All files ending with **"_test.go"** have the list of test cases started by the Golang test framework.
  - **"alias marketTest="go test checkoutsystem/checkout -v"**
```sh
type the below command to start all the test cases included in the project
$ marketTest
Farmers Market Bill
*******************
Code	 Discount	Price
OM1 	            	3.690000
OM1 	            	3.690000
OM1 	            	3.690000
AP1 	            	6.000000
     	APOM    	-3.000000
AP1 	            	6.000000
     	APOM    	-3.000000
AP1 	            	6.000000
     	APOM    	-3.000000
AP1 	            	6.000000
AP1 	            	6.000000
AP1 	            	6.000000
Total	       	38.070000

--- PASS: TestDepApplApom3 (0.00s)
PASS
ok  	checkoutsystem/checkout	0.004s
```
