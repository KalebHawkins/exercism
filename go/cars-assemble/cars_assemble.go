package cars

const carsPerHour = 221

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	switch speed {
	case 1, 2, 3, 4:
		return 1
	case 5, 6, 7, 8:
		return 0.9
	case 9, 10:
		return 0.77
	default:
		return 0
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(speed) * carsPerHour * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	if speed == 0 {
		return 0
	}

	rate := CalculateProductionRatePerHour(speed)

	if rate < 1000 {
		return rate
	}

	return limit
}
