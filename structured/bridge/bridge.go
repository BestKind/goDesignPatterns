package bridge

import "fmt"

type printer interface {
	printFile()
}

type epson struct {
}

func (p *epson) printFile() {
	fmt.Println("Printing by a epson printer")
}

type hp struct {
}

func (h *hp) printFile() {
	fmt.Println("Printing by a hp printer")
}

type computer interface {
	print()
	setPrinter(printer)
}

type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}
