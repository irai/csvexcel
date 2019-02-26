package csvexcel

import (
	"testing"
)

func TestAddRowWithValues(t *testing.T) {

	table := New()
	row := table.AddRowWithValues([]string{"first_name", "last_name", "username"})
	table.SetHeader(row.Number)
	// fmt.Println("Header ", table.header, row.Number, table.Rows)

	for i := 0; i < 10; i++ {
		table.AddRow()
	}

	table.AddColumn()

	table.Rows[3].Cell("B").Value = "bob"
	table.Cell("d11").Value = "mary"

	if len(table.Columns) != 4 || len(table.Rows) != 11 {
		t.Error("wrong rows or column count", table.Columns, table.Rows)
		table.Print()
	}
	if r := table.FindRow("last_name", "bob"); r == nil || r.Number != 4 {
		t.Error("wrong last name", table.Columns, table.Rows)
		table.Print()
	}
	if r := table.FindRow("D", "mary"); r == nil || r.Number != 11 {
		t.Error("wrong value in D", r, table.Columns, table.Rows)
		table.Print()
	}

	// table.Print()
}

func TestFind(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`

	table, err := ParseCSV(in)
	if err != nil {
		t.Error("error parsing string ", err)
	}
	// table.Print()

	if v := table.Cell("A1").Value; v != "first_name" {
		t.Error("invalid value in A1 ", v)
	}
	if v := table.Cell("C3").Value; v != "ken" {
		t.Error("invalid value in C3 ", v)
	}
	if v := table.Cell("A0").Value; v != InvalidRange {
		t.Error("invalid value in A1 ", v)
	}
	if v := table.Cell("d4").Value; v != OutOfRange {
		t.Error("invalid value in A4 ", v)
	}

	table.SetHeader(1)

	if v := table.Cell("A1").Value; v != "first_name" {
		t.Error("invalid value in A1 ", v)
	}
	if v := table.Rows[2].Cell("username").Value; v != "ken" {
		t.Error("invalid value in username ", v)
	}
	if r := table.FindRow("last_name", "Pike"); r == nil || r.Cells[0].Value != "Rob" {
		t.Error("could not find Pike ", r)
	}

	table.Rows[2].Cell("username").Value = "bob"
	if r := table.FindRow("username", "bob"); r == nil || r.Cell("username").Value != "bob" {
		t.Error("could not find Bob ", r, r.Cell("username").Value)
	}
}
