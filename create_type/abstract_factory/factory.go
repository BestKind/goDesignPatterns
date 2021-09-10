package abstract_factory

import "fmt"

type Door interface {
	GetDescription()
}

type WoodenDoor struct {
}

type IronDoor struct {
}

func (wd *WoodenDoor) GetDescription() {
	fmt.Println("I'm an wooden door")
}

func (id *IronDoor) GetDescription() {
	fmt.Println("I'm an iron door")
}

type DoorFittingExpert interface {
	GetDesc()
}

type Carpenter struct {
}

type Welder struct {
}

func (c *Carpenter) GetDesc() {
	fmt.Println("I can only fit wooden door")
}

func (w *Welder) GetDesc() {
	fmt.Println("I can only fit iron door")
}

type DoorFactory interface {
	MakeDoor()
	MakeFittingExpert()
}

type WoodenDoorFactory struct {
}

type IronDoorFactory struct {
}

func (wdf *WoodenDoorFactory) MakeDoor() *WoodenDoor {
	return &WoodenDoor{}
}

func (wdf *WoodenDoorFactory) MakeFittingExpert() *Carpenter {
	return &Carpenter{}
}

func (idf *IronDoorFactory) MakeDoor() *IronDoor {
	return &IronDoor{}
}

func (idf *IronDoorFactory) MakeFittingExpert() *Welder {
	return &Welder{}
}
