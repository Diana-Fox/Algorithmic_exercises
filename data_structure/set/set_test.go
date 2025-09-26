package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)
	fmt.Println(set.Strings())
}
