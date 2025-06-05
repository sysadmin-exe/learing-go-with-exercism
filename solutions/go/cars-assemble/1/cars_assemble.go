package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    floatProductionRate := float64(productionRate)
	successfulCars := successRate / float64(100) * floatProductionRate
    return successfulCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    floatProductionRate := float64(productionRate)
	successfulCars := successRate / float64(100) * floatProductionRate
    return int(successfulCars)/60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    costGroupOfTen := carsCount/10 * 95000
    remainingCarCount := carsCount%10 * 10000
	return uint(costGroupOfTen + remainingCarCount)
}
