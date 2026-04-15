package purchase

import "fmt"

func NeedsLicense(kind string) bool {
    return kind == "car" || kind == "truck"
}

func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
    	return fmt.Sprintf("%s is clearly the better choice.", option1)
    }
    return fmt.Sprintf("%s is clearly the better choice.", option2)
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
        case age >= 10:
            return originalPrice * 0.5
        case age >= 3:
            return originalPrice * 0.7
        default:
            return originalPrice * 0.8
    }
    return originalPrice
}
