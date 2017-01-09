package slices

import "testing"

func TestContains(t *testing.T) {
	s := []int{1, 2, 3}

	if !Contains(s, 1) {
		t.Error("not contains 1")
	}

	if Contains(s, 5) {
		t.Error("contains 5")
	}
}

type item struct {
	id    string
	token string
}

func TestContainsFunc(t *testing.T) {

	m := []item{{id: "id1", token: "token1"}, {id: "id2", token: "token2"}}

	if !ContainsFunc(m, item{id: "id1"}, func(v1 interface{}, v2 interface{}) bool {
		a1, _ := v1.(item)
		a2, _ := v2.(item)
		return a1.id == a2.id
	}) {
		t.Error("contains func error")
	}

	if ContainsFunc(m, item{id: "id3"}, func(v1 interface{}, v2 interface{}) bool {
		a1, _ := v1.(item)
		a2, _ := v2.(item)
		return a1.id == a2.id
	}) {
		t.Error("contains func error")
	}

}
