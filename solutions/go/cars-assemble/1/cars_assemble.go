package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

    var percentage = successRate / float64(100)
    var carsPerHour = percentage * float64(productionRate)
    
    return carsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var carsPerMinute = CalculateWorkingCarsPerHour(productionRate, successRate)

    return int(carsPerMinute / float64(60))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

    var groupsCars = carsCount / 10
    var remainingCars = carsCount % 10

    var cost1 = groupsCars * 95000
    var cost2 = remainingCars * 10000

	return uint(cost1 + cost2)
}
