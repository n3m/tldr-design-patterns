package main

import (
	"golang/logistics"
)

/*
	CONTEXT:
 In a Factory Method pattern, the whole point is to be able to generate instances with different
 properties that are property of the specific type, but also the ones that are from the parent.

 In our case, we're attepting to generate a new transport, and specifically we want a bus, so at
 the end, our variable type will be a Bus of some sorts
*/

/*
 SOLUTION:
 Go to file > logistics/main.go for further explanation
*/

func main() {

	// We're creating a new transport, and we're passing the type of transport we want to generate
	someTransport := logistics.CreateNewTransport("bus")

	/* bla bla bla */

}
