package chainResponsibility

import "fmt"

type Reception struct {
	next department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Patient is registering")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (d *Reception) SetNext(next department) {
    d.next = next
}