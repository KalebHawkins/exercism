package elon

import "fmt"

func (c *Car) Drive() {
	c.distance += c.speed
	c.battery -= c.batteryDrain
}

func (c *Car) CanFinish(trackDistance int) bool {
	return c.battery-trackDistance/c.speed*c.batteryDrain >= 0
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", c.battery)
}
