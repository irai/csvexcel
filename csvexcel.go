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
	OutOfRange   = "Out of range"
	InvalidRange = "Invalid cell range"
)

type table struct {
	Columns Columns
	header  *Row
	Rows    []*Row
}

func New() *table {
	return &table{Columns: []*Column{}, Rows: []*Row{}}
}

// SetHeader set the row number to be the header row for the table. The row value can then be
// used for for cell lookup. A value of 0 means no header row.
//
func (t *table) SetHeader(row int) {
	if row != 0 && row <= len(t.Rows) {
		t.header = t.Rows[row-1]
	} else {
		log.Println("Resetting header - n rows ", len(t.Rows))
		t.header = nil
	}
}

func ParseCSV(in string) (t *table, err error) {
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

		if row.table != t {
			log.Fatal("ERROR tables are different ", &t, row.table)
		}
	}
	return t, nil
}

func ParseFile(filename string) (t *table, err error) {
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

func (t *table) Save(filename string) error {
	outfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outfile.Close()

	writer := csv.NewWriter(outfile)
	buffer := make([]string, len(t.Columns))
	for _, row := range t.Rows {
		if row.Hide == true {
			continue
		}
		for x := range t.Columns {
			if t.Columns[x].Hide != true {
				buffer[x] = row.Cells[x].Value
			}
		}
		if err = writer.Write(buffer); err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}

func (t *table) Cell(name string) *Cell {
	c, r := split2colnumber(name)
	if c == "" || r == -1 {
		return &Cell{Value: InvalidRange}
	}

	r-- // "A1" means row 0
	if r > len(t.Rows) {
		return &Cell{Value: OutOfRange}
	}
	row := t.Rows[r]
	col := t.findColumn(c)
	if col == InvalidColumn || (col != nil && col.pos >= len(row.Cells)) {
		return &Cell{Value: OutOfRange}
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
