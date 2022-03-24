package codeAnalysis

import (
	"crypto/rand"
	"golang.org/x/crypto/argon2"
	"log"
)

type Password struct {
	salt []byte
	hash []byte
}

func NewPassword (password []byte)(Password){
	//Generate the salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		log.Panicln(err)
	}

	//Generate the hash
	hash := argon2.IDKey(password, salt, 1, 64*1024, 4, 64)

	//Generate the password
	return Password{salt, hash}
}