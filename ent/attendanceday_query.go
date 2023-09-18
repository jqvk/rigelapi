// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// AttendanceDayQuery is the builder for querying AttendanceDay entities.
type AttendanceDayQuery struct {
	config
	ctx             *QueryContext
	order           []attendanceday.OrderOption
	inters          []Interceptor
	predicates      []predicate.AttendanceDay
	withAttendances *AttendanceQuery
	withClassPeriod *ClassPeriodQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttendanceDayQuery builder.
func (adq *AttendanceDayQuery) Where(ps ...predicate.AttendanceDay) *AttendanceDayQuery {
	adq.predicates = append(adq.predicates, ps...)
	return adq
}

// Limit the number of records to be returned by this query.
func (adq *AttendanceDayQuery) Limit(limit int) *AttendanceDayQuery {
	adq.ctx.Limit = &limit
	return adq
}

// Offset to start from.
func (adq *AttendanceDayQuery) Offset(offset int) *AttendanceDayQuery {
	adq.ctx.Offset = &offset
	return adq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (adq *AttendanceDayQuery) Unique(unique bool) *AttendanceDayQuery {
	adq.ctx.Unique = &unique
	return adq
}

// Order specifies how the records should be ordered.
func (adq *AttendanceDayQuery) Order(o ...attendanceday.OrderOption) *AttendanceDayQuery {
	adq.order = append(adq.order, o...)
	return adq
}

// QueryAttendances chains the current query on the "attendances" edge.
func (adq *AttendanceDayQuery) QueryAttendances() *AttendanceQuery {
	query := (&AttendanceClient{config: adq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := adq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := adq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attendanceday.Table, attendanceday.FieldID, selector),
			sqlgraph.To(attendance.Table, attendance.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, attendanceday.AttendancesTable, attendanceday.AttendancesColumn),
		)
		fromU = sqlgraph.SetNeighbors(adq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClassPeriod chains the current query on the "classPeriod" edge.
func (adq *AttendanceDayQuery) QueryClassPeriod() *ClassPeriodQuery {
	query := (&ClassPeriodClient{config: adq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := adq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := adq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attendanceday.Table, attendanceday.FieldID, selector),
			sqlgraph.To(classperiod.Table, classperiod.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attendanceday.ClassPeriodTable, attendanceday.ClassPeriodColumn),
		)
		fromU = sqlgraph.SetNeighbors(adq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AttendanceDay entity from the query.
// Returns a *NotFoundError when no AttendanceDay was found.
func (adq *AttendanceDayQuery) First(ctx context.Context) (*AttendanceDay, error) {
	nodes, err := adq.Limit(1).All(setContextOp(ctx, adq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attendanceday.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (adq *AttendanceDayQuery) FirstX(ctx context.Context) *AttendanceDay {
	node, err := adq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttendanceDay ID from the query.
// Returns a *NotFoundError when no AttendanceDay ID was found.
func (adq *AttendanceDayQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = adq.Limit(1).IDs(setContextOp(ctx, adq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attendanceday.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (adq *AttendanceDayQuery) FirstIDX(ctx context.Context) string {
	id, err := adq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttendanceDay entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttendanceDay entity is found.
// Returns a *NotFoundError when no AttendanceDay entities are found.
func (adq *AttendanceDayQuery) Only(ctx context.Context) (*AttendanceDay, error) {
	nodes, err := adq.Limit(2).All(setContextOp(ctx, adq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attendanceday.Label}
	default:
		return nil, &NotSingularError{attendanceday.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (adq *AttendanceDayQuery) OnlyX(ctx context.Context) *AttendanceDay {
	node, err := adq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttendanceDay ID in the query.
// Returns a *NotSingularError when more than one AttendanceDay ID is found.
// Returns a *NotFoundError when no entities are found.
func (adq *AttendanceDayQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = adq.Limit(2).IDs(setContextOp(ctx, adq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attendanceday.Label}
	default:
		err = &NotSingularError{attendanceday.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (adq *AttendanceDayQuery) OnlyIDX(ctx context.Context) string {
	id, err := adq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttendanceDays.
func (adq *AttendanceDayQuery) All(ctx context.Context) ([]*AttendanceDay, error) {
	ctx = setContextOp(ctx, adq.ctx, "All")
	if err := adq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AttendanceDay, *AttendanceDayQuery]()
	return withInterceptors[[]*AttendanceDay](ctx, adq, qr, adq.inters)
}

// AllX is like All, but panics if an error occurs.
func (adq *AttendanceDayQuery) AllX(ctx context.Context) []*AttendanceDay {
	nodes, err := adq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttendanceDay IDs.
func (adq *AttendanceDayQuery) IDs(ctx context.Context) (ids []string, err error) {
	if adq.ctx.Unique == nil && adq.path != nil {
		adq.Unique(true)
	}
	ctx = setContextOp(ctx, adq.ctx, "IDs")
	if err = adq.Select(attendanceday.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (adq *AttendanceDayQuery) IDsX(ctx context.Context) []string {
	ids, err := adq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (adq *AttendanceDayQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, adq.ctx, "Count")
	if err := adq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, adq, querierCount[*AttendanceDayQuery](), adq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (adq *AttendanceDayQuery) CountX(ctx context.Context) int {
	count, err := adq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (adq *AttendanceDayQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, adq.ctx, "Exist")
	switch _, err := adq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (adq *AttendanceDayQuery) ExistX(ctx context.Context) bool {
	exist, err := adq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttendanceDayQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (adq *AttendanceDayQuery) Clone() *AttendanceDayQuery {
	if adq == nil {
		return nil
	}
	return &AttendanceDayQuery{
		config:          adq.config,
		ctx:             adq.ctx.Clone(),
		order:           append([]attendanceday.OrderOption{}, adq.order...),
		inters:          append([]Interceptor{}, adq.inters...),
		predicates:      append([]predicate.AttendanceDay{}, adq.predicates...),
		withAttendances: adq.withAttendances.Clone(),
		withClassPeriod: adq.withClassPeriod.Clone(),
		// clone intermediate query.
		sql:  adq.sql.Clone(),
		path: adq.path,
	}
}

// WithAttendances tells the query-builder to eager-load the nodes that are connected to
// the "attendances" edge. The optional arguments are used to configure the query builder of the edge.
func (adq *AttendanceDayQuery) WithAttendances(opts ...func(*AttendanceQuery)) *AttendanceDayQuery {
	query := (&AttendanceClient{config: adq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	adq.withAttendances = query
	return adq
}

// WithClassPeriod tells the query-builder to eager-load the nodes that are connected to
// the "classPeriod" edge. The optional arguments are used to configure the query builder of the edge.
func (adq *AttendanceDayQuery) WithClassPeriod(opts ...func(*ClassPeriodQuery)) *AttendanceDayQuery {
	query := (&ClassPeriodClient{config: adq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	adq.withClassPeriod = query
	return adq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Day time.Time `json:"day,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AttendanceDay.Query().
//		GroupBy(attendanceday.FieldDay).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (adq *AttendanceDayQuery) GroupBy(field string, fields ...string) *AttendanceDayGroupBy {
	adq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttendanceDayGroupBy{build: adq}
	grbuild.flds = &adq.ctx.Fields
	grbuild.label = attendanceday.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Day time.Time `json:"day,omitempty"`
//	}
//
//	client.AttendanceDay.Query().
//		Select(attendanceday.FieldDay).
//		Scan(ctx, &v)
func (adq *AttendanceDayQuery) Select(fields ...string) *AttendanceDaySelect {
	adq.ctx.Fields = append(adq.ctx.Fields, fields...)
	sbuild := &AttendanceDaySelect{AttendanceDayQuery: adq}
	sbuild.label = attendanceday.Label
	sbuild.flds, sbuild.scan = &adq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttendanceDaySelect configured with the given aggregations.
func (adq *AttendanceDayQuery) Aggregate(fns ...AggregateFunc) *AttendanceDaySelect {
	return adq.Select().Aggregate(fns...)
}

func (adq *AttendanceDayQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range adq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, adq); err != nil {
				return err
			}
		}
	}
	for _, f := range adq.ctx.Fields {
		if !attendanceday.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if adq.path != nil {
		prev, err := adq.path(ctx)
		if err != nil {
			return err
		}
		adq.sql = prev
	}
	return nil
}

func (adq *AttendanceDayQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttendanceDay, error) {
	var (
		nodes       = []*AttendanceDay{}
		withFKs     = adq.withFKs
		_spec       = adq.querySpec()
		loadedTypes = [2]bool{
			adq.withAttendances != nil,
			adq.withClassPeriod != nil,
		}
	)
	if adq.withClassPeriod != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, attendanceday.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AttendanceDay).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AttendanceDay{config: adq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, adq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := adq.withAttendances; query != nil {
		if err := adq.loadAttendances(ctx, query, nodes,
			func(n *AttendanceDay) { n.Edges.Attendances = []*Attendance{} },
			func(n *AttendanceDay, e *Attendance) { n.Edges.Attendances = append(n.Edges.Attendances, e) }); err != nil {
			return nil, err
		}
	}
	if query := adq.withClassPeriod; query != nil {
		if err := adq.loadClassPeriod(ctx, query, nodes, nil,
			func(n *AttendanceDay, e *ClassPeriod) { n.Edges.ClassPeriod = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (adq *AttendanceDayQuery) loadAttendances(ctx context.Context, query *AttendanceQuery, nodes []*AttendanceDay, init func(*AttendanceDay), assign func(*AttendanceDay, *Attendance)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*AttendanceDay)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Attendance(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(attendanceday.AttendancesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.attendance_day_attendances
		if fk == nil {
			return fmt.Errorf(`foreign-key "attendance_day_attendances" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "attendance_day_attendances" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (adq *AttendanceDayQuery) loadClassPeriod(ctx context.Context, query *ClassPeriodQuery, nodes []*AttendanceDay, init func(*AttendanceDay), assign func(*AttendanceDay, *ClassPeriod)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*AttendanceDay)
	for i := range nodes {
		if nodes[i].class_period_attendance_days == nil {
			continue
		}
		fk := *nodes[i].class_period_attendance_days
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(classperiod.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_period_attendance_days" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (adq *AttendanceDayQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := adq.querySpec()
	_spec.Node.Columns = adq.ctx.Fields
	if len(adq.ctx.Fields) > 0 {
		_spec.Unique = adq.ctx.Unique != nil && *adq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, adq.driver, _spec)
}

func (adq *AttendanceDayQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attendanceday.Table, attendanceday.Columns, sqlgraph.NewFieldSpec(attendanceday.FieldID, field.TypeString))
	_spec.From = adq.sql
	if unique := adq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if adq.path != nil {
		_spec.Unique = true
	}
	if fields := adq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attendanceday.FieldID)
		for i := range fields {
			if fields[i] != attendanceday.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := adq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := adq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := adq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := adq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (adq *AttendanceDayQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(adq.driver.Dialect())
	t1 := builder.Table(attendanceday.Table)
	columns := adq.ctx.Fields
	if len(columns) == 0 {
		columns = attendanceday.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if adq.sql != nil {
		selector = adq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if adq.ctx.Unique != nil && *adq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range adq.predicates {
		p(selector)
	}
	for _, p := range adq.order {
		p(selector)
	}
	if offset := adq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := adq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttendanceDayGroupBy is the group-by builder for AttendanceDay entities.
type AttendanceDayGroupBy struct {
	selector
	build *AttendanceDayQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (adgb *AttendanceDayGroupBy) Aggregate(fns ...AggregateFunc) *AttendanceDayGroupBy {
	adgb.fns = append(adgb.fns, fns...)
	return adgb
}

// Scan applies the selector query and scans the result into the given value.
func (adgb *AttendanceDayGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, adgb.build.ctx, "GroupBy")
	if err := adgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttendanceDayQuery, *AttendanceDayGroupBy](ctx, adgb.build, adgb, adgb.build.inters, v)
}

func (adgb *AttendanceDayGroupBy) sqlScan(ctx context.Context, root *AttendanceDayQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(adgb.fns))
	for _, fn := range adgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*adgb.flds)+len(adgb.fns))
		for _, f := range *adgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*adgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := adgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttendanceDaySelect is the builder for selecting fields of AttendanceDay entities.
type AttendanceDaySelect struct {
	*AttendanceDayQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ads *AttendanceDaySelect) Aggregate(fns ...AggregateFunc) *AttendanceDaySelect {
	ads.fns = append(ads.fns, fns...)
	return ads
}

// Scan applies the selector query and scans the result into the given value.
func (ads *AttendanceDaySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ads.ctx, "Select")
	if err := ads.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttendanceDayQuery, *AttendanceDaySelect](ctx, ads.AttendanceDayQuery, ads, ads.inters, v)
}

func (ads *AttendanceDaySelect) sqlScan(ctx context.Context, root *AttendanceDayQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ads.fns))
	for _, fn := range ads.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ads.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ads.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
