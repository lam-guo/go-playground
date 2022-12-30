package algorithm

import "testing"

func TestSubSet(t *testing.T) {
	backends := []string{"A", "B", "C", "D", "E", "F", "G"}
	clientId := 10
	subsetSize := 3
	t.Log(SubSet(backends, clientId, subsetSize))
	t.Log(SubSet(backends, clientId, subsetSize))
	t.Log(SubSet(backends, clientId, subsetSize))
	t.Log(SubSet(backends, clientId, subsetSize))
}
