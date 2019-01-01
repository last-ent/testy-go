package sublevel2

import "testing"

func TestFunc8(t *testing.T) {
	if file8Func() {
		t.Log("level3sub2.file8Func works as expected")
	} else {
		t.Error("level3sub2.file8Func returned false!")
	}
}
