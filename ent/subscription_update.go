// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/subscription"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

// SubscriptionUpdate is the builder for updating Subscription entities.
type SubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *SubscriptionMutation
}

// Where appends a list predicates to the SubscriptionUpdate builder.
func (su *SubscriptionUpdate) Where(ps ...predicate.Subscription) *SubscriptionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetMethod sets the "method" field.
func (su *SubscriptionUpdate) SetMethod(s string) *SubscriptionUpdate {
	su.mutation.SetMethod(s)
	return su
}

// SetQtty sets the "qtty" field.
func (su *SubscriptionUpdate) SetQtty(i int) *SubscriptionUpdate {
	su.mutation.ResetQtty()
	su.mutation.SetQtty(i)
	return su
}

// AddQtty adds i to the "qtty" field.
func (su *SubscriptionUpdate) AddQtty(i int) *SubscriptionUpdate {
	su.mutation.AddQtty(i)
	return su
}

// SetDate sets the "date" field.
func (su *SubscriptionUpdate) SetDate(t time.Time) *SubscriptionUpdate {
	su.mutation.SetDate(t)
	return su
}

// SetTeacherID sets the "teacher" edge to the Teacher entity by ID.
func (su *SubscriptionUpdate) SetTeacherID(id string) *SubscriptionUpdate {
	su.mutation.SetTeacherID(id)
	return su
}

// SetNillableTeacherID sets the "teacher" edge to the Teacher entity by ID if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableTeacherID(id *string) *SubscriptionUpdate {
	if id != nil {
		su = su.SetTeacherID(*id)
	}
	return su
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (su *SubscriptionUpdate) SetTeacher(t *Teacher) *SubscriptionUpdate {
	return su.SetTeacherID(t.ID)
}

// SetYearID sets the "year" edge to the Year entity by ID.
func (su *SubscriptionUpdate) SetYearID(id string) *SubscriptionUpdate {
	su.mutation.SetYearID(id)
	return su
}

// SetNillableYearID sets the "year" edge to the Year entity by ID if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableYearID(id *string) *SubscriptionUpdate {
	if id != nil {
		su = su.SetYearID(*id)
	}
	return su
}

// SetYear sets the "year" edge to the Year entity.
func (su *SubscriptionUpdate) SetYear(y *Year) *SubscriptionUpdate {
	return su.SetYearID(y.ID)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (su *SubscriptionUpdate) Mutation() *SubscriptionMutation {
	return su.mutation
}

// ClearTeacher clears the "teacher" edge to the Teacher entity.
func (su *SubscriptionUpdate) ClearTeacher() *SubscriptionUpdate {
	su.mutation.ClearTeacher()
	return su
}

// ClearYear clears the "year" edge to the Year entity.
func (su *SubscriptionUpdate) ClearYear() *SubscriptionUpdate {
	su.mutation.ClearYear()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubscriptionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubscriptionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(subscription.Table, subscription.Columns, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Method(); ok {
		_spec.SetField(subscription.FieldMethod, field.TypeString, value)
	}
	if value, ok := su.mutation.Qtty(); ok {
		_spec.SetField(subscription.FieldQtty, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedQtty(); ok {
		_spec.AddField(subscription.FieldQtty, field.TypeInt, value)
	}
	if value, ok := su.mutation.Date(); ok {
		_spec.SetField(subscription.FieldDate, field.TypeTime, value)
	}
	if su.mutation.TeacherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.TeacherTable,
			Columns: []string{subscription.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.TeacherTable,
			Columns: []string{subscription.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.YearCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.YearTable,
			Columns: []string{subscription.YearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(year.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.YearIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.YearTable,
			Columns: []string{subscription.YearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(year.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubscriptionUpdateOne is the builder for updating a single Subscription entity.
type SubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubscriptionMutation
}

// SetMethod sets the "method" field.
func (suo *SubscriptionUpdateOne) SetMethod(s string) *SubscriptionUpdateOne {
	suo.mutation.SetMethod(s)
	return suo
}

// SetQtty sets the "qtty" field.
func (suo *SubscriptionUpdateOne) SetQtty(i int) *SubscriptionUpdateOne {
	suo.mutation.ResetQtty()
	suo.mutation.SetQtty(i)
	return suo
}

// AddQtty adds i to the "qtty" field.
func (suo *SubscriptionUpdateOne) AddQtty(i int) *SubscriptionUpdateOne {
	suo.mutation.AddQtty(i)
	return suo
}

// SetDate sets the "date" field.
func (suo *SubscriptionUpdateOne) SetDate(t time.Time) *SubscriptionUpdateOne {
	suo.mutation.SetDate(t)
	return suo
}

// SetTeacherID sets the "teacher" edge to the Teacher entity by ID.
func (suo *SubscriptionUpdateOne) SetTeacherID(id string) *SubscriptionUpdateOne {
	suo.mutation.SetTeacherID(id)
	return suo
}

// SetNillableTeacherID sets the "teacher" edge to the Teacher entity by ID if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableTeacherID(id *string) *SubscriptionUpdateOne {
	if id != nil {
		suo = suo.SetTeacherID(*id)
	}
	return suo
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (suo *SubscriptionUpdateOne) SetTeacher(t *Teacher) *SubscriptionUpdateOne {
	return suo.SetTeacherID(t.ID)
}

// SetYearID sets the "year" edge to the Year entity by ID.
func (suo *SubscriptionUpdateOne) SetYearID(id string) *SubscriptionUpdateOne {
	suo.mutation.SetYearID(id)
	return suo
}

// SetNillableYearID sets the "year" edge to the Year entity by ID if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableYearID(id *string) *SubscriptionUpdateOne {
	if id != nil {
		suo = suo.SetYearID(*id)
	}
	return suo
}

// SetYear sets the "year" edge to the Year entity.
func (suo *SubscriptionUpdateOne) SetYear(y *Year) *SubscriptionUpdateOne {
	return suo.SetYearID(y.ID)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (suo *SubscriptionUpdateOne) Mutation() *SubscriptionMutation {
	return suo.mutation
}

// ClearTeacher clears the "teacher" edge to the Teacher entity.
func (suo *SubscriptionUpdateOne) ClearTeacher() *SubscriptionUpdateOne {
	suo.mutation.ClearTeacher()
	return suo
}

// ClearYear clears the "year" edge to the Year entity.
func (suo *SubscriptionUpdateOne) ClearYear() *SubscriptionUpdateOne {
	suo.mutation.ClearYear()
	return suo
}

// Where appends a list predicates to the SubscriptionUpdate builder.
func (suo *SubscriptionUpdateOne) Where(ps ...predicate.Subscription) *SubscriptionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubscriptionUpdateOne) Select(field string, fields ...string) *SubscriptionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subscription entity.
func (suo *SubscriptionUpdateOne) Save(ctx context.Context) (*Subscription, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubscriptionUpdateOne) SaveX(ctx context.Context) *Subscription {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *Subscription, err error) {
	_spec := sqlgraph.NewUpdateSpec(subscription.Table, subscription.Columns, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Subscription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscription.FieldID)
		for _, f := range fields {
			if !subscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subscription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Method(); ok {
		_spec.SetField(subscription.FieldMethod, field.TypeString, value)
	}
	if value, ok := suo.mutation.Qtty(); ok {
		_spec.SetField(subscription.FieldQtty, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedQtty(); ok {
		_spec.AddField(subscription.FieldQtty, field.TypeInt, value)
	}
	if value, ok := suo.mutation.Date(); ok {
		_spec.SetField(subscription.FieldDate, field.TypeTime, value)
	}
	if suo.mutation.TeacherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.TeacherTable,
			Columns: []string{subscription.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.TeacherTable,
			Columns: []string{subscription.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.YearCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.YearTable,
			Columns: []string{subscription.YearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(year.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.YearIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.YearTable,
			Columns: []string{subscription.YearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(year.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Subscription{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
