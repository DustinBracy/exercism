package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int {
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitsToAdd, isValidUnit := units[unit]
	if !isValidUnit {
		return false
	}
	bill[item] += unitsToAdd
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitsToRemove, isValidUnit := units[unit]
	_, isValidItem := bill[item]
	if !(isValidItem && isValidUnit) {
		return false
	}
	newQty := bill[item] - unitsToRemove
	response := true
	switch {
	case newQty < 0:
		response = false
	case newQty == 0:
		delete(bill, item)
	default:
		bill[item] = newQty
	}
	return response
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	return qty, ok
}
