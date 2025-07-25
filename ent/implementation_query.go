// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"requirements/ent/implementation"
	"requirements/ent/predicate"
	"requirements/ent/requirement"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ImplementationQuery is the builder for querying Implementation entities.
type ImplementationQuery struct {
	config
	ctx              *QueryContext
	order            []implementation.OrderOption
	inters           []Interceptor
	predicates       []predicate.Implementation
	withRequirements *RequirementQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImplementationQuery builder.
func (iq *ImplementationQuery) Where(ps ...predicate.Implementation) *ImplementationQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *ImplementationQuery) Limit(limit int) *ImplementationQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *ImplementationQuery) Offset(offset int) *ImplementationQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *ImplementationQuery) Unique(unique bool) *ImplementationQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *ImplementationQuery) Order(o ...implementation.OrderOption) *ImplementationQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryRequirements chains the current query on the "requirements" edge.
func (iq *ImplementationQuery) QueryRequirements() *RequirementQuery {
	query := (&RequirementClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(implementation.Table, implementation.FieldID, selector),
			sqlgraph.To(requirement.Table, requirement.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, implementation.RequirementsTable, implementation.RequirementsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Implementation entity from the query.
// Returns a *NotFoundError when no Implementation was found.
func (iq *ImplementationQuery) First(ctx context.Context) (*Implementation, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{implementation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *ImplementationQuery) FirstX(ctx context.Context) *Implementation {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Implementation ID from the query.
// Returns a *NotFoundError when no Implementation ID was found.
func (iq *ImplementationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{implementation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *ImplementationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Implementation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Implementation entity is found.
// Returns a *NotFoundError when no Implementation entities are found.
func (iq *ImplementationQuery) Only(ctx context.Context) (*Implementation, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{implementation.Label}
	default:
		return nil, &NotSingularError{implementation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *ImplementationQuery) OnlyX(ctx context.Context) *Implementation {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Implementation ID in the query.
// Returns a *NotSingularError when more than one Implementation ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *ImplementationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{implementation.Label}
	default:
		err = &NotSingularError{implementation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *ImplementationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Implementations.
func (iq *ImplementationQuery) All(ctx context.Context) ([]*Implementation, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryAll)
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Implementation, *ImplementationQuery]()
	return withInterceptors[[]*Implementation](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *ImplementationQuery) AllX(ctx context.Context) []*Implementation {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Implementation IDs.
func (iq *ImplementationQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryIDs)
	if err = iq.Select(implementation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *ImplementationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *ImplementationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryCount)
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*ImplementationQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *ImplementationQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *ImplementationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryExist)
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *ImplementationQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImplementationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *ImplementationQuery) Clone() *ImplementationQuery {
	if iq == nil {
		return nil
	}
	return &ImplementationQuery{
		config:           iq.config,
		ctx:              iq.ctx.Clone(),
		order:            append([]implementation.OrderOption{}, iq.order...),
		inters:           append([]Interceptor{}, iq.inters...),
		predicates:       append([]predicate.Implementation{}, iq.predicates...),
		withRequirements: iq.withRequirements.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithRequirements tells the query-builder to eager-load the nodes that are connected to
// the "requirements" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ImplementationQuery) WithRequirements(opts ...func(*RequirementQuery)) *ImplementationQuery {
	query := (&RequirementClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withRequirements = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		URL string `json:"url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Implementation.Query().
//		GroupBy(implementation.FieldURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *ImplementationQuery) GroupBy(field string, fields ...string) *ImplementationGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ImplementationGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = implementation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		URL string `json:"url,omitempty"`
//	}
//
//	client.Implementation.Query().
//		Select(implementation.FieldURL).
//		Scan(ctx, &v)
func (iq *ImplementationQuery) Select(fields ...string) *ImplementationSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &ImplementationSelect{ImplementationQuery: iq}
	sbuild.label = implementation.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ImplementationSelect configured with the given aggregations.
func (iq *ImplementationQuery) Aggregate(fns ...AggregateFunc) *ImplementationSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *ImplementationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !implementation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *ImplementationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Implementation, error) {
	var (
		nodes       = []*Implementation{}
		_spec       = iq.querySpec()
		loadedTypes = [1]bool{
			iq.withRequirements != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Implementation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Implementation{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withRequirements; query != nil {
		if err := iq.loadRequirements(ctx, query, nodes,
			func(n *Implementation) { n.Edges.Requirements = []*Requirement{} },
			func(n *Implementation, e *Requirement) { n.Edges.Requirements = append(n.Edges.Requirements, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *ImplementationQuery) loadRequirements(ctx context.Context, query *RequirementQuery, nodes []*Implementation, init func(*Implementation), assign func(*Implementation, *Requirement)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Implementation)
	nids := make(map[uuid.UUID]map[*Implementation]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(implementation.RequirementsTable)
		s.Join(joinT).On(s.C(requirement.FieldID), joinT.C(implementation.RequirementsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(implementation.RequirementsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(implementation.RequirementsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Implementation]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Requirement](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "requirements" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (iq *ImplementationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *ImplementationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(implementation.Table, implementation.Columns, sqlgraph.NewFieldSpec(implementation.FieldID, field.TypeUUID))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, implementation.FieldID)
		for i := range fields {
			if fields[i] != implementation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *ImplementationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(implementation.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = implementation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImplementationGroupBy is the group-by builder for Implementation entities.
type ImplementationGroupBy struct {
	selector
	build *ImplementationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *ImplementationGroupBy) Aggregate(fns ...AggregateFunc) *ImplementationGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *ImplementationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, ent.OpQueryGroupBy)
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImplementationQuery, *ImplementationGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *ImplementationGroupBy) sqlScan(ctx context.Context, root *ImplementationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ImplementationSelect is the builder for selecting fields of Implementation entities.
type ImplementationSelect struct {
	*ImplementationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *ImplementationSelect) Aggregate(fns ...AggregateFunc) *ImplementationSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *ImplementationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, ent.OpQuerySelect)
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImplementationQuery, *ImplementationSelect](ctx, is.ImplementationQuery, is, is.inters, v)
}

func (is *ImplementationSelect) sqlScan(ctx context.Context, root *ImplementationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
