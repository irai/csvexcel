
# csvexcel
csvexcel is a csv library with excel like table manipulation. It provide simple constructs on top of CSV file that 
simplify csv operations like:
- read and access cells in a CSV file i.e. table.Cell("A1)
- transform, search and update any cell in table using Excel like coordinates
- support for excel like header row
- hide rows or columsn an 
- write the output to another table
- automatically adjust cells in rows so all rows have the same number of cells

To use simply import and go modules should add the module to your go.mod file.
```golang
	import "github.com/irai/csvexcel
```

To create or load a CSV table use:
```golang
// Load file
table, err := csvexcel.ParseFile("filename")
table.Print()
table.Cell("a2").Value = "new value"
err = table.Save("./changed.csv")
```

To parse from a string CSV
```
in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
table, err := csvexcel.ParseCSV(in string)
table.Row(2).Hide = true // exclude row 2 in Save() or Print()
err = table.Save("./changed.csv")
```

To create a CSV table from scratch
```
table, err := csvexcel.New()
// You can pass row cells as slice
table.AddRowWithFields([]string{"first name", "last name", "username" })
table.SetHeader(1)  // Row 1 is first as in Excel

// or add an empty row
r := table.AddRow()
r.Cell("first name").Value = "Rob" // column by header
r.Cell("B").Value = "Pike" // column by excel label
table.Cell("C1").Value = "rob" // access the cell through the table
table.Print()
```

To manipulate each row
```golang
for _, row := table.Rows() {
  fmt.Println(row.Cell("A").Value, row.Cell("B").Value, row.Cell("username").Value)
}
```

To add a new column at the end
```golang
in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
table, err := csvexcel.ParseCSV(in string)
table.AddColumn() // add a forth column to all rows
table.Cell("d2").Value = "new column value"
```

## Error handling
The methods for cell lookup always return a valid pointer but to an invalid cell if the
cell coordinate are invalid or out of range. This has proven useful for quick transformations.

```golang
table.New()
if c := table.Cell("D10"); c.Value == csvexcel.OutOfRange || c.Value == csv.InvalidRange {
log.Error("Invalid range")
}
```



