package chainResponsibility

type department interface {
	SetNext(next department)
	Execute(patient *Patient)
}
