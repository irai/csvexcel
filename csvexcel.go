package csvexcel

func nextColIndex(n int) string {
	b := []byte("ABCDEFGHIJKLMNOPQRSTYUVXZ")
	prefix := []byte{}
	if n >= len(b) {
		prefix = append(prefix, b[(n/len(b))-1])
		n = n % len(b)
	}
	return string(prefix) + string(b[n])
}

type Table struct {
	Columns []Column
	Rows    []Row
}

func (t *Table) AddColumn(index string) {
	c := Column{Index: nextColIndex(len(t.Columns))}
	t.Columns = append(t.Columns, c)
}

type Row struct {
	Number int
	Cells  []*Cell
	Hide   bool
}

type Column struct {
	Index string
	Name  string
	Cells []*Cell
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
