package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(elapsedTime int) int {
	return OvenTime - elapsedTime
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers int) int {
	return layers * 2
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(layers int, elapsedTime int) int {
	return PreparationTime(layers) + elapsedTime
}
