package bridge

import (
	"fmt"
	"reflect"
)

type Mac struct {
	printer Printer
}

func (mac *Mac) SetPrinter(p Printer) {
	mac.printer = p
}

func (mac *Mac) Print() {
	fmt.Println(reflect.TypeOf(mac.printer))
	mac.printer.PrintFile()
}
