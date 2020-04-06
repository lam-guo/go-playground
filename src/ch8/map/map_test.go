package map_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = true
	n := 3
	if mySet[n] {
		t.Logf("%d exist", n)
	} else {
		t.Logf("%d not exist", n)
	}
	t.Log(len(mySet))
	delete(mySet, 1)
	if mySet[1] {
		t.Logf("%d exist", 1)
	} else {
		t.Logf("%d not exist", 1)
	}
}
