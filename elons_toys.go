package elon

import "fmt"

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func (car *Car) Drive() {
	if car.battery-car.batteryDrain > 0 {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// DisplayDistance returns the distance as displayed on the LED display
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery returns the battery percentage as displayed on the LED display
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if a car is able to finish a certain track distance.
func (car *Car) CanFinish(trackDistance int) bool {
	return trackDistance <= car.battery/car.batteryDrain*car.speed
}
