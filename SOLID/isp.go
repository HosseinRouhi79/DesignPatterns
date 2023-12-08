package main

import "fmt"

func main() {

}

type Machine interface {
	Print()
	Scan()
	Fax()
}

type MultiFunctionPrinter struct {
	//something
}

func (m MultiFunctionPrinter) Print() {
	fmt.Print("print")
}
func (m MultiFunctionPrinter) Scan() {
	fmt.Print("scan")
}
func (m MultiFunctionPrinter) Fax() {
	fmt.Print("fax")
}

// OldPrinter It breaks isp
type OldPrinter struct {
	//something
}

func (o OldPrinter) Print() {
	fmt.Print("print")
}

func (o OldPrinter) Scan() {
	fmt.Print("not valid")
}

func (o OldPrinter) Fax() {
	fmt.Print("not valid")
}

// Printer We have to break each large interface to small ones
type Printer interface {
	Print1()
}

type Scanner interface {
	Scan1()
}

type FaxMachine interface {
	Fax1()
}
