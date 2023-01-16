// Code generated by ent, DO NOT EDIT.

package upscalerealtime

const (
	// Label holds the string label denoting the upscalerealtime type in the database.
	Label = "upscale_realtime"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the upscalerealtime in the database.
	Table = "upscale_realtime"
)

// Columns holds all SQL columns for upscalerealtime fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
