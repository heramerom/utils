package maps

import "testing"

func TestDefault(t *testing.T) {

	m := map[string]int{"one": 1, "two": 2}

	if Default(m, "one", 10) != 1 {
		t.Error("error")
	}

	if Default(m, "three", 3) != 3 {
		t.Error("error")
	}

}
