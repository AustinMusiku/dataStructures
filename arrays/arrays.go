package arrays

import "fmt"

const arraySize = 5

func Array(){
	// array of 5 numbers
	var nums [arraySize]int
	fmt.Println(nums)

	// array of strings
	// shorthand
	names := [arraySize] string {"apple","facebook","tesla","google","microsoft"}
	
	// loop through array and print each
	for i:=arraySize; i > 0; i-- {
		fmt.Printf("%s\n", names[i-1])
	}
}