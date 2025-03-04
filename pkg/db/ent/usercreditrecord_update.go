// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/usercreditrecord"
	"github.com/google/uuid"
)

// UserCreditRecordUpdate is the builder for updating UserCreditRecord entities.
type UserCreditRecordUpdate struct {
	config
	hooks     []Hook
	mutation  *UserCreditRecordMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserCreditRecordUpdate builder.
func (ucru *UserCreditRecordUpdate) Where(ps ...predicate.UserCreditRecord) *UserCreditRecordUpdate {
	ucru.mutation.Where(ps...)
	return ucru
}

// SetCreatedAt sets the "created_at" field.
func (ucru *UserCreditRecordUpdate) SetCreatedAt(u uint32) *UserCreditRecordUpdate {
	ucru.mutation.ResetCreatedAt()
	ucru.mutation.SetCreatedAt(u)
	return ucru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ucru *UserCreditRecordUpdate) SetNillableCreatedAt(u *uint32) *UserCreditRecordUpdate {
	if u != nil {
		ucru.SetCreatedAt(*u)
	}
	return ucru
}

// AddCreatedAt adds u to the "created_at" field.
func (ucru *UserCreditRecordUpdate) AddCreatedAt(u int32) *UserCreditRecordUpdate {
	ucru.mutation.AddCreatedAt(u)
	return ucru
}

// SetUpdatedAt sets the "updated_at" field.
func (ucru *UserCreditRecordUpdate) SetUpdatedAt(u uint32) *UserCreditRecordUpdate {
	ucru.mutation.ResetUpdatedAt()
	ucru.mutation.SetUpdatedAt(u)
	return ucru
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ucru *UserCreditRecordUpdate) AddUpdatedAt(u int32) *UserCreditRecordUpdate {
	ucru.mutation.AddUpdatedAt(u)
	return ucru
}

// SetDeletedAt sets the "deleted_at" field.
func (ucru *UserCreditRecordUpdate) SetDeletedAt(u uint32) *UserCreditRecordUpdate {
	ucru.mutation.ResetDeletedAt()
	ucru.mutation.SetDeletedAt(u)
	return ucru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ucru *UserCreditRecordUpdate) SetNillableDeletedAt(u *uint32) *UserCreditRecordUpdate {
	if u != nil {
		ucru.SetDeletedAt(*u)
	}
	return ucru
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ucru *UserCreditRecordUpdate) AddDeletedAt(u int32) *UserCreditRecordUpdate {
	ucru.mutation.AddDeletedAt(u)
	return ucru
}

// SetEntID sets the "ent_id" field.
func (ucru *UserCreditRecordUpdate) SetEntID(u uuid.UUID) *UserCreditRecordUpdate {
	ucru.mutation.SetEntID(u)
	return ucru
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ucru *UserCreditRecordUpdate) SetNillableEntID(u *uuid.UUID) *UserCreditRecordUpdate {
	if u != nil {
		ucru.SetEntID(*u)
	}
	return ucru
}

// SetAppID sets the "app_id" field.
func (ucru *UserCreditRecordUpdate) SetAppID(u uuid.UUID) *UserCreditRecordUpdate {
	ucru.mutation.SetAppID(u)
	return ucru
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ucru *UserCreditRecordUpdate) SetNillableAppID(u *uuid.UUID) *UserCreditRecordUpdate {
	if u != nil {
		ucru.SetAppID(*u)
	}
	return ucru
}

// ClearAppID clears the value of the "app_id" field.
func (ucru *UserCreditRecordUpdate) ClearAppID() *UserCreditRecordUpdate {
	ucru.mutation.ClearAppID()
	return ucru
}

// SetUserID sets the "user_id" field.
func (ucru *UserCreditRecordUpdate) SetUserID(u uuid.UUID) *UserCreditRecordUpdate {
	ucru.mutation.SetUserID(u)
	return ucru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucru *UserCreditRecordUpdate) SetNillableUserID(u *uuid.UUID) *UserCreditRecordUpdate {
	if u != nil {
		ucru.SetUserID(*u)
	}
	return ucru
}

// ClearUserID clears the value of the "user_id" field.
func (ucru *UserCreditRecordUpdate) ClearUserID() *UserCreditRecordUpdate {
	ucru.mutation.ClearUserID()
	return ucru
}

// SetOperationType sets the "operation_type" field.
func (ucru *UserCreditRecordUpdate) SetOperationType(s string) *UserCreditRecordUpdate {
	ucru.mutation.SetOperationType(s)
	return ucru
}

// SetNillableOperationType sets the "operation_type" field if the given value is not nil.
func (ucru *UserCreditRecordUpdate) SetNillableOperationType(s *string) *UserCreditRecordUpdate {
	if s != nil {
		ucru.SetOperationType(*s)
	}
	return ucru
}

// ClearOperationType clears the value of the "operation_type" field.
func (ucru *UserCreditRecordUpdate) ClearOperationType() *UserCreditRecordUpdate {
	ucru.mutation.ClearOperationType()
	return ucru
}

// SetCreditsChange sets the "credits_change" field.
func (ucru *UserCreditRecordUpdate) SetCreditsChange(u uint32) *UserCreditRecordUpdate {
	ucru.mutation.ResetCreditsChange()
	ucru.mutation.SetCreditsChange(u)
	return ucru
}

// SetNillableCreditsChange sets the "credits_change" field if the given value is not nil.
func (ucru *UserCreditRecordUpdate) SetNillableCreditsChange(u *uint32) *UserCreditRecordUpdate {
	if u != nil {
		ucru.SetCreditsChange(*u)
	}
	return ucru
}

// AddCreditsChange adds u to the "credits_change" field.
func (ucru *UserCreditRecordUpdate) AddCreditsChange(u int32) *UserCreditRecordUpdate {
	ucru.mutation.AddCreditsChange(u)
	return ucru
}

// ClearCreditsChange clears the value of the "credits_change" field.
func (ucru *UserCreditRecordUpdate) ClearCreditsChange() *UserCreditRecordUpdate {
	ucru.mutation.ClearCreditsChange()
	return ucru
}

// SetExtra sets the "extra" field.
func (ucru *UserCreditRecordUpdate) SetExtra(s string) *UserCreditRecordUpdate {
	ucru.mutation.SetExtra(s)
	return ucru
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (ucru *UserCreditRecordUpdate) SetNillableExtra(s *string) *UserCreditRecordUpdate {
	if s != nil {
		ucru.SetExtra(*s)
	}
	return ucru
}

// ClearExtra clears the value of the "extra" field.
func (ucru *UserCreditRecordUpdate) ClearExtra() *UserCreditRecordUpdate {
	ucru.mutation.ClearExtra()
	return ucru
}

// Mutation returns the UserCreditRecordMutation object of the builder.
func (ucru *UserCreditRecordUpdate) Mutation() *UserCreditRecordMutation {
	return ucru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucru *UserCreditRecordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ucru.defaults(); err != nil {
		return 0, err
	}
	if len(ucru.hooks) == 0 {
		affected, err = ucru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCreditRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucru.mutation = mutation
			affected, err = ucru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ucru.hooks) - 1; i >= 0; i-- {
			if ucru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucru *UserCreditRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := ucru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucru *UserCreditRecordUpdate) Exec(ctx context.Context) error {
	_, err := ucru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucru *UserCreditRecordUpdate) ExecX(ctx context.Context) {
	if err := ucru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucru *UserCreditRecordUpdate) defaults() error {
	if _, ok := ucru.mutation.UpdatedAt(); !ok {
		if usercreditrecord.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized usercreditrecord.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := usercreditrecord.UpdateDefaultUpdatedAt()
		ucru.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ucru *UserCreditRecordUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserCreditRecordUpdate {
	ucru.modifiers = append(ucru.modifiers, modifiers...)
	return ucru
}

func (ucru *UserCreditRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercreditrecord.Table,
			Columns: usercreditrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: usercreditrecord.FieldID,
			},
		},
	}
	if ps := ucru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldCreatedAt,
		})
	}
	if value, ok := ucru.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldCreatedAt,
		})
	}
	if value, ok := ucru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldUpdatedAt,
		})
	}
	if value, ok := ucru.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldUpdatedAt,
		})
	}
	if value, ok := ucru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldDeletedAt,
		})
	}
	if value, ok := ucru.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldDeletedAt,
		})
	}
	if value, ok := ucru.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercreditrecord.FieldEntID,
		})
	}
	if value, ok := ucru.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercreditrecord.FieldAppID,
		})
	}
	if ucru.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: usercreditrecord.FieldAppID,
		})
	}
	if value, ok := ucru.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercreditrecord.FieldUserID,
		})
	}
	if ucru.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: usercreditrecord.FieldUserID,
		})
	}
	if value, ok := ucru.mutation.OperationType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercreditrecord.FieldOperationType,
		})
	}
	if ucru.mutation.OperationTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: usercreditrecord.FieldOperationType,
		})
	}
	if value, ok := ucru.mutation.CreditsChange(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldCreditsChange,
		})
	}
	if value, ok := ucru.mutation.AddedCreditsChange(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldCreditsChange,
		})
	}
	if ucru.mutation.CreditsChangeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: usercreditrecord.FieldCreditsChange,
		})
	}
	if value, ok := ucru.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercreditrecord.FieldExtra,
		})
	}
	if ucru.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: usercreditrecord.FieldExtra,
		})
	}
	_spec.Modifiers = ucru.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ucru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercreditrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserCreditRecordUpdateOne is the builder for updating a single UserCreditRecord entity.
type UserCreditRecordUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserCreditRecordMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ucruo *UserCreditRecordUpdateOne) SetCreatedAt(u uint32) *UserCreditRecordUpdateOne {
	ucruo.mutation.ResetCreatedAt()
	ucruo.mutation.SetCreatedAt(u)
	return ucruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ucruo *UserCreditRecordUpdateOne) SetNillableCreatedAt(u *uint32) *UserCreditRecordUpdateOne {
	if u != nil {
		ucruo.SetCreatedAt(*u)
	}
	return ucruo
}

// AddCreatedAt adds u to the "created_at" field.
func (ucruo *UserCreditRecordUpdateOne) AddCreatedAt(u int32) *UserCreditRecordUpdateOne {
	ucruo.mutation.AddCreatedAt(u)
	return ucruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ucruo *UserCreditRecordUpdateOne) SetUpdatedAt(u uint32) *UserCreditRecordUpdateOne {
	ucruo.mutation.ResetUpdatedAt()
	ucruo.mutation.SetUpdatedAt(u)
	return ucruo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ucruo *UserCreditRecordUpdateOne) AddUpdatedAt(u int32) *UserCreditRecordUpdateOne {
	ucruo.mutation.AddUpdatedAt(u)
	return ucruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ucruo *UserCreditRecordUpdateOne) SetDeletedAt(u uint32) *UserCreditRecordUpdateOne {
	ucruo.mutation.ResetDeletedAt()
	ucruo.mutation.SetDeletedAt(u)
	return ucruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ucruo *UserCreditRecordUpdateOne) SetNillableDeletedAt(u *uint32) *UserCreditRecordUpdateOne {
	if u != nil {
		ucruo.SetDeletedAt(*u)
	}
	return ucruo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ucruo *UserCreditRecordUpdateOne) AddDeletedAt(u int32) *UserCreditRecordUpdateOne {
	ucruo.mutation.AddDeletedAt(u)
	return ucruo
}

// SetEntID sets the "ent_id" field.
func (ucruo *UserCreditRecordUpdateOne) SetEntID(u uuid.UUID) *UserCreditRecordUpdateOne {
	ucruo.mutation.SetEntID(u)
	return ucruo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ucruo *UserCreditRecordUpdateOne) SetNillableEntID(u *uuid.UUID) *UserCreditRecordUpdateOne {
	if u != nil {
		ucruo.SetEntID(*u)
	}
	return ucruo
}

// SetAppID sets the "app_id" field.
func (ucruo *UserCreditRecordUpdateOne) SetAppID(u uuid.UUID) *UserCreditRecordUpdateOne {
	ucruo.mutation.SetAppID(u)
	return ucruo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ucruo *UserCreditRecordUpdateOne) SetNillableAppID(u *uuid.UUID) *UserCreditRecordUpdateOne {
	if u != nil {
		ucruo.SetAppID(*u)
	}
	return ucruo
}

// ClearAppID clears the value of the "app_id" field.
func (ucruo *UserCreditRecordUpdateOne) ClearAppID() *UserCreditRecordUpdateOne {
	ucruo.mutation.ClearAppID()
	return ucruo
}

// SetUserID sets the "user_id" field.
func (ucruo *UserCreditRecordUpdateOne) SetUserID(u uuid.UUID) *UserCreditRecordUpdateOne {
	ucruo.mutation.SetUserID(u)
	return ucruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucruo *UserCreditRecordUpdateOne) SetNillableUserID(u *uuid.UUID) *UserCreditRecordUpdateOne {
	if u != nil {
		ucruo.SetUserID(*u)
	}
	return ucruo
}

// ClearUserID clears the value of the "user_id" field.
func (ucruo *UserCreditRecordUpdateOne) ClearUserID() *UserCreditRecordUpdateOne {
	ucruo.mutation.ClearUserID()
	return ucruo
}

// SetOperationType sets the "operation_type" field.
func (ucruo *UserCreditRecordUpdateOne) SetOperationType(s string) *UserCreditRecordUpdateOne {
	ucruo.mutation.SetOperationType(s)
	return ucruo
}

// SetNillableOperationType sets the "operation_type" field if the given value is not nil.
func (ucruo *UserCreditRecordUpdateOne) SetNillableOperationType(s *string) *UserCreditRecordUpdateOne {
	if s != nil {
		ucruo.SetOperationType(*s)
	}
	return ucruo
}

// ClearOperationType clears the value of the "operation_type" field.
func (ucruo *UserCreditRecordUpdateOne) ClearOperationType() *UserCreditRecordUpdateOne {
	ucruo.mutation.ClearOperationType()
	return ucruo
}

// SetCreditsChange sets the "credits_change" field.
func (ucruo *UserCreditRecordUpdateOne) SetCreditsChange(u uint32) *UserCreditRecordUpdateOne {
	ucruo.mutation.ResetCreditsChange()
	ucruo.mutation.SetCreditsChange(u)
	return ucruo
}

// SetNillableCreditsChange sets the "credits_change" field if the given value is not nil.
func (ucruo *UserCreditRecordUpdateOne) SetNillableCreditsChange(u *uint32) *UserCreditRecordUpdateOne {
	if u != nil {
		ucruo.SetCreditsChange(*u)
	}
	return ucruo
}

// AddCreditsChange adds u to the "credits_change" field.
func (ucruo *UserCreditRecordUpdateOne) AddCreditsChange(u int32) *UserCreditRecordUpdateOne {
	ucruo.mutation.AddCreditsChange(u)
	return ucruo
}

// ClearCreditsChange clears the value of the "credits_change" field.
func (ucruo *UserCreditRecordUpdateOne) ClearCreditsChange() *UserCreditRecordUpdateOne {
	ucruo.mutation.ClearCreditsChange()
	return ucruo
}

// SetExtra sets the "extra" field.
func (ucruo *UserCreditRecordUpdateOne) SetExtra(s string) *UserCreditRecordUpdateOne {
	ucruo.mutation.SetExtra(s)
	return ucruo
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (ucruo *UserCreditRecordUpdateOne) SetNillableExtra(s *string) *UserCreditRecordUpdateOne {
	if s != nil {
		ucruo.SetExtra(*s)
	}
	return ucruo
}

// ClearExtra clears the value of the "extra" field.
func (ucruo *UserCreditRecordUpdateOne) ClearExtra() *UserCreditRecordUpdateOne {
	ucruo.mutation.ClearExtra()
	return ucruo
}

// Mutation returns the UserCreditRecordMutation object of the builder.
func (ucruo *UserCreditRecordUpdateOne) Mutation() *UserCreditRecordMutation {
	return ucruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucruo *UserCreditRecordUpdateOne) Select(field string, fields ...string) *UserCreditRecordUpdateOne {
	ucruo.fields = append([]string{field}, fields...)
	return ucruo
}

// Save executes the query and returns the updated UserCreditRecord entity.
func (ucruo *UserCreditRecordUpdateOne) Save(ctx context.Context) (*UserCreditRecord, error) {
	var (
		err  error
		node *UserCreditRecord
	)
	if err := ucruo.defaults(); err != nil {
		return nil, err
	}
	if len(ucruo.hooks) == 0 {
		node, err = ucruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCreditRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucruo.mutation = mutation
			node, err = ucruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ucruo.hooks) - 1; i >= 0; i-- {
			if ucruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ucruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserCreditRecord)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserCreditRecordMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucruo *UserCreditRecordUpdateOne) SaveX(ctx context.Context) *UserCreditRecord {
	node, err := ucruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucruo *UserCreditRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := ucruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucruo *UserCreditRecordUpdateOne) ExecX(ctx context.Context) {
	if err := ucruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucruo *UserCreditRecordUpdateOne) defaults() error {
	if _, ok := ucruo.mutation.UpdatedAt(); !ok {
		if usercreditrecord.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized usercreditrecord.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := usercreditrecord.UpdateDefaultUpdatedAt()
		ucruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ucruo *UserCreditRecordUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserCreditRecordUpdateOne {
	ucruo.modifiers = append(ucruo.modifiers, modifiers...)
	return ucruo
}

func (ucruo *UserCreditRecordUpdateOne) sqlSave(ctx context.Context) (_node *UserCreditRecord, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercreditrecord.Table,
			Columns: usercreditrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: usercreditrecord.FieldID,
			},
		},
	}
	id, ok := ucruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserCreditRecord.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercreditrecord.FieldID)
		for _, f := range fields {
			if !usercreditrecord.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usercreditrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldCreatedAt,
		})
	}
	if value, ok := ucruo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldCreatedAt,
		})
	}
	if value, ok := ucruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldUpdatedAt,
		})
	}
	if value, ok := ucruo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldUpdatedAt,
		})
	}
	if value, ok := ucruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldDeletedAt,
		})
	}
	if value, ok := ucruo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldDeletedAt,
		})
	}
	if value, ok := ucruo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercreditrecord.FieldEntID,
		})
	}
	if value, ok := ucruo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercreditrecord.FieldAppID,
		})
	}
	if ucruo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: usercreditrecord.FieldAppID,
		})
	}
	if value, ok := ucruo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usercreditrecord.FieldUserID,
		})
	}
	if ucruo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: usercreditrecord.FieldUserID,
		})
	}
	if value, ok := ucruo.mutation.OperationType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercreditrecord.FieldOperationType,
		})
	}
	if ucruo.mutation.OperationTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: usercreditrecord.FieldOperationType,
		})
	}
	if value, ok := ucruo.mutation.CreditsChange(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldCreditsChange,
		})
	}
	if value, ok := ucruo.mutation.AddedCreditsChange(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usercreditrecord.FieldCreditsChange,
		})
	}
	if ucruo.mutation.CreditsChangeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: usercreditrecord.FieldCreditsChange,
		})
	}
	if value, ok := ucruo.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercreditrecord.FieldExtra,
		})
	}
	if ucruo.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: usercreditrecord.FieldExtra,
		})
	}
	_spec.Modifiers = ucruo.modifiers
	_node = &UserCreditRecord{config: ucruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercreditrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
