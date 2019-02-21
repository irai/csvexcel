# csvexcel
csvexcel is a csv library with excel like table manipulation.

table := csvexcel.Open("filename")
table.Columns()



table.Cell("A1").String

for row := table.Rows() {
  row.Column("A")
}
