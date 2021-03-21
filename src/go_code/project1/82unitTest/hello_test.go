package hello

import "testing"

func TestAddUpper(t *testing.T) {
	wantInt := 11
	r := AddUpper(wantInt)
	if r != 11 {
		t.Fatalf("期望的结果%v,实际结果%v", wantInt, r)
	}
}

func TestMinusUpper(t *testing.T) {
	wantInt := 10
	r := AddUpper(wantInt)
	if r != 11 {
		t.Fatalf("期望的结果%v,实际结果%v", wantInt, r)
	}
}
