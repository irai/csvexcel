package csvexcel

type Cell struct {
	Row    *Row
	Column *Column
	Type   string
	Value  string
}
