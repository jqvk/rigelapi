// Code generated by ent, DO NOT EDIT.

package subscription

const (
	// Label holds the string label denoting the subscription type in the database.
	Label = "subscription"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMethod holds the string denoting the method field in the database.
	FieldMethod = "method"
	// FieldQtty holds the string denoting the qtty field in the database.
	FieldQtty = "qtty"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// EdgeTeacher holds the string denoting the teacher edge name in mutations.
	EdgeTeacher = "teacher"
	// EdgeYear holds the string denoting the year edge name in mutations.
	EdgeYear = "year"
	// Table holds the table name of the subscription in the database.
	Table = "subscriptions"
	// TeacherTable is the table that holds the teacher relation/edge.
	TeacherTable = "subscriptions"
	// TeacherInverseTable is the table name for the Teacher entity.
	// It exists in this package in order to avoid circular dependency with the "teacher" package.
	TeacherInverseTable = "teachers"
	// TeacherColumn is the table column denoting the teacher relation/edge.
	TeacherColumn = "teacher_subscriptions"
	// YearTable is the table that holds the year relation/edge.
	YearTable = "subscriptions"
	// YearInverseTable is the table name for the Year entity.
	// It exists in this package in order to avoid circular dependency with the "year" package.
	YearInverseTable = "years"
	// YearColumn is the table column denoting the year relation/edge.
	YearColumn = "year_subscriptions"
)

// Columns holds all SQL columns for subscription fields.
var Columns = []string{
	FieldID,
	FieldMethod,
	FieldQtty,
	FieldDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "subscriptions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"teacher_subscriptions",
	"year_subscriptions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
