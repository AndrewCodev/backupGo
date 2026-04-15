package speed

type Car struct {
	battery int
	speed int
	batteryDrain int
	distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery: 100,
		speed: speed,
		batteryDrain: batteryDrain,
		distance: 0,
	}
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}

	car.battery -= car.batteryDrain
	car.distance += car.speed
	return car
}

func CanFinish(car Car, track Track) bool {
	maxDistance := (car.battery / car.batteryDrain) * car.speed
	return maxDistance >= track.distance
}