package abstract_factory

import "testing"

func TestFactory(t *testing.T) {
	woodenFactory := WoodenDoorFactory{}
	wDoor := woodenFactory.MakeDoor()
	wExpert := woodenFactory.MakeFittingExpert()

	wDoor.GetDescription()
	wExpert.GetDesc()

	ironFactory := IronDoorFactory{}
	iDoor := ironFactory.MakeDoor()
	iExpert := ironFactory.MakeFittingExpert()

	iDoor.GetDescription()
	iExpert.GetDesc()
}
