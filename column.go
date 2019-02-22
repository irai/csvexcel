package csvexcel

import (
	"strconv"
	"strings"
)

type Column struct {
	pos   int
	Index string
	Name  string
	Hide  bool
}

var (
	InvalidColumn = Column{Index: "Invalid", Name: "Invalid"}
)

type Columns []*Column

func (t *table) AddColumn() *Column {

	c := Column{}
	c.pos = len(t.Columns)
	c.Index = nextColIndex(c.pos)
	c.Name = c.Index
	t.Columns = append(t.Columns, &c)

	for _, row := range t.Rows {
		cell := Cell{Row: row, Column: &c}
		row.addCell(&cell)
	}
	return &c
}

func nextColIndex(n int) string {
	b := []byte("ABCDEFGHIJKLMNOPQRSTYUVXZ")
	prefix := []byte{}
	if n >= len(b) {
		prefix = append(prefix, b[(n/len(b))-1])
		n = n % len(b)
	}
	return string(prefix) + string(b[n])
}

func (t *table) findColumn(index string) *Column {
	for _, c := range t.Columns {
		if c.Index == index {
			return c
		}
	}
	return nil
}

func toIndex(index string) (col string, row int) {

	if len(index) < 2 {
		return "", 0
	}

	s := strings.ToUpper(index)

	if s[0] <= 'A' && s[0] >= 'Z' {
		return "", 0
	}
	col = s[0:1]
	r := s[1:]

	if s[1] >= 'A' && s[1] <= 'Z' {
		col = s[0:2]
		if len(s) < 3 {
			return "", 0
		}
		r = s[2:]

	}

	var err error
	if row, err = strconv.Atoi(r); err != nil {
		return "", 0
	}
	return col, row
}
