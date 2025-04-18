// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/usercreditrecord"
)

// UserCreditRecordDelete is the builder for deleting a UserCreditRecord entity.
type UserCreditRecordDelete struct {
	config
	hooks    []Hook
	mutation *UserCreditRecordMutation
}

// Where appends a list predicates to the UserCreditRecordDelete builder.
func (ucrd *UserCreditRecordDelete) Where(ps ...predicate.UserCreditRecord) *UserCreditRecordDelete {
	ucrd.mutation.Where(ps...)
	return ucrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucrd *UserCreditRecordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ucrd.hooks) == 0 {
		affected, err = ucrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCreditRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucrd.mutation = mutation
			affected, err = ucrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ucrd.hooks) - 1; i >= 0; i-- {
			if ucrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucrd *UserCreditRecordDelete) ExecX(ctx context.Context) int {
	n, err := ucrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucrd *UserCreditRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: usercreditrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: usercreditrecord.FieldID,
			},
		},
	}
	if ps := ucrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ucrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// UserCreditRecordDeleteOne is the builder for deleting a single UserCreditRecord entity.
type UserCreditRecordDeleteOne struct {
	ucrd *UserCreditRecordDelete
}

// Exec executes the deletion query.
func (ucrdo *UserCreditRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := ucrdo.ucrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usercreditrecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucrdo *UserCreditRecordDeleteOne) ExecX(ctx context.Context) {
	ucrdo.ucrd.ExecX(ctx)
}
