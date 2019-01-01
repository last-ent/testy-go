package sublevel

import "testing"

func TestFile7(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test as expected.")
	} else {
		t.Error("This was supposed to be skipped!")
	}
}
