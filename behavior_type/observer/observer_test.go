package observer

import "testing"

func TestObserver(t *testing.T) {
	shirtItem := newItem("Shirt")

	observerFirst := &customer{id: "ttt"}
	observerSecond := &customer{id: "fff"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
