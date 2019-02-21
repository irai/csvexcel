package csvexcel

import (
	"testing"
)

func TestColIndex(t *testing.T) {
	for i := 0; i < 100; i++ {
		nextColIndex(i)
		// log.Println("nextColIndex ", s)
	}
	if s := nextColIndex(100); s != "DA" {
		t.Error("unexpect index ", s)
	}
}
