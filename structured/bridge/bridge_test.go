package bridge

import "testing"

func TestBridge(t *testing.T) {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()

	winComputer := &windows{}
	winComputer.setPrinter(hpPrinter)
	winComputer.print()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
}
