package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string `json:"firstname"`
}

func main() {
	p1 := person{
		First: "vincent",
	}

	p2 := person{
		First: "barbara",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)

	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Print JSON", string(bs))

	xp2 := []person{}

	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Print unmarshall", xp2)

	http.HandleFunc("/encode", encodeJSON)
	http.HandleFunc("/decode", decodeJSON)
	http.ListenAndServe(":8080", nil)
}

func encodeJSON(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "vincent",
	}

	p2 := person{
		First: "barbara",
	}

	xp := []person{p1, p2}

	err := json.NewEncoder(w).Encode(xp)
	if err != nil {
		log.Println("Encode failed", err)
	}

}

func decodeJSON(w http.ResponseWriter, r *http.Request) {
	var xp []person
	err := json.NewDecoder(r.Body).Decode(&xp)
	if err != nil {
		log.Println("Decode failed", err)
	}

	log.Println("Person:", xp)
}
