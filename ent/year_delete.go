// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/year"
)

// YearDelete is the builder for deleting a Year entity.
type YearDelete struct {
	config
	hooks    []Hook
	mutation *YearMutation
}

// Where appends a list predicates to the YearDelete builder.
func (yd *YearDelete) Where(ps ...predicate.Year) *YearDelete {
	yd.mutation.Where(ps...)
	return yd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (yd *YearDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, yd.sqlExec, yd.mutation, yd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (yd *YearDelete) ExecX(ctx context.Context) int {
	n, err := yd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (yd *YearDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(year.Table, sqlgraph.NewFieldSpec(year.FieldID, field.TypeString))
	if ps := yd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, yd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	yd.mutation.done = true
	return affected, err
}

// YearDeleteOne is the builder for deleting a single Year entity.
type YearDeleteOne struct {
	yd *YearDelete
}

// Where appends a list predicates to the YearDelete builder.
func (ydo *YearDeleteOne) Where(ps ...predicate.Year) *YearDeleteOne {
	ydo.yd.mutation.Where(ps...)
	return ydo
}

// Exec executes the deletion query.
func (ydo *YearDeleteOne) Exec(ctx context.Context) error {
	n, err := ydo.yd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{year.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ydo *YearDeleteOne) ExecX(ctx context.Context) {
	if err := ydo.Exec(ctx); err != nil {
		panic(err)
	}
}
