package hangul

import (
	"testing"
)

func ExampleHasConsonantSuffix(t *testing.T) {
	if HasConsonanatSuffix("Go 언어") {
		t.Error("No")
	}
	// Output:
	// false
}
