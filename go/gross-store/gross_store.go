package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var units = map[string]int{
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
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	validUnit := validateKey(units, unit)

	if !validUnit {
		return false
	}

	bill[item] = units[unit]

	return validUnit
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	// Return false if the given item is not in the bill
	isValidKey := validateKey(bill, item)
	if !isValidKey {
		return false
	}
	// Return false if the given unit is not in the units map.
	isValidKey = validateKey(units, unit)
	if !isValidKey {
		return false
	}

	// Return false if the new quantity would be less than 0.
	// If the new quantity is 0, completely remove the item from the bill then return true.
	if bill[item]-units[unit] < 0 {
		return false
	} else if bill[item]-units[unit] == 0 {
		delete(bill, item)
		return true
	}

	// Otherwise, reduce the quantity of the item and return true.
	bill[item] = bill[item] - units[unit]
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	isValidKey := validateKey(bill, item)
	if !isValidKey {
		return 0, false
	}

	return bill[item], true
}

func validateKey(units map[string]int, unit string) bool {
	_, ok := units[unit]
	return ok
}
