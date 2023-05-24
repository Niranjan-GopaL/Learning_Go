//Excuetables must be of package mainpackage main

package main
import "fmt"

// Every thing is pa`ssed by value except arrays, slices, maps and channels which some calls 
// reference types, these types are passed by reference ( they internally have pointers, so 
// no copying of the actual data happens when passing them) .

func main() {
	var g string = "Hello golang"
	println(g)
	fmt.Println(g)
}


func function() string {
	//Shorthand - Declration & Assignmnet
	a := 10
	b := "golang"
	println(a,b)
	return "Five of Diamonds"
}

// If variables/functions starts with Uppercase character, it is accessible
//  outside the scope of its package, if lowercase then it is only accessible 
// inside its package.