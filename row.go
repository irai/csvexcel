package csvexcel

type Row struct {
	table  *table
	Number int
	Cells  []*Cell
	Hide   bool
}

func (r *Row) addCell(cell *Cell) {
	r.Cells = append(r.Cells, cell)
}

func (r *Row) Cell(col string) *Cell {
	c := r.table.findColumn(col)
	// fmt.Println("find cell ", col, c)
	if c == nil || c == InvalidColumn {
		return &Cell{Value: InvalidRange}
	}
	return r.Cells[c.pos]
}

func (t *table) AddRow() *Row {
	row := Row{table: t, Number: len(t.Rows) + 1, Cells: []*Cell{}}
	for _, column := range t.Columns {
		cell := Cell{Row: &row, Column: column}
		row.addCell(&cell)
	}
	t.Rows = append(t.Rows, &row)
	return &row
}

func (t *table) AddRowWithValues(values []string) *Row {
	if len(t.Columns) < len(values) { // expand columns
		n := len(values) - len(t.Columns)
		for i := 0; i < n; i++ {
			t.AddColumn()
		}
	}

	row := t.AddRow()
	for i, value := range values {
		row.Cells[i].Value = value
	}
	return row
}

func (r *Row) Find(value string) (pos int, cell *Cell) {
	for i, c := range r.Cells {
		if c.Value == value {
			return i, c
		}
	}
	return -1, nil
}

func (t *table) FindRow(col string, value string) *Row {
	c := t.findColumn(col)
	if c == nil || c.pos == -1 {
		return nil
	}

	for _, row := range t.Rows {
		if row.Cells[c.pos].Value == value {
			return row
		}
	}
	return nil
}
