package iter

import (
	"testing"
)

func TestAccumulate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	for a := Accumulate(&nums); a.HasNext(); {
		t.Log(a.Next())
	}
}

func TestCount(t *testing.T) {
	for c := Count(1, 2); c.HasNext(); {
		result := c.Next()
		if result > 100 {
			break
		}
		t.Log(result)
	}
}

func TestCycle(t *testing.T) {
	for c := Cycle(&[]int{1, 2, 3, 4}, 2); c.HasNext(); {
		t.Log(c.Next())
	}
}
