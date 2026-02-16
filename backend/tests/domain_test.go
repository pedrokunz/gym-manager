package tests

import "testing"

func TestCoreLogic(t *testing.T) {
	if 1+1 != 2 {
		t.Error("Math broke")
	}
}

func TestMemberLogic(t *testing.T) {
	name := "Alice"
	if len(name) <= 0 {
		t.Errorf("Name length is impossible")
	}
}
