package codeAnalysis

type Institute struct {
	code SchoolCode
	name string
}

type School struct {
	referenceInstitute SchoolCode
	code               SchoolCode
	classification     Classification
	name               string
	emailAddress       EmailAddress
	pecAddress         EmailAddress
	location           Location
}

type SchoolCode string
type Classification struct {
	order string
	grade int
}
