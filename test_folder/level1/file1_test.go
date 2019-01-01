package level1

import "testing"

func TestFunc1(t *testing.T) {
	if file1Func() {
		t.Log("level1.file1Func works as expected")
	} else {
		t.Error("level1.file1Func returned false!")
	}
}
