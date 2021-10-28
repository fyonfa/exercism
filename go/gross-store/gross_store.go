package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	// measurements := make(map[string]int) // also like this
	// map literal
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	//1 check that units exists:
	value, exists := units[unit]
	if !exists {
		return false
	}
	//2 add item to the bill if exist
	bill[item] = value
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	valueUnits, existsUnits := units[unit]
	if !existsUnits {
		return false
	}
	if bill[item] >= valueUnits {
		bill[item] -= valueUnits
		if bill[item] == 0 {
			delete(bill, item)
		}
		return true
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, exist := bill[item]
	if !exist {
		return 0, false
	}
	return bill[item], true
}
