package gross

import "fmt"
// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    unitMap := map[string]int{}
    unitMap["quarter_of_a_dozen"] = 3
	unitMap["half_of_a_dozen"] = 6
	unitMap["dozen"] = 12
	unitMap["small_gross"] = 120
	unitMap["gross"] = 144
	unitMap["great_gross"] = 1728
	return unitMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	newBill := map[string]int{}
    return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	fmt.Println("Input units: \n", units)
    fmt.Println("Input bill: \n", bill)
    fmt.Println("Input item: \n", item)
    fmt.Println("Input unit: \n", unit)
    unitValue, exists := units[unit]
    fmt.Println("Input unitvalue: \n", unitValue)
    if exists == false{
        return false
    }
    billItem, exists := bill[item]
    if exists == true {
        fmt.Println("billItem: \n", billItem)
        bill[item] = bill[item] + unitValue
        return true
    } else {
        bill[item] = unitValue
        return true
    }
	
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	fmt.Println("Input units: \n", units)
    fmt.Println("Input bill: \n", bill)
    fmt.Println("Input item: \n", item)
    fmt.Println("Input unit: \n", unit)
    billItem, billExists := bill[item]
    unitValue, unitExists := units[unit]
    newQuantity := billItem - unitValue
    if billExists == false || unitExists == false || newQuantity < 0 {
        return false
    }
    fmt.Println("newQuantity: \n", newQuantity)
    if newQuantity == 0 {
        delete(bill, item)
        return true
    } else {
        bill[item] = newQuantity
        return true
    }
	
}
// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    fmt.Println("Input bill: \n", bill)
    fmt.Println("Input item: \n", item)
    billItem, billExists := bill[item]
    return billItem, billExists
}
