package chainResponsibility

import "fmt"

type Medical struct {
	next department
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicalCheck {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medicine is giving")
	p.MedicalCheck = true
	// m.next.Execute(p)
}

func (d *Medical) SetNext(next department) {
    d.next = next
}