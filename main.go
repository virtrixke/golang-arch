package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "P@ssw0rd"
	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	passOther := "P@sseuirgiosuh"

	err = comparePassword(passOther, hashedPass)
	if err != nil {
		log.Fatalln("NOT logged in")
	}

	log.Println("Logged in")

}

func hashPassword(pw string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("We got an error during compare pw: %w", err)
	}
	return nil
}
