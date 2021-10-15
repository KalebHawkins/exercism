package savings

import "math"

const interestNegativePercent float32 = -3.213
const interestPercentLessThan1000 float32 = 0.5
const interestPercentBetween1000And5000 float32 = 1.621
const interestPercent5000 float32 = 2.475

// InterestRate calculates the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return interestNegativePercent
	} else if balance < 1000 && balance >= 0 {
		return interestPercentLessThan1000
	} else if balance >= 1000 && balance < 5000 {
		return interestPercentBetween1000And5000
	} else {
		return interestPercent5000
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * math.Abs(float64(InterestRate(balance))) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	if balance > targetBalance {
		return 0
	}

	return YearsBeforeDesiredBalance(AnnualBalanceUpdate(balance), targetBalance) + 1
}
