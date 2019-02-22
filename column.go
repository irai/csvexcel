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

const letters = "ABCDEFGHIJKLMNOPQRSTYUVXZ"

func nextColIndex(n int) string {
	// letters := []byte('ABCDEFGHIJKLMNOPQRSTYUVXZ')
	prefix := []byte{}
	if n >= len(letters) {
		prefix = append(prefix, letters[(n/len(letters))-1])
		n = n % len(letters)
	}
	return string(prefix) + string(letters[n])
}

func str2Pos(index string) (pos int) {
	// letters := []byte('ABCDEFGHIJKLMNOPQRSTYUVXZ')
	if len(index) < 1 {
		return 0
	}

	s := strings.ToUpper(index)

	if s[0] < 'A' && s[0] > 'Z' {
		return 0
	}
	pos = int(s[0]) - 'A'
	if len(index) == 1 {
		return pos
	}
	pos = pos * len(letters)

	if s[0] < 'A' && s[0] > 'Z' {
		return 0
	}
	pos = pos + int(s[1]) - 'A'
	return pos
}

func (t *table) findColumn(index string) *Column {
	index = strings.ToUpper(index)
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
