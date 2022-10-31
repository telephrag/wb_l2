package pattern

import "fmt"

// Command interface and concrete implementation
type Command interface {
	Do()
}

// Increase gas pressure
type IncreaseGasPressureCommand struct {
	furnance *Furnance
}

func (c *IncreaseGasPressureCommand) Init(furnance *Furnance) (self *IncreaseGasPressureCommand) {
	c.furnance = furnance
	return c
}

func (c *IncreaseGasPressureCommand) Do() {
	if c.furnance.GetGasPresure()+10 > 100 {
		fmt.Println("can't increase pressure even more")
		return
	}
	c.furnance.setGasPressure(c.furnance.GetGasPresure() + 10)
}

// Decrease gas pressure
type DecreaseGasPressureCommand struct {
	furnance *Furnance
}

func (c *DecreaseGasPressureCommand) Init(furnance *Furnance) (self *DecreaseGasPressureCommand) {
	c.furnance = furnance
	return c
}

func (c *DecreaseGasPressureCommand) Do() {
	if c.furnance.GetGasPresure() == 0 {
		fmt.Println("can't decrease pressure even more")
		return
	}

	if c.furnance.GetGasPresure()-10 < 0 {
		c.furnance.setGasPressure(0)
		c.furnance.setBurn(false)
		fmt.Println("Furnance has been extinguished.")
		return
	}

	c.furnance.setGasPressure(c.furnance.GetGasPresure() - 10)
}

// Ignite
type IgniteCommand struct {
	furnance *Furnance
}

func (c *IgniteCommand) Init(furnance *Furnance) (self *IgniteCommand) {
	c.furnance = furnance
	return c
}

func (c *IgniteCommand) Do() {
	if c.furnance.GetGasPresure() > 0 {
		c.furnance.setBurn(true)
		fmt.Println("furnance is ignited")
	} else {
		fmt.Println("can't ignite, no gas is coming out")
	}
}

// Object the commands are executed over
type Furnance struct {
	gasPressure int
	isBurning   bool
}

func (f *Furnance) GetGasPresure() int { return f.gasPressure }

func (f *Furnance) setGasPressure(gasPressure int) { f.gasPressure = gasPressure }

func (f *Furnance) setBurn(status bool) { f.isBurning = status }

// Demonstration
func init() {
	fmt.Println("Demonstrating Command pattern:")
	defer fmt.Println()

	f := &Furnance{}

	incGasCmd := (&IncreaseGasPressureCommand{}).Init(f)
	decGasCmd := (&DecreaseGasPressureCommand{}).Init(f)
	igniteCmd := (&IgniteCommand{}).Init(f)

	// Do I need a seperate struct for furnance's control panel?
	igniteCmd.Do() // cant ignite
	decGasCmd.Do() // cant decrease

	incGasCmd.Do() // increase

	igniteCmd.Do() // ignite

	decGasCmd.Do() // extinguish

	incGasCmd.Do()
	incGasCmd.Do()
	incGasCmd.Do()
	incGasCmd.Do()
	incGasCmd.Do()
	incGasCmd.Do()
	incGasCmd.Do()
	incGasCmd.Do()
	incGasCmd.Do()
	incGasCmd.Do()
	incGasCmd.Do() // cant increase
}
