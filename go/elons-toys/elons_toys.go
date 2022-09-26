package elon

import "fmt"

//Drive updates the number of meters driven based on the car's speed,
//and reduces the battery according to the battery drainage
func (c *Car) Drive()  {
	if c.battery - c.batteryDrain >= 0 {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

//DisplayDistance shows distance travelled
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", c.distance)
}

//DisplayBattery shows remaining battery life
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", c.battery)
}

//CanFinish determines if the car can finish a race given distance and current battery state
func (c *Car) CanFinish(trackDistance int) bool {
	return trackDistance / c.speed * c.batteryDrain <= c.battery
}
