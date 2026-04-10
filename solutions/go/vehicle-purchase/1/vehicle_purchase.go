package purchase

import (
	"fmt"
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck"{
        return true
    }else{
        return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	lista := []string{option1, option2}
    sort.Strings(lista)
    
    return fmt.Sprintf("%s is clearly the better choice.", lista[0])
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

    var totalPrice float64
	if age >= 10 {
        totalPrice = originalPrice * 50 / 100
    }else if age >= 3 && age < 10 {
        totalPrice = originalPrice * 70 / 100
    }else if age >= 0 && age < 3 {
        totalPrice = originalPrice * 80 / 100
    }else{
        totalPrice = 0.0
    }

    return totalPrice
}
