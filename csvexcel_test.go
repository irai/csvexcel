package csvexcel

import (
	"os"
	"testing"
)

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
	// table.Print()

	if v := table.Cell("A1"); v.Value != "first_name" {
		table.Print()
		t.Error("invalid value in A1 ", v)
	}
	if v := table.Cell("C3").Value; v != "ken" {
		t.Error("invalid value in C3 ", v)
	}
	if v := table.Cell("A0").Value; v != InvalidRange {
		t.Error("invalid value in A0 ", v)
	}
	if v := table.Cell("d4").Value; v != OutOfRange {
		t.Error("invalid value in A4 ", v)
	}

	if r := table.FindRow("b", "Pike"); r == nil || r.Cells[0].Value != "Rob" {
		t.Error("could not find Pike ", r)
	}
}

func TestParseFile(t *testing.T) {
	table, err := ParseFile("testdata/list.csv")
	if err != nil {
		t.Error("error parsing file ", err)
	}

	for _, c := range table.Columns {
		c.Hide = true
	}
	table.Columns[3].Hide = false
	// table.Print()
	// for _, c := range table.Columns {
	// log.Println(*c)
	// }

}

func TestSave(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`

	const filename = "./csvtest.out"
	defer os.Remove(filename)

	t1, err := ParseCSV(in)
	if err != nil {
		t.Error("error parsing string ", err)
	}

	if err := t1.Save(filename); err != nil {
		t.Error("error parsing string ", err)
	}

	t2, err := ParseFile(filename)
	if err != nil {
		t.Error("error loading file ", err)
	}
	// t2.Print()

	for i := range t1.Rows() {
		for x := range t1.Columns {
			if t1.Rows()[i].Cells[x].Value != t2.Rows()[i].Cells[x].Value {
				t.Error("cells dont match ", i, x, t1.Rows()[i].Cells[x].Value, t2.Rows()[i].Cells[x].Value)
			}
		}
	}
}
