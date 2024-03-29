package elon

import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	c.battery = c.battery - c.batteryDrain
	c.distance = c.speed - c.distance
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", c.distance)
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", c.battery)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) CanFinish(trackDistance int) bool {
	return c.battery/c.batteryDrain*c.speed >= trackDistance

}
