// Code generated by ent, DO NOT EDIT.

package adminaction

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the adminaction type in the database.
	Label = "admin_action"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldInfo holds the string denoting the info field in the database.
	FieldInfo = "info"
	// EdgeTeacher holds the string denoting the teacher edge name in mutations.
	EdgeTeacher = "teacher"
	// Table holds the table name of the adminaction in the database.
	Table = "admin_actions"
	// TeacherTable is the table that holds the teacher relation/edge.
	TeacherTable = "admin_actions"
	// TeacherInverseTable is the table name for the Teacher entity.
	// It exists in this package in order to avoid circular dependency with the "teacher" package.
	TeacherInverseTable = "teachers"
	// TeacherColumn is the table column denoting the teacher relation/edge.
	TeacherColumn = "teacher_actions"
)

// Columns holds all SQL columns for adminaction fields.
var Columns = []string{
	FieldID,
	FieldAction,
	FieldInfo,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "admin_actions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"teacher_actions",
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

// OrderOption defines the ordering options for the AdminAction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAction orders the results by the action field.
func ByAction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAction, opts...).ToFunc()
}

// ByInfo orders the results by the info field.
func ByInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInfo, opts...).ToFunc()
}

// ByTeacherField orders the results by teacher field.
func ByTeacherField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeacherStep(), sql.OrderByField(field, opts...))
	}
}
func newTeacherStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeacherInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TeacherTable, TeacherColumn),
	)
}
