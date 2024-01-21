package main

import (
	chainResponsibility "designPatterns/Patterns/COR"
	"designPatterns/Patterns/bridge"
	"fmt"
)

func main() {

	e := &bridge.Epson{}
	mac := &bridge.Mac{}
	mac.SetPrinter(e)
	mac.Print()
	//COR
	
	p := &chainResponsibility.Patient{
		RegistrationDone:  true,
		DoctorCheckupDone: false,
		MedicalCheck:      false,
	}
	m := &chainResponsibility.Medical{}
	d := &chainResponsibility.Doctor{}
	r := &chainResponsibility.Reception{}

	r.SetNext(d)
	d.SetNext(m)

	fmt.Println("test")
	r.Execute(p)

}
