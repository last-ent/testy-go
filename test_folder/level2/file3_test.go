package level2

import "testing"

func TestFunc3(t *testing.T) {
	if file3Func() {
		t.Log("level2.file3Func works as expected")
	} else {
		t.Error("level2.file3Func returned false!")
	}
}
