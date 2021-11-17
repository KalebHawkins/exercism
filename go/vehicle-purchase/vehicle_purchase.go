package purchase

import "fmt"

// NeedsLicence determines whether a licence is need to drive a type of vehicle. Only "car" and "truck" require a licence.
func NeedsLicence(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}

	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	fmtString := "%v is clearly the better choice."
	if option1 < option2 {
		return fmt.Sprintf(fmtString, option1)
	}

	return fmt.Sprintf(fmtString, option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * .80
	} else if age > 3 && age < 10 {
		return originalPrice * .70
	}

	return originalPrice * .50
}
