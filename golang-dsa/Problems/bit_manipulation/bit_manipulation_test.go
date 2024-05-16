package bitmanipulation_test

import (
	"fmt"
	"testing"
)

func TestBitManipulation(t *testing.T) {
	a := uintptr(2)
	fmt.Printf("Bitwise NOT: %d - %.10b\n", ^a, ^a)
}
