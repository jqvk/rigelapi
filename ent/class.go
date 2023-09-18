// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/grade"
	"github.com/vmkevv/rigelapi/ent/school"
	"github.com/vmkevv/rigelapi/ent/subject"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

// Class is the model entity for the Class schema.
type Class struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Parallel holds the value of the "parallel" field.
	Parallel string `json:"parallel,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassQuery when eager-loading is set.
	Edges           ClassEdges `json:"edges"`
	grade_classes   *string
	school_classes  *string
	subject_classes *string
	teacher_classes *string
	year_classes    *string
	selectValues    sql.SelectValues
}

// ClassEdges holds the relations/edges for other nodes in the graph.
type ClassEdges struct {
	// Students holds the value of the students edge.
	Students []*Student `json:"students,omitempty"`
	// ClassPeriods holds the value of the classPeriods edge.
	ClassPeriods []*ClassPeriod `json:"classPeriods,omitempty"`
	// School holds the value of the school edge.
	School *School `json:"school,omitempty"`
	// Teacher holds the value of the teacher edge.
	Teacher *Teacher `json:"teacher,omitempty"`
	// Subject holds the value of the subject edge.
	Subject *Subject `json:"subject,omitempty"`
	// Grade holds the value of the grade edge.
	Grade *Grade `json:"grade,omitempty"`
	// Year holds the value of the year edge.
	Year *Year `json:"year,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) StudentsOrErr() ([]*Student, error) {
	if e.loadedTypes[0] {
		return e.Students, nil
	}
	return nil, &NotLoadedError{edge: "students"}
}

// ClassPeriodsOrErr returns the ClassPeriods value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) ClassPeriodsOrErr() ([]*ClassPeriod, error) {
	if e.loadedTypes[1] {
		return e.ClassPeriods, nil
	}
	return nil, &NotLoadedError{edge: "classPeriods"}
}

// SchoolOrErr returns the School value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) SchoolOrErr() (*School, error) {
	if e.loadedTypes[2] {
		if e.School == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: school.Label}
		}
		return e.School, nil
	}
	return nil, &NotLoadedError{edge: "school"}
}

// TeacherOrErr returns the Teacher value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) TeacherOrErr() (*Teacher, error) {
	if e.loadedTypes[3] {
		if e.Teacher == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: teacher.Label}
		}
		return e.Teacher, nil
	}
	return nil, &NotLoadedError{edge: "teacher"}
}

// SubjectOrErr returns the Subject value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) SubjectOrErr() (*Subject, error) {
	if e.loadedTypes[4] {
		if e.Subject == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: subject.Label}
		}
		return e.Subject, nil
	}
	return nil, &NotLoadedError{edge: "subject"}
}

// GradeOrErr returns the Grade value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) GradeOrErr() (*Grade, error) {
	if e.loadedTypes[5] {
		if e.Grade == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: grade.Label}
		}
		return e.Grade, nil
	}
	return nil, &NotLoadedError{edge: "grade"}
}

// YearOrErr returns the Year value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) YearOrErr() (*Year, error) {
	if e.loadedTypes[6] {
		if e.Year == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: year.Label}
		}
		return e.Year, nil
	}
	return nil, &NotLoadedError{edge: "year"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Class) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case class.FieldID, class.FieldParallel:
			values[i] = new(sql.NullString)
		case class.ForeignKeys[0]: // grade_classes
			values[i] = new(sql.NullString)
		case class.ForeignKeys[1]: // school_classes
			values[i] = new(sql.NullString)
		case class.ForeignKeys[2]: // subject_classes
			values[i] = new(sql.NullString)
		case class.ForeignKeys[3]: // teacher_classes
			values[i] = new(sql.NullString)
		case class.ForeignKeys[4]: // year_classes
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Class fields.
func (c *Class) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case class.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case class.FieldParallel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parallel", values[i])
			} else if value.Valid {
				c.Parallel = value.String
			}
		case class.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field grade_classes", values[i])
			} else if value.Valid {
				c.grade_classes = new(string)
				*c.grade_classes = value.String
			}
		case class.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field school_classes", values[i])
			} else if value.Valid {
				c.school_classes = new(string)
				*c.school_classes = value.String
			}
		case class.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject_classes", values[i])
			} else if value.Valid {
				c.subject_classes = new(string)
				*c.subject_classes = value.String
			}
		case class.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field teacher_classes", values[i])
			} else if value.Valid {
				c.teacher_classes = new(string)
				*c.teacher_classes = value.String
			}
		case class.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field year_classes", values[i])
			} else if value.Valid {
				c.year_classes = new(string)
				*c.year_classes = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Class.
// This includes values selected through modifiers, order, etc.
func (c *Class) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryStudents queries the "students" edge of the Class entity.
func (c *Class) QueryStudents() *StudentQuery {
	return NewClassClient(c.config).QueryStudents(c)
}

// QueryClassPeriods queries the "classPeriods" edge of the Class entity.
func (c *Class) QueryClassPeriods() *ClassPeriodQuery {
	return NewClassClient(c.config).QueryClassPeriods(c)
}

// QuerySchool queries the "school" edge of the Class entity.
func (c *Class) QuerySchool() *SchoolQuery {
	return NewClassClient(c.config).QuerySchool(c)
}

// QueryTeacher queries the "teacher" edge of the Class entity.
func (c *Class) QueryTeacher() *TeacherQuery {
	return NewClassClient(c.config).QueryTeacher(c)
}

// QuerySubject queries the "subject" edge of the Class entity.
func (c *Class) QuerySubject() *SubjectQuery {
	return NewClassClient(c.config).QuerySubject(c)
}

// QueryGrade queries the "grade" edge of the Class entity.
func (c *Class) QueryGrade() *GradeQuery {
	return NewClassClient(c.config).QueryGrade(c)
}

// QueryYear queries the "year" edge of the Class entity.
func (c *Class) QueryYear() *YearQuery {
	return NewClassClient(c.config).QueryYear(c)
}

// Update returns a builder for updating this Class.
// Note that you need to call Class.Unwrap() before calling this method if this Class
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Class) Update() *ClassUpdateOne {
	return NewClassClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Class entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Class) Unwrap() *Class {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Class is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Class) String() string {
	var builder strings.Builder
	builder.WriteString("Class(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("parallel=")
	builder.WriteString(c.Parallel)
	builder.WriteByte(')')
	return builder.String()
}

// Classes is a parsable slice of Class.
type Classes []*Class
