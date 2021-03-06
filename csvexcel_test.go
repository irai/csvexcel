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

func TestErrors(t *testing.T) {
	table := New()
	table.Cell("B1").Value = "new name"
	if len(table.Errors.Rows()) != 1 {
		t.Error("failed to log error B1", len(table.Errors.Rows()))
		table.Errors.Print()
	}

	row := table.AddRowWithValues([]string{"cell A row 1", "cell B row 1"})
	table.Cell("A1").Value = "changed1"
	row.Cell("B").Value = "changed2"
	if len(table.Errors.Rows()) != 1 {
		t.Error("failed to log error A1 and B", len(table.Errors.Rows()))
		table.Errors.Print()
	}

	table.Cell("B3").Value = "invalid line"
	table.Cell("C3").Value = "invalid column"
	table.Cell("aa3445z").Value = "invalid"
	row.Cell("C").Value = "invalid column"

	if len(table.Errors.Rows()) != 5 {
		t.Error("failed to log error B1", len(table.Errors.Rows()))
		table.Errors.Print()
	}
	// table.Errors.Print()
}

func TestParseBigFile(t *testing.T) {
	table, err := ParseFile("testdata/FL_insurance_sample1.csv")
	if err != nil {
		t.Error("error parsing file ", err)
	}

	for _, c := range table.Columns {
		c.Hide = true
	}
	table.Columns[3].Hide = false

	if len(table.Rows()) != 36635 || table.Row(36635).Cell("O").Value != "-82.77459" {
		t.Error("error in line count  ", len(table.Rows()), table.Cell("O36635").Value)
	}
	// for i := 1; i < 1000; i++ {
	// log.Println("row i=", i, "value ", table.Rows()[i].Cell("A").Value)
	// }
}
