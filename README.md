# csvToSlice
Converts a cvs file to a slice based on the column specified.

Example

var s []string
s = csvToSlice([column starting at 0], [Whether it has a header you want to cut], [file location])<br />
s = csvToSlice(0, true, "test.csv")
