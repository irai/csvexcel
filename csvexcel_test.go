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
func TestIndex(t *testing.T) {
	if col, row := toIndex("A1"); col != "A" || row != 1 {
		t.Error("failed to index A1", col, row)
	}
	if col, row := toIndex("az1"); col != "AZ" || row != 1 {
		t.Error("failed to index az1", col, row)
	}
	if col, row := toIndex(" "); col != "" || row != 0 {
		t.Error("failed to index space", col, row)
	}
	if col, row := toIndex(""); col != "" || row != 0 {
		t.Error("failed to index empty", col, row)
	}
	if col, row := toIndex("A"); col != "" || row != 0 {
		t.Error("failed to index A", col, row)
	}
	if col, row := toIndex("AAA1"); col != "" || row != 0 {
		t.Error("failed to index AAA1", col, row)
	}
	if col, row := toIndex("1"); col != "" || row != 0 {
		t.Error("failed to index 1", col, row)
	}
	if col, row := toIndex("X123456789"); col != "X" || row != 123456789 {
		t.Error("failed to index X123456789", col, row)
	}
}

func TestParseCSV(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`

	table, err := ParseCSV(in)
	if err != nil {
		t.Error("error parsing string ", err)
	}
	table.Print()

	if v := table.Cell("A1").Value; v != "first_name" {
		t.Error("invalid value in A1 ", v)
	}
	if v := table.Cell("C3").Value; v != "ken" {
		t.Error("invalid value in C3 ", v)
	}
	if v := table.Cell("A0").Value; v != InvalidRange.Value {
		t.Error("invalid value in A1 ", v)
	}
	if v := table.Cell("d4").Value; v != OutOfRange.Value {
		t.Error("invalid value in A4 ", v)
	}

	if r := table.FindRow("b", "Pike"); r == nil || r.Cells[0].Value != "Rob" {
		t.Error("could not find Pike ", r)
	}
}

func TestOpen(t *testing.T) {
	table, err := Open("/mnt/c/Users/fabio/Desktop/TEST Member Database.csv")
	if err != nil {
		t.Error("error parsing file ", err)
	}

	for _, c := range table.Columns {
		c.Hide = true
	}
	table.Columns[3].Hide = false
	table.Print()
	// for _, c := range table.Columns {
	// log.Println(*c)
	// }

}
