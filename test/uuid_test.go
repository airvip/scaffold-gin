package test

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
)


func TestUuid(t *testing.T) {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1.String())

	// Parsing UUID from string input
	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}