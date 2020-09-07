package main

import (
	"encoding/json"
	"fmt"
	"log"
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
}

/*
--git notes
git push
git tag v.0.2.0
git push --tags
*/
