package codeAnalysis

type User struct {
	id UserCode
	firstname string
	lastname string
	email EmailAddress
	password Password
	gender string
	birthday Date
}

type Student struct {
	User
	subcriptionYear Year
	graduationYear Year
}

type Teacher struct {
	User
}

type ATAStaff struct {
	User
}

type UserCode string