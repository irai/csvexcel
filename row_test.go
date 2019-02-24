package csvexcel

import (
	"testing"
)

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
	if v := table.Cell("A0").Value; v != InvalidRange.Value {
		t.Error("invalid value in A1 ", v)
	}
	if v := table.Cell("d4").Value; v != OutOfRange.Value {
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
