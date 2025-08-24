package cars
import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate/60) * (successRate/100))
    /*
    var ratePerMin int = productionRate/60
    var successPercent int = successRate/100
    return ratePerMin * successPercent
    */
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const groupedCost int = 95000
    const individualCost int = 10000
	groupedCars := carsCount / 10
    individualCars := carsCount % 10
    cost := (groupedCars * groupedCost) + (individualCars * individualCost)
    fmt.Println("Individul cars",individualCars)
    fmt.Println("Grouped cars",groupedCars)
    return uint(cost)
}
