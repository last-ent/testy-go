package sublevel

import "testing"

func TestFunc4(t *testing.T) {
	if file4Func() {
		t.Log("level2sub.file4Func works as expected")
	} else {
		t.Error("level2sub.file4Func returned false!")
	}
}
