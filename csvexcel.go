package csvexcel

import (
	"fmt"
	"log"
)

func nextColIndex(n int) string {
	b := []byte("ABCDEFGHIJKLMNOPQRSTYUVXZ")
	prefix := []byte{}
	if n >= len(b) {
		prefix = append(prefix, b[(n/len(b))-1])
		n = n % len(b)
	}
	return string(prefix) + string(b[n])
}

type table struct {
	Columns []*Column
	Rows    []*Row
}

func New() table {
	return table{Columns: []*Column{}, Rows: []*Row{}}
}

func (t *table) AddColumn() {
	c := Column{Index: nextColIndex(len(t.Columns))}
	c.Name = c.Index
	t.Columns = append(t.Columns, &c)

	for _, row := range t.Rows {
		cell := Cell{Row: row, Column: &c}
		row.addCell(&cell)
	}
}

func (t *table) AddRow() {
	row := Row{Number: len(t.Rows), Cells: []*Cell{}}
	for _, column := range t.Columns {
		cell := Cell{Row: &row, Column: column}
		row.addCell(&cell)
	}
	t.Rows = append(t.Rows, &row)
}

func (t *table) Print() {
	line := ""
	for _, column := range t.Columns {
		line = fmt.Sprintf("%s%10s,", line, column.Index)
	}
	log.Println(line)

	for _, row := range t.Rows {
		line = ""
		for _, cell := range row.Cells {
			line = fmt.Sprintf("%s%10s,", line, cell.Value)
		}
		log.Println(line)
	}
}

type Row struct {
	Number int
	Cells  []*Cell
	Hide   bool
}

func (r *Row) addCell(cell *Cell) {
	r.Cells = append(r.Cells, cell)
}

type Column struct {
	Index string
	Name  string
	Hide  bool
}

type Cell struct {
	Row    *Row
	Column *Column
	Type   string
	Value  string
}

func main() {

}
