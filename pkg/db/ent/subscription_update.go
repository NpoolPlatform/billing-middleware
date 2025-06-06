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
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/subscription"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// SubscriptionUpdate is the builder for updating Subscription entities.
type SubscriptionUpdate struct {
	config
	hooks     []Hook
	mutation  *SubscriptionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SubscriptionUpdate builder.
func (su *SubscriptionUpdate) Where(ps ...predicate.Subscription) *SubscriptionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SubscriptionUpdate) SetCreatedAt(u uint32) *SubscriptionUpdate {
	su.mutation.ResetCreatedAt()
	su.mutation.SetCreatedAt(u)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableCreatedAt(u *uint32) *SubscriptionUpdate {
	if u != nil {
		su.SetCreatedAt(*u)
	}
	return su
}

// AddCreatedAt adds u to the "created_at" field.
func (su *SubscriptionUpdate) AddCreatedAt(u int32) *SubscriptionUpdate {
	su.mutation.AddCreatedAt(u)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SubscriptionUpdate) SetUpdatedAt(u uint32) *SubscriptionUpdate {
	su.mutation.ResetUpdatedAt()
	su.mutation.SetUpdatedAt(u)
	return su
}

// AddUpdatedAt adds u to the "updated_at" field.
func (su *SubscriptionUpdate) AddUpdatedAt(u int32) *SubscriptionUpdate {
	su.mutation.AddUpdatedAt(u)
	return su
}

// SetDeletedAt sets the "deleted_at" field.
func (su *SubscriptionUpdate) SetDeletedAt(u uint32) *SubscriptionUpdate {
	su.mutation.ResetDeletedAt()
	su.mutation.SetDeletedAt(u)
	return su
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableDeletedAt(u *uint32) *SubscriptionUpdate {
	if u != nil {
		su.SetDeletedAt(*u)
	}
	return su
}

// AddDeletedAt adds u to the "deleted_at" field.
func (su *SubscriptionUpdate) AddDeletedAt(u int32) *SubscriptionUpdate {
	su.mutation.AddDeletedAt(u)
	return su
}

// SetEntID sets the "ent_id" field.
func (su *SubscriptionUpdate) SetEntID(u uuid.UUID) *SubscriptionUpdate {
	su.mutation.SetEntID(u)
	return su
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableEntID(u *uuid.UUID) *SubscriptionUpdate {
	if u != nil {
		su.SetEntID(*u)
	}
	return su
}

// SetAppID sets the "app_id" field.
func (su *SubscriptionUpdate) SetAppID(u uuid.UUID) *SubscriptionUpdate {
	su.mutation.SetAppID(u)
	return su
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableAppID(u *uuid.UUID) *SubscriptionUpdate {
	if u != nil {
		su.SetAppID(*u)
	}
	return su
}

// ClearAppID clears the value of the "app_id" field.
func (su *SubscriptionUpdate) ClearAppID() *SubscriptionUpdate {
	su.mutation.ClearAppID()
	return su
}

// SetPackageName sets the "package_name" field.
func (su *SubscriptionUpdate) SetPackageName(s string) *SubscriptionUpdate {
	su.mutation.SetPackageName(s)
	return su
}

// SetNillablePackageName sets the "package_name" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillablePackageName(s *string) *SubscriptionUpdate {
	if s != nil {
		su.SetPackageName(*s)
	}
	return su
}

// ClearPackageName clears the value of the "package_name" field.
func (su *SubscriptionUpdate) ClearPackageName() *SubscriptionUpdate {
	su.mutation.ClearPackageName()
	return su
}

// SetUsdPrice sets the "usd_price" field.
func (su *SubscriptionUpdate) SetUsdPrice(d decimal.Decimal) *SubscriptionUpdate {
	su.mutation.SetUsdPrice(d)
	return su
}

// SetNillableUsdPrice sets the "usd_price" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableUsdPrice(d *decimal.Decimal) *SubscriptionUpdate {
	if d != nil {
		su.SetUsdPrice(*d)
	}
	return su
}

// ClearUsdPrice clears the value of the "usd_price" field.
func (su *SubscriptionUpdate) ClearUsdPrice() *SubscriptionUpdate {
	su.mutation.ClearUsdPrice()
	return su
}

// SetDescription sets the "description" field.
func (su *SubscriptionUpdate) SetDescription(s string) *SubscriptionUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableDescription(s *string) *SubscriptionUpdate {
	if s != nil {
		su.SetDescription(*s)
	}
	return su
}

// ClearDescription clears the value of the "description" field.
func (su *SubscriptionUpdate) ClearDescription() *SubscriptionUpdate {
	su.mutation.ClearDescription()
	return su
}

// SetSortOrder sets the "sort_order" field.
func (su *SubscriptionUpdate) SetSortOrder(u uint32) *SubscriptionUpdate {
	su.mutation.ResetSortOrder()
	su.mutation.SetSortOrder(u)
	return su
}

// SetNillableSortOrder sets the "sort_order" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableSortOrder(u *uint32) *SubscriptionUpdate {
	if u != nil {
		su.SetSortOrder(*u)
	}
	return su
}

// AddSortOrder adds u to the "sort_order" field.
func (su *SubscriptionUpdate) AddSortOrder(u int32) *SubscriptionUpdate {
	su.mutation.AddSortOrder(u)
	return su
}

// ClearSortOrder clears the value of the "sort_order" field.
func (su *SubscriptionUpdate) ClearSortOrder() *SubscriptionUpdate {
	su.mutation.ClearSortOrder()
	return su
}

// SetPackageType sets the "package_type" field.
func (su *SubscriptionUpdate) SetPackageType(s string) *SubscriptionUpdate {
	su.mutation.SetPackageType(s)
	return su
}

// SetNillablePackageType sets the "package_type" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillablePackageType(s *string) *SubscriptionUpdate {
	if s != nil {
		su.SetPackageType(*s)
	}
	return su
}

// ClearPackageType clears the value of the "package_type" field.
func (su *SubscriptionUpdate) ClearPackageType() *SubscriptionUpdate {
	su.mutation.ClearPackageType()
	return su
}

// SetCredit sets the "credit" field.
func (su *SubscriptionUpdate) SetCredit(u uint32) *SubscriptionUpdate {
	su.mutation.ResetCredit()
	su.mutation.SetCredit(u)
	return su
}

// SetNillableCredit sets the "credit" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableCredit(u *uint32) *SubscriptionUpdate {
	if u != nil {
		su.SetCredit(*u)
	}
	return su
}

// AddCredit adds u to the "credit" field.
func (su *SubscriptionUpdate) AddCredit(u int32) *SubscriptionUpdate {
	su.mutation.AddCredit(u)
	return su
}

// ClearCredit clears the value of the "credit" field.
func (su *SubscriptionUpdate) ClearCredit() *SubscriptionUpdate {
	su.mutation.ClearCredit()
	return su
}

// SetResetType sets the "reset_type" field.
func (su *SubscriptionUpdate) SetResetType(s string) *SubscriptionUpdate {
	su.mutation.SetResetType(s)
	return su
}

// SetNillableResetType sets the "reset_type" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableResetType(s *string) *SubscriptionUpdate {
	if s != nil {
		su.SetResetType(*s)
	}
	return su
}

// ClearResetType clears the value of the "reset_type" field.
func (su *SubscriptionUpdate) ClearResetType() *SubscriptionUpdate {
	su.mutation.ClearResetType()
	return su
}

// SetQPSLimit sets the "qps_limit" field.
func (su *SubscriptionUpdate) SetQPSLimit(u uint32) *SubscriptionUpdate {
	su.mutation.ResetQPSLimit()
	su.mutation.SetQPSLimit(u)
	return su
}

// SetNillableQPSLimit sets the "qps_limit" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableQPSLimit(u *uint32) *SubscriptionUpdate {
	if u != nil {
		su.SetQPSLimit(*u)
	}
	return su
}

// AddQPSLimit adds u to the "qps_limit" field.
func (su *SubscriptionUpdate) AddQPSLimit(u int32) *SubscriptionUpdate {
	su.mutation.AddQPSLimit(u)
	return su
}

// ClearQPSLimit clears the value of the "qps_limit" field.
func (su *SubscriptionUpdate) ClearQPSLimit() *SubscriptionUpdate {
	su.mutation.ClearQPSLimit()
	return su
}

// Mutation returns the SubscriptionMutation object of the builder.
func (su *SubscriptionUpdate) Mutation() *SubscriptionMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubscriptionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := su.defaults(); err != nil {
		return 0, err
	}
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
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

// defaults sets the default values of the builder before save.
func (su *SubscriptionUpdate) defaults() error {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		if subscription.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized subscription.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := subscription.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SubscriptionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SubscriptionUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscription.Table,
			Columns: subscription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: subscription.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldDeletedAt,
		})
	}
	if value, ok := su.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldDeletedAt,
		})
	}
	if value, ok := su.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subscription.FieldEntID,
		})
	}
	if value, ok := su.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subscription.FieldAppID,
		})
	}
	if su.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: subscription.FieldAppID,
		})
	}
	if value, ok := su.mutation.PackageName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldPackageName,
		})
	}
	if su.mutation.PackageNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscription.FieldPackageName,
		})
	}
	if value, ok := su.mutation.UsdPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: subscription.FieldUsdPrice,
		})
	}
	if su.mutation.UsdPriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: subscription.FieldUsdPrice,
		})
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldDescription,
		})
	}
	if su.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscription.FieldDescription,
		})
	}
	if value, ok := su.mutation.SortOrder(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldSortOrder,
		})
	}
	if value, ok := su.mutation.AddedSortOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldSortOrder,
		})
	}
	if su.mutation.SortOrderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: subscription.FieldSortOrder,
		})
	}
	if value, ok := su.mutation.PackageType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldPackageType,
		})
	}
	if su.mutation.PackageTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscription.FieldPackageType,
		})
	}
	if value, ok := su.mutation.Credit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldCredit,
		})
	}
	if value, ok := su.mutation.AddedCredit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldCredit,
		})
	}
	if su.mutation.CreditCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: subscription.FieldCredit,
		})
	}
	if value, ok := su.mutation.ResetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldResetType,
		})
	}
	if su.mutation.ResetTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscription.FieldResetType,
		})
	}
	if value, ok := su.mutation.QPSLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldQPSLimit,
		})
	}
	if value, ok := su.mutation.AddedQPSLimit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldQPSLimit,
		})
	}
	if su.mutation.QPSLimitCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: subscription.FieldQPSLimit,
		})
	}
	_spec.Modifiers = su.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SubscriptionUpdateOne is the builder for updating a single Subscription entity.
type SubscriptionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SubscriptionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (suo *SubscriptionUpdateOne) SetCreatedAt(u uint32) *SubscriptionUpdateOne {
	suo.mutation.ResetCreatedAt()
	suo.mutation.SetCreatedAt(u)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableCreatedAt(u *uint32) *SubscriptionUpdateOne {
	if u != nil {
		suo.SetCreatedAt(*u)
	}
	return suo
}

// AddCreatedAt adds u to the "created_at" field.
func (suo *SubscriptionUpdateOne) AddCreatedAt(u int32) *SubscriptionUpdateOne {
	suo.mutation.AddCreatedAt(u)
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SubscriptionUpdateOne) SetUpdatedAt(u uint32) *SubscriptionUpdateOne {
	suo.mutation.ResetUpdatedAt()
	suo.mutation.SetUpdatedAt(u)
	return suo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (suo *SubscriptionUpdateOne) AddUpdatedAt(u int32) *SubscriptionUpdateOne {
	suo.mutation.AddUpdatedAt(u)
	return suo
}

// SetDeletedAt sets the "deleted_at" field.
func (suo *SubscriptionUpdateOne) SetDeletedAt(u uint32) *SubscriptionUpdateOne {
	suo.mutation.ResetDeletedAt()
	suo.mutation.SetDeletedAt(u)
	return suo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableDeletedAt(u *uint32) *SubscriptionUpdateOne {
	if u != nil {
		suo.SetDeletedAt(*u)
	}
	return suo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (suo *SubscriptionUpdateOne) AddDeletedAt(u int32) *SubscriptionUpdateOne {
	suo.mutation.AddDeletedAt(u)
	return suo
}

// SetEntID sets the "ent_id" field.
func (suo *SubscriptionUpdateOne) SetEntID(u uuid.UUID) *SubscriptionUpdateOne {
	suo.mutation.SetEntID(u)
	return suo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableEntID(u *uuid.UUID) *SubscriptionUpdateOne {
	if u != nil {
		suo.SetEntID(*u)
	}
	return suo
}

// SetAppID sets the "app_id" field.
func (suo *SubscriptionUpdateOne) SetAppID(u uuid.UUID) *SubscriptionUpdateOne {
	suo.mutation.SetAppID(u)
	return suo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableAppID(u *uuid.UUID) *SubscriptionUpdateOne {
	if u != nil {
		suo.SetAppID(*u)
	}
	return suo
}

// ClearAppID clears the value of the "app_id" field.
func (suo *SubscriptionUpdateOne) ClearAppID() *SubscriptionUpdateOne {
	suo.mutation.ClearAppID()
	return suo
}

// SetPackageName sets the "package_name" field.
func (suo *SubscriptionUpdateOne) SetPackageName(s string) *SubscriptionUpdateOne {
	suo.mutation.SetPackageName(s)
	return suo
}

// SetNillablePackageName sets the "package_name" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillablePackageName(s *string) *SubscriptionUpdateOne {
	if s != nil {
		suo.SetPackageName(*s)
	}
	return suo
}

// ClearPackageName clears the value of the "package_name" field.
func (suo *SubscriptionUpdateOne) ClearPackageName() *SubscriptionUpdateOne {
	suo.mutation.ClearPackageName()
	return suo
}

// SetUsdPrice sets the "usd_price" field.
func (suo *SubscriptionUpdateOne) SetUsdPrice(d decimal.Decimal) *SubscriptionUpdateOne {
	suo.mutation.SetUsdPrice(d)
	return suo
}

// SetNillableUsdPrice sets the "usd_price" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableUsdPrice(d *decimal.Decimal) *SubscriptionUpdateOne {
	if d != nil {
		suo.SetUsdPrice(*d)
	}
	return suo
}

// ClearUsdPrice clears the value of the "usd_price" field.
func (suo *SubscriptionUpdateOne) ClearUsdPrice() *SubscriptionUpdateOne {
	suo.mutation.ClearUsdPrice()
	return suo
}

// SetDescription sets the "description" field.
func (suo *SubscriptionUpdateOne) SetDescription(s string) *SubscriptionUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableDescription(s *string) *SubscriptionUpdateOne {
	if s != nil {
		suo.SetDescription(*s)
	}
	return suo
}

// ClearDescription clears the value of the "description" field.
func (suo *SubscriptionUpdateOne) ClearDescription() *SubscriptionUpdateOne {
	suo.mutation.ClearDescription()
	return suo
}

// SetSortOrder sets the "sort_order" field.
func (suo *SubscriptionUpdateOne) SetSortOrder(u uint32) *SubscriptionUpdateOne {
	suo.mutation.ResetSortOrder()
	suo.mutation.SetSortOrder(u)
	return suo
}

// SetNillableSortOrder sets the "sort_order" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableSortOrder(u *uint32) *SubscriptionUpdateOne {
	if u != nil {
		suo.SetSortOrder(*u)
	}
	return suo
}

// AddSortOrder adds u to the "sort_order" field.
func (suo *SubscriptionUpdateOne) AddSortOrder(u int32) *SubscriptionUpdateOne {
	suo.mutation.AddSortOrder(u)
	return suo
}

// ClearSortOrder clears the value of the "sort_order" field.
func (suo *SubscriptionUpdateOne) ClearSortOrder() *SubscriptionUpdateOne {
	suo.mutation.ClearSortOrder()
	return suo
}

// SetPackageType sets the "package_type" field.
func (suo *SubscriptionUpdateOne) SetPackageType(s string) *SubscriptionUpdateOne {
	suo.mutation.SetPackageType(s)
	return suo
}

// SetNillablePackageType sets the "package_type" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillablePackageType(s *string) *SubscriptionUpdateOne {
	if s != nil {
		suo.SetPackageType(*s)
	}
	return suo
}

// ClearPackageType clears the value of the "package_type" field.
func (suo *SubscriptionUpdateOne) ClearPackageType() *SubscriptionUpdateOne {
	suo.mutation.ClearPackageType()
	return suo
}

// SetCredit sets the "credit" field.
func (suo *SubscriptionUpdateOne) SetCredit(u uint32) *SubscriptionUpdateOne {
	suo.mutation.ResetCredit()
	suo.mutation.SetCredit(u)
	return suo
}

// SetNillableCredit sets the "credit" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableCredit(u *uint32) *SubscriptionUpdateOne {
	if u != nil {
		suo.SetCredit(*u)
	}
	return suo
}

// AddCredit adds u to the "credit" field.
func (suo *SubscriptionUpdateOne) AddCredit(u int32) *SubscriptionUpdateOne {
	suo.mutation.AddCredit(u)
	return suo
}

// ClearCredit clears the value of the "credit" field.
func (suo *SubscriptionUpdateOne) ClearCredit() *SubscriptionUpdateOne {
	suo.mutation.ClearCredit()
	return suo
}

// SetResetType sets the "reset_type" field.
func (suo *SubscriptionUpdateOne) SetResetType(s string) *SubscriptionUpdateOne {
	suo.mutation.SetResetType(s)
	return suo
}

// SetNillableResetType sets the "reset_type" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableResetType(s *string) *SubscriptionUpdateOne {
	if s != nil {
		suo.SetResetType(*s)
	}
	return suo
}

// ClearResetType clears the value of the "reset_type" field.
func (suo *SubscriptionUpdateOne) ClearResetType() *SubscriptionUpdateOne {
	suo.mutation.ClearResetType()
	return suo
}

// SetQPSLimit sets the "qps_limit" field.
func (suo *SubscriptionUpdateOne) SetQPSLimit(u uint32) *SubscriptionUpdateOne {
	suo.mutation.ResetQPSLimit()
	suo.mutation.SetQPSLimit(u)
	return suo
}

// SetNillableQPSLimit sets the "qps_limit" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableQPSLimit(u *uint32) *SubscriptionUpdateOne {
	if u != nil {
		suo.SetQPSLimit(*u)
	}
	return suo
}

// AddQPSLimit adds u to the "qps_limit" field.
func (suo *SubscriptionUpdateOne) AddQPSLimit(u int32) *SubscriptionUpdateOne {
	suo.mutation.AddQPSLimit(u)
	return suo
}

// ClearQPSLimit clears the value of the "qps_limit" field.
func (suo *SubscriptionUpdateOne) ClearQPSLimit() *SubscriptionUpdateOne {
	suo.mutation.ClearQPSLimit()
	return suo
}

// Mutation returns the SubscriptionMutation object of the builder.
func (suo *SubscriptionUpdateOne) Mutation() *SubscriptionMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubscriptionUpdateOne) Select(field string, fields ...string) *SubscriptionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subscription entity.
func (suo *SubscriptionUpdateOne) Save(ctx context.Context) (*Subscription, error) {
	var (
		err  error
		node *Subscription
	)
	if err := suo.defaults(); err != nil {
		return nil, err
	}
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Subscription)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SubscriptionMutation", v)
		}
		node = nv
	}
	return node, err
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

// defaults sets the default values of the builder before save.
func (suo *SubscriptionUpdateOne) defaults() error {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		if subscription.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized subscription.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := subscription.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SubscriptionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SubscriptionUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *Subscription, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscription.Table,
			Columns: subscription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: subscription.FieldID,
			},
		},
	}
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
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldDeletedAt,
		})
	}
	if value, ok := suo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldDeletedAt,
		})
	}
	if value, ok := suo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subscription.FieldEntID,
		})
	}
	if value, ok := suo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subscription.FieldAppID,
		})
	}
	if suo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: subscription.FieldAppID,
		})
	}
	if value, ok := suo.mutation.PackageName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldPackageName,
		})
	}
	if suo.mutation.PackageNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscription.FieldPackageName,
		})
	}
	if value, ok := suo.mutation.UsdPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: subscription.FieldUsdPrice,
		})
	}
	if suo.mutation.UsdPriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: subscription.FieldUsdPrice,
		})
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldDescription,
		})
	}
	if suo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscription.FieldDescription,
		})
	}
	if value, ok := suo.mutation.SortOrder(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldSortOrder,
		})
	}
	if value, ok := suo.mutation.AddedSortOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldSortOrder,
		})
	}
	if suo.mutation.SortOrderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: subscription.FieldSortOrder,
		})
	}
	if value, ok := suo.mutation.PackageType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldPackageType,
		})
	}
	if suo.mutation.PackageTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscription.FieldPackageType,
		})
	}
	if value, ok := suo.mutation.Credit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldCredit,
		})
	}
	if value, ok := suo.mutation.AddedCredit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldCredit,
		})
	}
	if suo.mutation.CreditCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: subscription.FieldCredit,
		})
	}
	if value, ok := suo.mutation.ResetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscription.FieldResetType,
		})
	}
	if suo.mutation.ResetTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscription.FieldResetType,
		})
	}
	if value, ok := suo.mutation.QPSLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldQPSLimit,
		})
	}
	if value, ok := suo.mutation.AddedQPSLimit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscription.FieldQPSLimit,
		})
	}
	if suo.mutation.QPSLimitCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: subscription.FieldQPSLimit,
		})
	}
	_spec.Modifiers = suo.modifiers
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
	return _node, nil
}
