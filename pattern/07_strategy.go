package pattern

import "fmt"

// IAction (strategy) interface and it's concrete implmentations
type IAction interface {
	Do(string)
}

type Boil struct {
	// state and stuff...
}

func (a *Boil) Do(sbj string) { fmt.Printf("Boiling %s...\n", sbj) }

type Fry struct {
	// state and stuff...
}

func (a *Fry) Do(sbj string) { fmt.Printf("Frying %s...\n", sbj) }

// Cook (context)
type CookingMachine struct {
	action IAction
}

func (cm *CookingMachine) SetAction(action IAction) (self *CookingMachine) {
	cm.action = action
	return cm
}

func (cm *CookingMachine) Execute(sbj string) {
	cm.action.Do(sbj)
}

// Demonstration
func init() {
	fmt.Println("Demonstrating Strategy pattern:")
	defer fmt.Println()

	cm := CookingMachine{}

	fry := &Fry{}
	boil := &Boil{}

	sbj := "egg"

	cm.SetAction(fry).Execute(sbj)
	cm.SetAction(boil).Execute(sbj)
}
