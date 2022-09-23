package stacksArray

import (
	"errors"
)

const stackSize = 5
var stack [] int
var top int = 0

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
func Push(x int) error {
	if(IsFull()){
		return errors.New("stack is full")
	}else{
		if(!IsEmpty()){
			top++
		}
		stack = append(stack, x)
	}
	return nil
}
// pop
func Pop() (int, error){
	var popped int
	if(IsEmpty()){
		return 0, errors.New("stack is empty")
	}else{
		popped = stack[top]
		stack[top] = 0
		stack = stack[:top]
		top--
	}
	return popped, nil
}
// peek
func Peek() (int, error){
	if(IsEmpty()){
		return 0, errors.New("stack is empty")
	}else{
		return stack[top], nil
	}
}