package csvexcel

import (
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
func TestAddColumn(t *testing.T) {
	table := New()

	for i := 0; i < 6; i++ {
		table.AddColumn()
	}

	for i := 0; i < 10; i++ {
		table.AddRow()
	}

	table.AddColumn()

	if len(table.Columns) != 7 || len(table.Rows()) != 10 {
		t.Error("wrong rows or column count", table.Columns, table.rows)
	}
	// table.Print()
}

func TestFindColumn(t *testing.T) {
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
	table.SetHeader(1)

	if v := table.Row(2).Cell("first_name").Value; v != "Rob" {
		table.Print()
		t.Error("invalid value in first_name ", v)
	}
	if v := table.Row(1).Cell("e not valid"); v.Value != InvalidRange {
		t.Error("invalid column ", v)
	}

}
