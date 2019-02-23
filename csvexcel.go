package csvexcel

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	OutOfRange   = &Cell{Value: "Out of range"}
	InvalidRange = &Cell{Value: "Invalid cell range"}
)

type table struct {
	Columns Columns
	Rows    []*Row
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

func Open(filename string) (t table, err error) {
	csvfile, err := os.Open(filename)
	if err != nil {
		return t, err
	}
	defer csvfile.Close()

	t = New()
	r := csv.NewReader(csvfile)
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
	col := t.findColumn(c)
	if col == nil || (col != nil && col.pos >= len(row.Cells)) {
		return OutOfRange
	}
	return row.Cells[col.pos]
}

func (t *table) Print() {
	line := "row/name  "
	for _, column := range t.Columns {
		if column.Hide == true {
			continue
		}
		line = fmt.Sprintf("%s%10s,", line, column.col)
	}
	log.Println(line)

	for i, row := range t.Rows {
		if row.Hide == true {
			continue
		}
		line = fmt.Sprintf("  %3d", i+1)
		for i, cell := range row.Cells {
			if t.Columns[i].Hide == true {
				continue
			}
			line = fmt.Sprintf("%s%10s,", line, cell.Value)
		}
		log.Println(line)
	}
}
