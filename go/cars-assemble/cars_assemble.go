package cars

const carsEachHour = 221

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {

	result := carsEachHour * float64(speed)
	if speed >= 5 && speed <= 8 {
		result *= successRate(speed)
	}
	if speed >= 9 && speed <= 10 {
		result *= successRate(speed)
	}

	return result
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	speedFloat := float64(speed)
	carEachMinutes := carsEachHour / 60.0
	result := carEachMinutes * speedFloat
	if speed >= 5 && speed <= 8 {
		result *= successRate(speed)
	}
	if speed >= 9 && speed <= 10 {
		result *= successRate(speed)
	}

	return int(result)

}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	if speed == 0 {
		return 0.0
	}
	if speed >= 1 && speed <= 4 {
		return 1
	}
	if speed >= 5 && speed <= 8 {
		return 0.9
	}
	if speed >= 9 && speed <= 10 {
		return 0.77
	} else {
		return 0.77
	}
}
