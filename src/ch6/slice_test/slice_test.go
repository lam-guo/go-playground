package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
	s1 := []int{1, 2, 3}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShare(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug"}
	Q2 := year[2:4]
	t.Log(Q2, len(Q2), cap(Q2))
	Q3 := year[3:6]
	t.Log(Q3, len(Q3), cap(Q3))
	Q2[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	// t.Log(a == b)
	t.Log(&a == &b)
}
