package speed

// Car defines a Car struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
	return car
}

// Track is a struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{distance: distance}
	return track
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		return car
	}

	car.battery = car.battery - car.batteryDrain
	car.distance = car.speed
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	drained := car.distance * car.batteryDrain
	if drained > track.distance {
		return true
	}
	return false

}
