package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {

	c := map[string]*Employee{}

	c["Lamine"] = &Employee{
		Name:   "Lamine",
		Number: 2,
		Hired:  time.Now(),
	}

	c["Matt"] = &Employee{
		Name:   "Matt",
		Number: 1,
		Boss:   c["Lamine"],
		Hired:  time.Now(),
	}

	fmt.Printf("%T %+[1]v\n", c["Lamine"])
	fmt.Printf("%T %+[1]v\n", c["Matt"])

}
