package gross

func Units() map[string]int {
    return map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

func NewBill() map[string]int {
    return make(map[string]int)
}

func AddItem(bill, units map[string]int, item, unit string) bool {
    unitQty, unitExists := units[unit]
    
    if !unitExists {
        return false
    }
    
    bill[item] += unitQty
    return true
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billQty, billExists := bill[item]
	unitQty, unitExists := units[unit]
    
    if !billExists || !unitExists {
        return false
    }
    
    newQty := billQty - unitQty

    if newQty < 0 {
        return false
    }

    if newQty == 0 {
        delete(bill, item)
        return true
    }

    bill[item] = newQty
    return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	billQty, billExists := bill[item]
	return billQty, billExists  
}
