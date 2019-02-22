package csvexcel

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

var (
	OutOfRange   = &Cell{Value: "Out of range"}
	InvalidRange = &Cell{Value: "Invalid cell range"}
)

type table struct {
	Columns []*Column
	Rows    []*Row
}

func (t *table) toColumn(index string) *Column {
	for _, c := range t.Columns {
		if c.Index == index {
			return c
		}
	}
	return nil
}

func New() table {
	return table{Columns: []*Column{}, Rows: []*Row{}}
}

func ParseCSV(in string) (t table, err error) {
	r := csv.NewReader(strings.NewReader(in))

	t = New()

	lc := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if lc == 0 {
			for range record {
				t.AddColumn()
			}
		}
		lc++
		row := t.AddRow()
		for i, cell := range row.Cells {
			cell.Value = record[i]
		}
	}
	return t, nil
}

func (t *table) Cell(name string) *Cell {
	c, r := toIndex(name)
	if c == "" || r == 0 {
		return InvalidRange
	}

	r-- // "A1" means row 0
	if r >= len(t.Rows) {
		return OutOfRange
	}
	row := t.Rows[r]
	col := t.toColumn(c)
	if col == nil || (col != nil && col.pos >= len(row.Cells)) {
		return OutOfRange
	}
	return row.Cells[col.pos]
}

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

func (t *table) AddRow() *Row {
	row := Row{Number: len(t.Rows), Cells: []*Cell{}}
	for _, column := range t.Columns {
		cell := Cell{Row: &row, Column: column}
		row.addCell(&cell)
	}
	t.Rows = append(t.Rows, &row)
	return &row
}

func (t *table) Print() {
	line := "row/name  "
	for _, column := range t.Columns {
		line = fmt.Sprintf("%s%10s,", line, column.Index)
	}
	log.Println(line)

	for i, row := range t.Rows {
		line = fmt.Sprintf("  %3d", i+1)
		for _, cell := range row.Cells {
			line = fmt.Sprintf("%s%10s,", line, cell.Value)
		}
		log.Println(line)
	}
}
