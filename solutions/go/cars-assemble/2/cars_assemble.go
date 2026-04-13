package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * (successRate / 100.0)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

func CalculateCost(carsCount int) uint {
    groups := carsCount / 10
	remainder := carsCount % 10
    
	return uint(groups*95000 + remainder*10000)
}
