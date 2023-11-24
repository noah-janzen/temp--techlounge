package calculator_test

import (
	"testing"

	"github.com/thenativeweb/hallowelt/calculator"
)

func TestAdd(t *testing.T) {
	sum := calculator.Add(21, 42)
	if sum != 63 {
		t.Errorf("Expected 63, got %d instead.", sum)
	}
}