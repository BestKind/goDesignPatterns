package chain

import "testing"

func TestChain(t *testing.T) {
	cashier := &cashier{}

	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name:"123"}

	reception.execute(patient)
}
