package csvexcel

import (
	"testing"
)

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

	if v := table.Rows[1].Cell("first_name").Value; v != "Rob" {
		t.Error("invalid value in first_name ", v)
	}
	if v := table.Rows[1].Cell("e not valid").Value; v != InvalidRange.Value {
		t.Error("invalid column ", v)
	}

}
