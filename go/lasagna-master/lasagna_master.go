package lasagna

// PreparationTime() returns the estimated total preparation time based on the number of layers.
// If the avgPrepTime passed is 0 the assumed prep time is 2 minutes.
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		return len(layers) * 2
	}

	return len(layers) * avgPrepTime
}

// Quantities returns the amount of noodles and sauce needed for each layer.
func Quantities(layers []string) (noodles int, sauce float64) {
	const gramsPerNoodleLayer = 50
	const litersPerSauceLayer = 0.2

	var totalNoodles int
	var totalSauce float64

	for _, v := range layers {
		switch v {
		case "noodles":
			totalNoodles += gramsPerNoodleLayer
		case "sauce":
			totalSauce += litersPerSauceLayer
		}
	}

	return totalNoodles, totalSauce
}

func AddSecretIngredient(friendsIngredients []string, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendsIngredients[len(friendsIngredients)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities = make([]float64, len(quantities))
	copy(scaledQuantities, quantities)

	for i, v := range scaledQuantities {
		scaledQuantities[i] = v * float64(portions) / 2
	}

	return scaledQuantities
}
