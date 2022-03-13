package main

type Document struct {
}

// this is ok for interface definition, but better to break more granular control
// so that you are not force to implement all of interface for your struct

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// above one interface have too many functions it's ok if you need a multifunction device
// which confirm have Print, Fax, Scan capability

type MultiFunctionPrinter struct {
	// ...some fields
}

func (m MultiFunctionPrinter) Print(d Document) {}

func (m MultiFunctionPrinter) Fax(d Document) {}

func (m MultiFunctionPrinter) Scan(d Document) {}

// but let say you have a old printer which only can print, but you are forcing to implement all of function

type OldFashionedPrinter struct {
	// ...
}

func (o OldFashionedPrinter) Print(d Document) {
	// ok
}

// for un support function you have to return some info

func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Or you add Deprecated so consumer wont use it

func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// Better approach: split into several interfaces

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// printer only
type MyPrinter struct {
	// ...some fields
}

func (m MyPrinter) Print(d Document) {
	// ...
}

// combine interfaces
type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	//
}

func (p Photocopier) Print(d Document) {
	//
}

// or you can also nested two interfaces together
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// interface combination pattern
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

// decorator design pattern: since we already have printer and scanner implemented as seprate components
func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

// decorator design pattern: since we already have printer and scanner implemented as seprate components
func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {

}
