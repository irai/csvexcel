package csvexcel

import (
	"os"
	"testing"
)

func TestColIndex(t *testing.T) {
	for k := 0; k < 100; k++ {
		s := pos2col(k)
		// t.Error("i ", s, k)
		if x := col2pos(s); x != k {
			t.Error("unexpect str2Pos ", s, x, k)
		}
		// log.Println("nextColIndex ", s)
	}
	if s := pos2col(100); s != "CW" {
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
	// table.Print()
}
func TestIndex(t *testing.T) {
	if col, row := split2colnumber("A1"); col != "A" || row != 1 {
		t.Error("failed to index A1", col, row)
	}
	if col, row := split2colnumber("az1"); col != "AZ" || row != 1 {
		t.Error("failed to index az1", col, row)
	}
	if col, row := split2colnumber(" "); col != "" || row != -1 {
		t.Error("failed to index space", col, row)
	}
	if col, row := split2colnumber(""); col != "" || row != -1 {
		t.Error("failed to index empty", col, row)
	}
	if col, row := split2colnumber("A"); col != "" || row != -1 {
		t.Error("failed to index A", col, row)
	}
	if col, row := split2colnumber("AAA1"); col != "" || row != -1 {
		t.Error("failed to index AAA1", col, row)
	}
	if col, row := split2colnumber("1"); col != "" || row != -1 {
		t.Error("failed to index 1", col, row)
	}
	if col, row := split2colnumber("X123456789"); col != "X" || row != 123456789 {
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

	if r := table.FindRow("b", "Pike"); r == nil || r.Cells[0].Value != "Rob" {
		t.Error("could not find Pike ", r)
	}
}

func TestParseFile(t *testing.T) {
	table, err := ParseFile("/mnt/c/Users/fabio/Desktop/TEST Member Database.csv")
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

	for i := range t1.Rows {
		for x := range t1.Columns {
			if t1.Rows[i].Cells[x].Value != t2.Rows[i].Cells[x].Value {
				t.Error("cells dont match ", i, x, t1.Rows[i].Cells[x].Value, t2.Rows[i].Cells[x].Value)
			}
		}
	}
}
