package csvexcel

type Row struct {
	Number int
	Cells  []*Cell
	Hide   bool
}

func (r *Row) addCell(cell *Cell) {
	r.Cells = append(r.Cells, cell)
}

func (r *Row) Column(c string) {
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

func (t *table) FindRow(col string, value string) *Row {
	c := t.findColumn(col)
	if c == nil {
		return nil
	}

	for _, row := range t.Rows {
		if row.Cells[c.pos].Value == value {
			return row
		}
	}
	return nil
}
