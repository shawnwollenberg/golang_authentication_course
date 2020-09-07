package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "Shawn",
	}
	p2 := person{
		First: "Stephanie",
	}
	xp := []person{p1, p2}
	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))

	fmt.Println("PRINT JSON", string(bs))

	xp2 := []person{}
	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Back to a Go Data Structure", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)

}
func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Shawn",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}
func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data", err)
	}
	log.Println("New Person", p1)

}

/*
--git notes
git add -A
git commit -m "Encode Example"
git push
git tag v.0.2.0
git push --tags

curl notes
	curl localhost:8080/encode
	curl -XGET -H "Content-type: application/json" -d'{"First":"Shawn"}' 'localhost:8080/decode'

*/
