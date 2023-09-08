package mascot_test

import (
	"testing"

	"github.com/bsmith578/4143-PLC/tree/main/Assignments/P01/mascot"
)

// Checks if BestMascot() prints "Go Gopher"
func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot")
	}
}
