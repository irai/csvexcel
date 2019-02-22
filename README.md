
# csvexcel
csvexcel is a csv library with excel like table manipulation.

// Load file
table, err := csvexcel.Open("filename")

// or parse
table, err := csvexcel.ParseCSV(in string)

// or create from scratch
table, err := csvexcel.New(nCol int)

table.Print()

table.AddColumn()

table.Columns()



table.Cell("A1").Value

for row := table.Rows() {
  row.Cell("A")
}
