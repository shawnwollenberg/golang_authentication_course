package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
}

func test64Base() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
}

func hashPassword(pwd string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

//separate function for comparing
func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password %w", err)
	}
	return nil
}

func main() {
	test64Base()
	pass := "12345678"
	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}
	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}
	log.Println("Logged In!")
	/*p1 := person{
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
	*/

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
git tag v0.5.0
git push --tags

//clean version
git stash
git stash drop
git stash pull

curl notes
	https://curlbuilder.com
	curl localhost:8080/encode
	curl -XGET -H "Content-type: application/json" -d'{"First":"Shawn"}' 'localhost:8080/decode'

*/
