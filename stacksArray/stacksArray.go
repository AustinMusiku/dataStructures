package stacksArray

import "fmt"

const stackSize = 5
var stack [] int
var top int

// isEmpty
func IsEmpty() bool{
	if(len(stack) == 0){
		return true
	}else{
		return false
	}
}
// isFull
func IsFull() bool{
	if(len(stack) < stackSize){
		return false
	}else{
		return true
	}
}
// push
// pop
// peek