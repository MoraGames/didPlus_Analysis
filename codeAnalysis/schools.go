package codeAnalysis

type Institute struct {
	code schoolCode
	name string
}

type School struct {
	referenceInstitute schoolCode
	code               schoolCode
	classification     Classification
	name               string
	emailAddress       emailAddress
	pecAddress         emailAddress
	location           Location
}

type schoolCode string
type Classification struct {
	order string
	grade int
}
