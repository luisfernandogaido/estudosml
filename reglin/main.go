package main

import (
	"fmt"
	"log"
)

func main() {
	modelo, err := NewModeloPropaganda("./Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", modelo.CovVendas())
}
