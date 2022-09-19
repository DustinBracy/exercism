package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	const message = " is clearly the better choice."
	vehicle := option1
	if option1 > option2 {
		vehicle = option2
	}
	return vehicle + message	
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resellPrice float64
	if age < 3 {
		resellPrice = originalPrice * .80
	} else if age >= 10 {
		resellPrice = originalPrice * .5
	} else {
		resellPrice = originalPrice * .70
	}
	return resellPrice
}

