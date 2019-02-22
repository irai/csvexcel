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

func TestAddColumns(t *testing.T) {
	table := New()

	for i := 0; i < 6; i++ {
		table.AddColumn()
	}

	for i := 0; i < 10; i++ {
		table.AddRow()
	}

	table.AddColumn()

	if len(table.Columns) != 7 || len(table.Rows) != 10 {
		t.Error("wrong rows or column count", table.Columns, table.Rows)
	}
	table.Print()
}
