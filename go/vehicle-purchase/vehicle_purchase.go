package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	msg := "%s is clearly the better choice."

	if option1 < option2 {
		return fmt.Sprintf(msg, option1)
	}

	return fmt.Sprintf(msg, option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var priceReduction float64

	if age < 3 {
		priceReduction = 80.0 / 100.0 // 80% off the original price of a used vehicle less than 3 years old.
	} else if age >= 3 && age < 10 {
		priceReduction = 70.0 / 100.0 // 70% off the original price of a used vehicle more than 3 years old but less than 10.
	} else if age >= 10 {
		priceReduction = 50.0 / 100.0 // 50% off the original price of a used vehicle 10 or more years old.
	}

	return originalPrice * priceReduction
}
