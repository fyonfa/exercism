package lasagna

// TODO: define the 'OvenTime()' function

func OvenTime() int {
	return 40
}

// TODO: define the 'RemainingOvenTime()' function

func RemainingOvenTime(timeInOven int) int {
	cookingTime := OvenTime()
	return cookingTime - timeInOven
}

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers int) int {
	layerCookingTime := 2
	return layerCookingTime * layers
}

// TODO: define the 'ElapsedTime()' function

func ElapsedTime(layers, timeInOven int) int {
	preparationTime := PreparationTime(layers)
	return preparationTime + timeInOven
}
