// Package
package main

// import "rsc.io/quote"
import (
	"fmt"
	"rsc.io/quote"
)

// function for å hente fra rsc.io/quote GLASS
func main() {
	fmt.Print(quote.Glass())

}
