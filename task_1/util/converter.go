package util

import (
	"fmt"
	// "reflect"
)

// keyConv returns the string representation of i
func keyConv(i interface{}) string {
	return fmt.Sprintf("%v", i)
}