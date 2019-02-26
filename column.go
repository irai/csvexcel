package csvexcel

import (
	"strconv"
	"strings"
)

type Column struct {
	pos  int
	col  string
	Hide bool
}

var (
	InvalidColumn = &Column{pos: -1, col: "Invalid"}
)

type Columns []*Column

func (t *table) AddColumn() *Column {

	c := Column{}
	c.pos = len(t.Columns)
	c.col = pos2col(c.pos)
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
	col = strings.ToUpper(col)
	if len(col) < 1 || len(col) > 2 { // only accept two letters for now
		return -1
	}

	col = strings.ToUpper(col)

	if col[0] < 'A' && col[0] > 'Z' {
		return -1
	}
	pos = int(col[0]) - 'A'
	if len(col) == 1 {
		return pos
	}
	pos = (pos + 1) * len(ascii)

	if col[1] < 'A' && col[1] > 'Z' {
		return -1
	}
	pos = pos + int(col[1]) - 'A'
	return pos
}

func (t *table) findColumn(col string) *Column {
	pos := col2pos(col)
	if pos == -1 {
		if t.header != nil {
			pos, _ := t.header.Find(col)
			if pos == -1 {
				return InvalidColumn
			}
			return t.Columns[pos]
		}
		return InvalidColumn
	}

	if pos >= len(t.Columns) {
		return InvalidColumn
	}
	return t.Columns[pos]
}

func split2colnumber(index string) (col string, row int) {

	if len(index) < 2 {
		return "", -1
	}

	s := strings.ToUpper(index)

	if s[0] <= 'A' && s[0] >= 'Z' {
		return "", -1
	}
	col = s[0:1]
	r := s[1:]

	if s[1] >= 'A' && s[1] <= 'Z' {
		col = s[0:2]
		if len(s) < 3 {
			return "", -1
		}
		r = s[2:]

	}

	var err error
	if row, err = strconv.Atoi(r); err != nil || row == 0 {
		// log.Println("strconv error ", err)
		return "", -1
	}
	return col, row
}
