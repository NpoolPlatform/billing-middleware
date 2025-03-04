// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/detail"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// DetailUpdate is the builder for updating Detail entities.
type DetailUpdate struct {
	config
	hooks     []Hook
	mutation  *DetailMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DetailUpdate builder.
func (du *DetailUpdate) Where(ps ...predicate.Detail) *DetailUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DetailUpdate) SetCreatedAt(u uint32) *DetailUpdate {
	du.mutation.ResetCreatedAt()
	du.mutation.SetCreatedAt(u)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DetailUpdate) SetNillableCreatedAt(u *uint32) *DetailUpdate {
	if u != nil {
		du.SetCreatedAt(*u)
	}
	return du
}

// AddCreatedAt adds u to the "created_at" field.
func (du *DetailUpdate) AddCreatedAt(u int32) *DetailUpdate {
	du.mutation.AddCreatedAt(u)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DetailUpdate) SetUpdatedAt(u uint32) *DetailUpdate {
	du.mutation.ResetUpdatedAt()
	du.mutation.SetUpdatedAt(u)
	return du
}

// AddUpdatedAt adds u to the "updated_at" field.
func (du *DetailUpdate) AddUpdatedAt(u int32) *DetailUpdate {
	du.mutation.AddUpdatedAt(u)
	return du
}

// SetDeletedAt sets the "deleted_at" field.
func (du *DetailUpdate) SetDeletedAt(u uint32) *DetailUpdate {
	du.mutation.ResetDeletedAt()
	du.mutation.SetDeletedAt(u)
	return du
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (du *DetailUpdate) SetNillableDeletedAt(u *uint32) *DetailUpdate {
	if u != nil {
		du.SetDeletedAt(*u)
	}
	return du
}

// AddDeletedAt adds u to the "deleted_at" field.
func (du *DetailUpdate) AddDeletedAt(u int32) *DetailUpdate {
	du.mutation.AddDeletedAt(u)
	return du
}

// SetEntID sets the "ent_id" field.
func (du *DetailUpdate) SetEntID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetEntID(u)
	return du
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableEntID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetEntID(*u)
	}
	return du
}

// SetSampleCol sets the "sample_col" field.
func (du *DetailUpdate) SetSampleCol(s string) *DetailUpdate {
	du.mutation.SetSampleCol(s)
	return du
}

// SetNillableSampleCol sets the "sample_col" field if the given value is not nil.
func (du *DetailUpdate) SetNillableSampleCol(s *string) *DetailUpdate {
	if s != nil {
		du.SetSampleCol(*s)
	}
	return du
}

// ClearSampleCol clears the value of the "sample_col" field.
func (du *DetailUpdate) ClearSampleCol() *DetailUpdate {
	du.mutation.ClearSampleCol()
	return du
}

// Mutation returns the DetailMutation object of the builder.
func (du *DetailUpdate) Mutation() *DetailMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := du.defaults(); err != nil {
		return 0, err
	}
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DetailUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DetailUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DetailUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DetailUpdate) defaults() error {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		if detail.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized detail.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := detail.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (du *DetailUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DetailUpdate {
	du.modifiers = append(du.modifiers, modifiers...)
	return du
}

func (du *DetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   detail.Table,
			Columns: detail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: detail.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := du.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := du.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldEntID,
		})
	}
	if value, ok := du.mutation.SampleCol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detail.FieldSampleCol,
		})
	}
	if du.mutation.SampleColCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: detail.FieldSampleCol,
		})
	}
	_spec.Modifiers = du.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{detail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DetailUpdateOne is the builder for updating a single Detail entity.
type DetailUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DetailMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (duo *DetailUpdateOne) SetCreatedAt(u uint32) *DetailUpdateOne {
	duo.mutation.ResetCreatedAt()
	duo.mutation.SetCreatedAt(u)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableCreatedAt(u *uint32) *DetailUpdateOne {
	if u != nil {
		duo.SetCreatedAt(*u)
	}
	return duo
}

// AddCreatedAt adds u to the "created_at" field.
func (duo *DetailUpdateOne) AddCreatedAt(u int32) *DetailUpdateOne {
	duo.mutation.AddCreatedAt(u)
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DetailUpdateOne) SetUpdatedAt(u uint32) *DetailUpdateOne {
	duo.mutation.ResetUpdatedAt()
	duo.mutation.SetUpdatedAt(u)
	return duo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (duo *DetailUpdateOne) AddUpdatedAt(u int32) *DetailUpdateOne {
	duo.mutation.AddUpdatedAt(u)
	return duo
}

// SetDeletedAt sets the "deleted_at" field.
func (duo *DetailUpdateOne) SetDeletedAt(u uint32) *DetailUpdateOne {
	duo.mutation.ResetDeletedAt()
	duo.mutation.SetDeletedAt(u)
	return duo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableDeletedAt(u *uint32) *DetailUpdateOne {
	if u != nil {
		duo.SetDeletedAt(*u)
	}
	return duo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (duo *DetailUpdateOne) AddDeletedAt(u int32) *DetailUpdateOne {
	duo.mutation.AddDeletedAt(u)
	return duo
}

// SetEntID sets the "ent_id" field.
func (duo *DetailUpdateOne) SetEntID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetEntID(u)
	return duo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableEntID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetEntID(*u)
	}
	return duo
}

// SetSampleCol sets the "sample_col" field.
func (duo *DetailUpdateOne) SetSampleCol(s string) *DetailUpdateOne {
	duo.mutation.SetSampleCol(s)
	return duo
}

// SetNillableSampleCol sets the "sample_col" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableSampleCol(s *string) *DetailUpdateOne {
	if s != nil {
		duo.SetSampleCol(*s)
	}
	return duo
}

// ClearSampleCol clears the value of the "sample_col" field.
func (duo *DetailUpdateOne) ClearSampleCol() *DetailUpdateOne {
	duo.mutation.ClearSampleCol()
	return duo
}

// Mutation returns the DetailMutation object of the builder.
func (duo *DetailUpdateOne) Mutation() *DetailMutation {
	return duo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DetailUpdateOne) Select(field string, fields ...string) *DetailUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Detail entity.
func (duo *DetailUpdateOne) Save(ctx context.Context) (*Detail, error) {
	var (
		err  error
		node *Detail
	)
	if err := duo.defaults(); err != nil {
		return nil, err
	}
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Detail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DetailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DetailUpdateOne) SaveX(ctx context.Context) *Detail {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DetailUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DetailUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DetailUpdateOne) defaults() error {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		if detail.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized detail.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := detail.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (duo *DetailUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DetailUpdateOne {
	duo.modifiers = append(duo.modifiers, modifiers...)
	return duo
}

func (duo *DetailUpdateOne) sqlSave(ctx context.Context) (_node *Detail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   detail.Table,
			Columns: detail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: detail.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Detail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, detail.FieldID)
		for _, f := range fields {
			if !detail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != detail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := duo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := duo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldEntID,
		})
	}
	if value, ok := duo.mutation.SampleCol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detail.FieldSampleCol,
		})
	}
	if duo.mutation.SampleColCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: detail.FieldSampleCol,
		})
	}
	_spec.Modifiers = duo.modifiers
	_node = &Detail{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{detail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
