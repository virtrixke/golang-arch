package main

import (
	"crypto"
	"crypto/hmac"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var key = []byte{}

func main() {

	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}
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

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(crypto.SHA512.New, key)

	_, err := h.Write((msg))
	if err != nil {
		return nil, err
	}

	signature := h.Sum((nil))
	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSig %w", err)
	}

	same := hmac.Equal(newSig, sig)
	return same, nil

}
