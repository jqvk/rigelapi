// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/grade"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// GradeDelete is the builder for deleting a Grade entity.
type GradeDelete struct {
	config
	hooks    []Hook
	mutation *GradeMutation
}

// Where appends a list predicates to the GradeDelete builder.
func (gd *GradeDelete) Where(ps ...predicate.Grade) *GradeDelete {
	gd.mutation.Where(ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GradeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gd.sqlExec, gd.mutation, gd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GradeDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GradeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(grade.Table, sqlgraph.NewFieldSpec(grade.FieldID, field.TypeString))
	if ps := gd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gd.mutation.done = true
	return affected, err
}

// GradeDeleteOne is the builder for deleting a single Grade entity.
type GradeDeleteOne struct {
	gd *GradeDelete
}

// Where appends a list predicates to the GradeDelete builder.
func (gdo *GradeDeleteOne) Where(ps ...predicate.Grade) *GradeDeleteOne {
	gdo.gd.mutation.Where(ps...)
	return gdo
}

// Exec executes the deletion query.
func (gdo *GradeDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{grade.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GradeDeleteOne) ExecX(ctx context.Context) {
	if err := gdo.Exec(ctx); err != nil {
		panic(err)
	}
}
