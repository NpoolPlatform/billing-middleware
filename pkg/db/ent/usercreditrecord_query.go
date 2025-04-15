// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/usercreditrecord"
)

// UserCreditRecordQuery is the builder for querying UserCreditRecord entities.
type UserCreditRecordQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserCreditRecord
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserCreditRecordQuery builder.
func (ucrq *UserCreditRecordQuery) Where(ps ...predicate.UserCreditRecord) *UserCreditRecordQuery {
	ucrq.predicates = append(ucrq.predicates, ps...)
	return ucrq
}

// Limit adds a limit step to the query.
func (ucrq *UserCreditRecordQuery) Limit(limit int) *UserCreditRecordQuery {
	ucrq.limit = &limit
	return ucrq
}

// Offset adds an offset step to the query.
func (ucrq *UserCreditRecordQuery) Offset(offset int) *UserCreditRecordQuery {
	ucrq.offset = &offset
	return ucrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ucrq *UserCreditRecordQuery) Unique(unique bool) *UserCreditRecordQuery {
	ucrq.unique = &unique
	return ucrq
}

// Order adds an order step to the query.
func (ucrq *UserCreditRecordQuery) Order(o ...OrderFunc) *UserCreditRecordQuery {
	ucrq.order = append(ucrq.order, o...)
	return ucrq
}

// First returns the first UserCreditRecord entity from the query.
// Returns a *NotFoundError when no UserCreditRecord was found.
func (ucrq *UserCreditRecordQuery) First(ctx context.Context) (*UserCreditRecord, error) {
	nodes, err := ucrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usercreditrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ucrq *UserCreditRecordQuery) FirstX(ctx context.Context) *UserCreditRecord {
	node, err := ucrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserCreditRecord ID from the query.
// Returns a *NotFoundError when no UserCreditRecord ID was found.
func (ucrq *UserCreditRecordQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = ucrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usercreditrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ucrq *UserCreditRecordQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := ucrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserCreditRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserCreditRecord entity is found.
// Returns a *NotFoundError when no UserCreditRecord entities are found.
func (ucrq *UserCreditRecordQuery) Only(ctx context.Context) (*UserCreditRecord, error) {
	nodes, err := ucrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usercreditrecord.Label}
	default:
		return nil, &NotSingularError{usercreditrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ucrq *UserCreditRecordQuery) OnlyX(ctx context.Context) *UserCreditRecord {
	node, err := ucrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserCreditRecord ID in the query.
// Returns a *NotSingularError when more than one UserCreditRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (ucrq *UserCreditRecordQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = ucrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usercreditrecord.Label}
	default:
		err = &NotSingularError{usercreditrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ucrq *UserCreditRecordQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := ucrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserCreditRecords.
func (ucrq *UserCreditRecordQuery) All(ctx context.Context) ([]*UserCreditRecord, error) {
	if err := ucrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ucrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ucrq *UserCreditRecordQuery) AllX(ctx context.Context) []*UserCreditRecord {
	nodes, err := ucrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserCreditRecord IDs.
func (ucrq *UserCreditRecordQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := ucrq.Select(usercreditrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ucrq *UserCreditRecordQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := ucrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ucrq *UserCreditRecordQuery) Count(ctx context.Context) (int, error) {
	if err := ucrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ucrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ucrq *UserCreditRecordQuery) CountX(ctx context.Context) int {
	count, err := ucrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ucrq *UserCreditRecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := ucrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ucrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ucrq *UserCreditRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := ucrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserCreditRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ucrq *UserCreditRecordQuery) Clone() *UserCreditRecordQuery {
	if ucrq == nil {
		return nil
	}
	return &UserCreditRecordQuery{
		config:     ucrq.config,
		limit:      ucrq.limit,
		offset:     ucrq.offset,
		order:      append([]OrderFunc{}, ucrq.order...),
		predicates: append([]predicate.UserCreditRecord{}, ucrq.predicates...),
		// clone intermediate query.
		sql:    ucrq.sql.Clone(),
		path:   ucrq.path,
		unique: ucrq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserCreditRecord.Query().
//		GroupBy(usercreditrecord.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ucrq *UserCreditRecordQuery) GroupBy(field string, fields ...string) *UserCreditRecordGroupBy {
	grbuild := &UserCreditRecordGroupBy{config: ucrq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ucrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ucrq.sqlQuery(ctx), nil
	}
	grbuild.label = usercreditrecord.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.UserCreditRecord.Query().
//		Select(usercreditrecord.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ucrq *UserCreditRecordQuery) Select(fields ...string) *UserCreditRecordSelect {
	ucrq.fields = append(ucrq.fields, fields...)
	selbuild := &UserCreditRecordSelect{UserCreditRecordQuery: ucrq}
	selbuild.label = usercreditrecord.Label
	selbuild.flds, selbuild.scan = &ucrq.fields, selbuild.Scan
	return selbuild
}

func (ucrq *UserCreditRecordQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ucrq.fields {
		if !usercreditrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ucrq.path != nil {
		prev, err := ucrq.path(ctx)
		if err != nil {
			return err
		}
		ucrq.sql = prev
	}
	if usercreditrecord.Policy == nil {
		return errors.New("ent: uninitialized usercreditrecord.Policy (forgotten import ent/runtime?)")
	}
	if err := usercreditrecord.Policy.EvalQuery(ctx, ucrq); err != nil {
		return err
	}
	return nil
}

func (ucrq *UserCreditRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserCreditRecord, error) {
	var (
		nodes = []*UserCreditRecord{}
		_spec = ucrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserCreditRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserCreditRecord{config: ucrq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ucrq.modifiers) > 0 {
		_spec.Modifiers = ucrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ucrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ucrq *UserCreditRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ucrq.querySpec()
	if len(ucrq.modifiers) > 0 {
		_spec.Modifiers = ucrq.modifiers
	}
	_spec.Node.Columns = ucrq.fields
	if len(ucrq.fields) > 0 {
		_spec.Unique = ucrq.unique != nil && *ucrq.unique
	}
	return sqlgraph.CountNodes(ctx, ucrq.driver, _spec)
}

func (ucrq *UserCreditRecordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ucrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ucrq *UserCreditRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercreditrecord.Table,
			Columns: usercreditrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: usercreditrecord.FieldID,
			},
		},
		From:   ucrq.sql,
		Unique: true,
	}
	if unique := ucrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ucrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercreditrecord.FieldID)
		for i := range fields {
			if fields[i] != usercreditrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ucrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ucrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ucrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ucrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ucrq *UserCreditRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ucrq.driver.Dialect())
	t1 := builder.Table(usercreditrecord.Table)
	columns := ucrq.fields
	if len(columns) == 0 {
		columns = usercreditrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ucrq.sql != nil {
		selector = ucrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ucrq.unique != nil && *ucrq.unique {
		selector.Distinct()
	}
	for _, m := range ucrq.modifiers {
		m(selector)
	}
	for _, p := range ucrq.predicates {
		p(selector)
	}
	for _, p := range ucrq.order {
		p(selector)
	}
	if offset := ucrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ucrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ucrq *UserCreditRecordQuery) ForUpdate(opts ...sql.LockOption) *UserCreditRecordQuery {
	if ucrq.driver.Dialect() == dialect.Postgres {
		ucrq.Unique(false)
	}
	ucrq.modifiers = append(ucrq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ucrq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ucrq *UserCreditRecordQuery) ForShare(opts ...sql.LockOption) *UserCreditRecordQuery {
	if ucrq.driver.Dialect() == dialect.Postgres {
		ucrq.Unique(false)
	}
	ucrq.modifiers = append(ucrq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ucrq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ucrq *UserCreditRecordQuery) Modify(modifiers ...func(s *sql.Selector)) *UserCreditRecordSelect {
	ucrq.modifiers = append(ucrq.modifiers, modifiers...)
	return ucrq.Select()
}

// UserCreditRecordGroupBy is the group-by builder for UserCreditRecord entities.
type UserCreditRecordGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ucrgb *UserCreditRecordGroupBy) Aggregate(fns ...AggregateFunc) *UserCreditRecordGroupBy {
	ucrgb.fns = append(ucrgb.fns, fns...)
	return ucrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ucrgb *UserCreditRecordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ucrgb.path(ctx)
	if err != nil {
		return err
	}
	ucrgb.sql = query
	return ucrgb.sqlScan(ctx, v)
}

func (ucrgb *UserCreditRecordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ucrgb.fields {
		if !usercreditrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ucrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ucrgb *UserCreditRecordGroupBy) sqlQuery() *sql.Selector {
	selector := ucrgb.sql.Select()
	aggregation := make([]string, 0, len(ucrgb.fns))
	for _, fn := range ucrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ucrgb.fields)+len(ucrgb.fns))
		for _, f := range ucrgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ucrgb.fields...)...)
}

// UserCreditRecordSelect is the builder for selecting fields of UserCreditRecord entities.
type UserCreditRecordSelect struct {
	*UserCreditRecordQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ucrs *UserCreditRecordSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ucrs.prepareQuery(ctx); err != nil {
		return err
	}
	ucrs.sql = ucrs.UserCreditRecordQuery.sqlQuery(ctx)
	return ucrs.sqlScan(ctx, v)
}

func (ucrs *UserCreditRecordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ucrs.sql.Query()
	if err := ucrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ucrs *UserCreditRecordSelect) Modify(modifiers ...func(s *sql.Selector)) *UserCreditRecordSelect {
	ucrs.modifiers = append(ucrs.modifiers, modifiers...)
	return ucrs
}
