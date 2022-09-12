package main

import (
	"awesomeProject/center"
	"fmt"
)

func main() {
	var good, bad int
	numTest := 100
	var res bool
	for i := 0; i < numTest; i++ {
		res = center.GeneratedKeys()
		fmt.Println(i, " : ", res)
		if res {
			good++
		} else {
			bad++
		}
	}
	fmt.Println("Good: ", float64(good)/float64(numTest))
	fmt.Print("Bad: ", float64(bad)/float64(numTest))

	//center.GeneratedKeys()

	// maybe i should to realize function to find Primal number and then this can work normal
	// but maybe not )))
}
