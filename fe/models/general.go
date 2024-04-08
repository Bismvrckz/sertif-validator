package models

type ToeflCertificate struct {
	ID            int
	TestID        string
	Name          string
	StudentNumber string
	Major         string
	DateOfTest    string
	ToeflScore    string
	InsertDate    string
}
