package helpers

import (
	"fmt"
	"strconv"
)

// PtrToHex converts uintptr to hex string.
func PtrToHex(ptr uintptr) string {
	s := fmt.Sprintf("%d", ptr)
	n, _ := strconv.Atoi(s)
	h := fmt.Sprintf("0x%x", n)
	return h
}
