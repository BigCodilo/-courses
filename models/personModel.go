package models

type Person struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	Gender       string
	GenderIota   Gender
	RegisterDate Date
	Loan         string
}

const (
	Male = iota
	Female
)
