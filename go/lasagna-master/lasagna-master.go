package lasagna

// PreparationTime preparation time of lasagna
func PreparationTime(layers []string, avgTime int) int {

	numberOfLayers := len(layers)
	if avgTime <= 0 {
		return numberOfLayers * 2
	}
	return numberOfLayers * avgTime
}

// Quantities defines quantities recipe
func Quantities(layers []string) (int, float64) {
	var noodles, noodleQuantity int
	var sauces, sauceQuantity float64
	for _, v := range layers {
		if v == "noodles" {
			noodles += 1
			noodleQuantity = noodles * 50
		}
	}
	for _, v := range layers {
		if v == "sauce" {
			sauces += 1
			sauceQuantity = sauces * 0.2
		}
	}

	return noodleQuantity, sauceQuantity
}

func popLastItem(friendList []string) string {
	lastIndex := len(friendList) - 1
	return friendList[lastIndex]
}

// AddSecretIngredient adds secret ingredient to my list
func AddSecretIngredient(friendList, myList []string) []string {
	lastItem := popLastItem(friendList)
	myList = append(myList, lastItem)
	return myList
}

// ScaleRecipe scales the recipe
func ScaleRecipe(portionsAmt []float64, portions int) []float64 {
	// portionsAmt ----- 2
	// portionsAmt ----- x

	// 0.5portionsAmount --- 2
	// x -------------------portions  ---> X * 2 = 0.5 * portions --> x = (posrtionsAmt[i] * portions)/2
	var scaledRecipe []float64
	var result float64

	for _, v := range portionsAmt {
		result = (v * float64(portions)) / 2
		scaledRecipe = append(scaledRecipe, result)

	}
	return scaledRecipe
}
