package observer

import "fmt"

type observer interface {
	update(string)
	getID() string
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s \n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}

type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}

type item struct {
	observerList []observer
	name         string
	isStock      bool
}

func newItem(name string) *item {
	return &item{name: name}
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, ob := range i.observerList {
		ob.update(i.name)
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock \n", i.name)
	i.isStock = true
	i.notifyAll()
}

func removeFromSlice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
