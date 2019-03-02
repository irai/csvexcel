
# csvexcel
csvexcel is a csv library with excel like table manipulation. It provide simple constructs on top of CSV file that 
simplify csv operations like:
- read and access cells in a CSV file i.e. table.Cell("A1)
- transform, search and update any cell in table using Excel like coordinates
- support for excel like header row and lookup by header
- hide rows or columsn to easily delete from output file
- write the table to a file
- automatically adjust columns in rows so all rows have the same number of cells

## Getting started
Simply import csvexcel and go modules will add the rest to your go.mod file.
```golang
	import "github.com/irai/csvexcel
```

To load a CSV file use:
```golang
// Load file
table, err := csvexcel.ParseFile("filename")
table.Print()
table.Cell("a2").Value = "new value"
err = table.Save("./changed.csv")
```

To parse from a CSV string
```
in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
table, err := csvexcel.ParseCSV(in)
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
Methods for cell lookup always return a valid pointer. If the coordinate is invalid or out of range the 
value will be set to either csvexcel.OutOfRange or csvexcel.InvalidRange. This has proven useful for quick transformations.

```golang
table.New()
if c := table.Cell("D10"); c.Value == csvexcel.OutOfRange || c.Value == csv.InvalidRange {
log.Error("Invalid range")
}
```

## Limitations
The library is in use with files containing tens of thousands of records however the library manipulate the tables 
in memory so you are limited by the memory available. Lots of optimisations could be done to reduce the memory footprint.
