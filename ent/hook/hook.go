// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/vmkevv/rigelapi/ent"
)

// The ActivityFunc type is an adapter to allow the use of ordinary
// function as Activity mutator.
type ActivityFunc func(context.Context, *ent.ActivityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ActivityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ActivityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ActivityMutation", m)
	}
	return f(ctx, mv)
}

// The AppErrorFunc type is an adapter to allow the use of ordinary
// function as AppError mutator.
type AppErrorFunc func(context.Context, *ent.AppErrorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppErrorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppErrorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppErrorMutation", m)
	}
	return f(ctx, mv)
}

// The AreaFunc type is an adapter to allow the use of ordinary
// function as Area mutator.
type AreaFunc func(context.Context, *ent.AreaMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AreaFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AreaMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AreaMutation", m)
	}
	return f(ctx, mv)
}

// The AttendanceFunc type is an adapter to allow the use of ordinary
// function as Attendance mutator.
type AttendanceFunc func(context.Context, *ent.AttendanceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AttendanceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AttendanceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AttendanceMutation", m)
	}
	return f(ctx, mv)
}

// The AttendanceDayFunc type is an adapter to allow the use of ordinary
// function as AttendanceDay mutator.
type AttendanceDayFunc func(context.Context, *ent.AttendanceDayMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AttendanceDayFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AttendanceDayMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AttendanceDayMutation", m)
	}
	return f(ctx, mv)
}

// The ClassFunc type is an adapter to allow the use of ordinary
// function as Class mutator.
type ClassFunc func(context.Context, *ent.ClassMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ClassFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ClassMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ClassMutation", m)
	}
	return f(ctx, mv)
}

// The ClassPeriodFunc type is an adapter to allow the use of ordinary
// function as ClassPeriod mutator.
type ClassPeriodFunc func(context.Context, *ent.ClassPeriodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ClassPeriodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ClassPeriodMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ClassPeriodMutation", m)
	}
	return f(ctx, mv)
}

// The DptoFunc type is an adapter to allow the use of ordinary
// function as Dpto mutator.
type DptoFunc func(context.Context, *ent.DptoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DptoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DptoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DptoMutation", m)
	}
	return f(ctx, mv)
}

// The GradeFunc type is an adapter to allow the use of ordinary
// function as Grade mutator.
type GradeFunc func(context.Context, *ent.GradeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GradeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GradeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GradeMutation", m)
	}
	return f(ctx, mv)
}

// The MunicipioFunc type is an adapter to allow the use of ordinary
// function as Municipio mutator.
type MunicipioFunc func(context.Context, *ent.MunicipioMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MunicipioFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MunicipioMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MunicipioMutation", m)
	}
	return f(ctx, mv)
}

// The PeriodFunc type is an adapter to allow the use of ordinary
// function as Period mutator.
type PeriodFunc func(context.Context, *ent.PeriodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PeriodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PeriodMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PeriodMutation", m)
	}
	return f(ctx, mv)
}

// The ProvinciaFunc type is an adapter to allow the use of ordinary
// function as Provincia mutator.
type ProvinciaFunc func(context.Context, *ent.ProvinciaMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProvinciaFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProvinciaMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProvinciaMutation", m)
	}
	return f(ctx, mv)
}

// The SchoolFunc type is an adapter to allow the use of ordinary
// function as School mutator.
type SchoolFunc func(context.Context, *ent.SchoolMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SchoolFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SchoolMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SchoolMutation", m)
	}
	return f(ctx, mv)
}

// The ScoreFunc type is an adapter to allow the use of ordinary
// function as Score mutator.
type ScoreFunc func(context.Context, *ent.ScoreMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScoreFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ScoreMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScoreMutation", m)
	}
	return f(ctx, mv)
}

// The StudentFunc type is an adapter to allow the use of ordinary
// function as Student mutator.
type StudentFunc func(context.Context, *ent.StudentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StudentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StudentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StudentMutation", m)
	}
	return f(ctx, mv)
}

// The SubjectFunc type is an adapter to allow the use of ordinary
// function as Subject mutator.
type SubjectFunc func(context.Context, *ent.SubjectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SubjectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SubjectMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SubjectMutation", m)
	}
	return f(ctx, mv)
}

// The TeacherFunc type is an adapter to allow the use of ordinary
// function as Teacher mutator.
type TeacherFunc func(context.Context, *ent.TeacherMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TeacherFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TeacherMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TeacherMutation", m)
	}
	return f(ctx, mv)
}

// The YearFunc type is an adapter to allow the use of ordinary
// function as Year mutator.
type YearFunc func(context.Context, *ent.YearMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f YearFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.YearMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.YearMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
