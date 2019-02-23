package csvexcel

import (
	"strconv"
	"strings"
)

type Column struct {
	pos  int
	col  string
	Name string
	Hide bool
}

var (
	InvalidColumn = Column{col: "Invalid", Name: "Invalid"}
)

type Columns []*Column

func (t *table) AddColumn() *Column {

	c := Column{}
	c.pos = len(t.Columns)
	c.col = pos2col(c.pos)
	c.Name = c.col
	t.Columns = append(t.Columns, &c)

	for _, row := range t.Rows {
		cell := Cell{Row: row, Column: &c}
		row.addCell(&cell)
	}
	return &c
}

const ascii = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func pos2col(n int) string {
	prefix := []byte{}
	if n >= len(ascii) {
		prefix = append(prefix, ascii[(n/len(ascii))-1])
		n = n % len(ascii)
	}
	return string(prefix) + string(ascii[n])
}

func col2pos(col string) (pos int) {
	if len(col) < 1 {
		return 0
	}

	col = strings.ToUpper(col)

	if col[0] < 'A' && col[0] > 'Z' {
		return 0
	}
	pos = int(col[0]) - 'A'
	if len(col) == 1 {
		return pos
	}
	pos = (pos + 1) * len(ascii)

	if col[0] < 'A' && col[0] > 'Z' {
		return 0
	}
	pos = pos + int(col[1]) - 'A'
	return pos
}

func (t *table) findColumn(col string) *Column {
	col = strings.ToUpper(col)
	for _, c := range t.Columns {
		if c.col == col {
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
