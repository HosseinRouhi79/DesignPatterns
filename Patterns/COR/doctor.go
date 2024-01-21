package chainResponsibility

import "fmt"

type Doctor struct {
	next department
}

func (r *Doctor) Execute(p *Patient) {
	if p.DoctorCheckupDone {
		fmt.Println("Doctor checkup already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("doctor is visiting patient")
	p.DoctorCheckupDone = true
	r.next.Execute(p)
}

func (d *Doctor) SetNext(next department) {
    d.next = next
}