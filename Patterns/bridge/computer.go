package bridge

type Computer interface {
	SetPrinter(Printer)
	Print()
}
