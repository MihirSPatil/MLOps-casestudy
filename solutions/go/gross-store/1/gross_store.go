package gross
import ("fmt")
// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    grossUnits := map[string]int{}
    keys := []string{"quarter_of_a_dozen", "half_of_a_dozen", "dozen", "small_gross", "gross", "great_gross"}
    scores := []int{3, 6, 12, 120, 144, 1728}

    for idx:=0; idx < len(scores); idx++{
        grossUnits[keys[idx]]=scores[idx]
    }
    return grossUnits

}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
    _, itemExists := bill[item]
    fmt.Println("exists", exists)
    if !exists{
        return exists
    }
    
    if itemExists{
        bill[item] += units[unit]
        return exists
    } else {
        bill[item] = units[unit]
        return exists
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    itemQty, itemExists := bill[item]
    unitQty, unitExists := units[unit]

    if !itemExists || !unitExists {
        return false
    }

    newQty := itemQty - unitQty

    if newQty < 0 {
        return false
    }

    if newQty == 0 {
        delete(bill, item)
    } else {
        bill[item] = newQty
    }

    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]
    if !exists{
        return 0,false
    }else{
    return qty, exists
    }
}
