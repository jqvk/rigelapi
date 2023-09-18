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
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/grade"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/school"
	"github.com/vmkevv/rigelapi/ent/student"
	"github.com/vmkevv/rigelapi/ent/subject"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

// ClassQuery is the builder for querying Class entities.
type ClassQuery struct {
	config
	ctx              *QueryContext
	order            []class.OrderOption
	inters           []Interceptor
	predicates       []predicate.Class
	withStudents     *StudentQuery
	withClassPeriods *ClassPeriodQuery
	withSchool       *SchoolQuery
	withTeacher      *TeacherQuery
	withSubject      *SubjectQuery
	withGrade        *GradeQuery
	withYear         *YearQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClassQuery builder.
func (cq *ClassQuery) Where(ps ...predicate.Class) *ClassQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ClassQuery) Limit(limit int) *ClassQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ClassQuery) Offset(offset int) *ClassQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ClassQuery) Unique(unique bool) *ClassQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ClassQuery) Order(o ...class.OrderOption) *ClassQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryStudents chains the current query on the "students" edge.
func (cq *ClassQuery) QueryStudents() *StudentQuery {
	query := (&StudentClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, class.StudentsTable, class.StudentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClassPeriods chains the current query on the "classPeriods" edge.
func (cq *ClassQuery) QueryClassPeriods() *ClassPeriodQuery {
	query := (&ClassPeriodClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(classperiod.Table, classperiod.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, class.ClassPeriodsTable, class.ClassPeriodsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySchool chains the current query on the "school" edge.
func (cq *ClassQuery) QuerySchool() *SchoolQuery {
	query := (&SchoolClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(school.Table, school.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, class.SchoolTable, class.SchoolColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeacher chains the current query on the "teacher" edge.
func (cq *ClassQuery) QueryTeacher() *TeacherQuery {
	query := (&TeacherClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(teacher.Table, teacher.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, class.TeacherTable, class.TeacherColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubject chains the current query on the "subject" edge.
func (cq *ClassQuery) QuerySubject() *SubjectQuery {
	query := (&SubjectClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, class.SubjectTable, class.SubjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGrade chains the current query on the "grade" edge.
func (cq *ClassQuery) QueryGrade() *GradeQuery {
	query := (&GradeClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(grade.Table, grade.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, class.GradeTable, class.GradeColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryYear chains the current query on the "year" edge.
func (cq *ClassQuery) QueryYear() *YearQuery {
	query := (&YearClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(year.Table, year.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, class.YearTable, class.YearColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Class entity from the query.
// Returns a *NotFoundError when no Class was found.
func (cq *ClassQuery) First(ctx context.Context) (*Class, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{class.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ClassQuery) FirstX(ctx context.Context) *Class {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Class ID from the query.
// Returns a *NotFoundError when no Class ID was found.
func (cq *ClassQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{class.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ClassQuery) FirstIDX(ctx context.Context) string {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Class entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Class entity is found.
// Returns a *NotFoundError when no Class entities are found.
func (cq *ClassQuery) Only(ctx context.Context) (*Class, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{class.Label}
	default:
		return nil, &NotSingularError{class.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ClassQuery) OnlyX(ctx context.Context) *Class {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Class ID in the query.
// Returns a *NotSingularError when more than one Class ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ClassQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{class.Label}
	default:
		err = &NotSingularError{class.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ClassQuery) OnlyIDX(ctx context.Context) string {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Classes.
func (cq *ClassQuery) All(ctx context.Context) ([]*Class, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Class, *ClassQuery]()
	return withInterceptors[[]*Class](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ClassQuery) AllX(ctx context.Context) []*Class {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Class IDs.
func (cq *ClassQuery) IDs(ctx context.Context) (ids []string, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(class.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ClassQuery) IDsX(ctx context.Context) []string {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ClassQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ClassQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ClassQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ClassQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ClassQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClassQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ClassQuery) Clone() *ClassQuery {
	if cq == nil {
		return nil
	}
	return &ClassQuery{
		config:           cq.config,
		ctx:              cq.ctx.Clone(),
		order:            append([]class.OrderOption{}, cq.order...),
		inters:           append([]Interceptor{}, cq.inters...),
		predicates:       append([]predicate.Class{}, cq.predicates...),
		withStudents:     cq.withStudents.Clone(),
		withClassPeriods: cq.withClassPeriods.Clone(),
		withSchool:       cq.withSchool.Clone(),
		withTeacher:      cq.withTeacher.Clone(),
		withSubject:      cq.withSubject.Clone(),
		withGrade:        cq.withGrade.Clone(),
		withYear:         cq.withYear.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithStudents tells the query-builder to eager-load the nodes that are connected to
// the "students" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithStudents(opts ...func(*StudentQuery)) *ClassQuery {
	query := (&StudentClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withStudents = query
	return cq
}

// WithClassPeriods tells the query-builder to eager-load the nodes that are connected to
// the "classPeriods" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithClassPeriods(opts ...func(*ClassPeriodQuery)) *ClassQuery {
	query := (&ClassPeriodClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withClassPeriods = query
	return cq
}

// WithSchool tells the query-builder to eager-load the nodes that are connected to
// the "school" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithSchool(opts ...func(*SchoolQuery)) *ClassQuery {
	query := (&SchoolClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withSchool = query
	return cq
}

// WithTeacher tells the query-builder to eager-load the nodes that are connected to
// the "teacher" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithTeacher(opts ...func(*TeacherQuery)) *ClassQuery {
	query := (&TeacherClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withTeacher = query
	return cq
}

// WithSubject tells the query-builder to eager-load the nodes that are connected to
// the "subject" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithSubject(opts ...func(*SubjectQuery)) *ClassQuery {
	query := (&SubjectClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withSubject = query
	return cq
}

// WithGrade tells the query-builder to eager-load the nodes that are connected to
// the "grade" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithGrade(opts ...func(*GradeQuery)) *ClassQuery {
	query := (&GradeClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withGrade = query
	return cq
}

// WithYear tells the query-builder to eager-load the nodes that are connected to
// the "year" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithYear(opts ...func(*YearQuery)) *ClassQuery {
	query := (&YearClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withYear = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Parallel string `json:"parallel,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Class.Query().
//		GroupBy(class.FieldParallel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ClassQuery) GroupBy(field string, fields ...string) *ClassGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ClassGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = class.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Parallel string `json:"parallel,omitempty"`
//	}
//
//	client.Class.Query().
//		Select(class.FieldParallel).
//		Scan(ctx, &v)
func (cq *ClassQuery) Select(fields ...string) *ClassSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ClassSelect{ClassQuery: cq}
	sbuild.label = class.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ClassSelect configured with the given aggregations.
func (cq *ClassQuery) Aggregate(fns ...AggregateFunc) *ClassSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ClassQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !class.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ClassQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Class, error) {
	var (
		nodes       = []*Class{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [7]bool{
			cq.withStudents != nil,
			cq.withClassPeriods != nil,
			cq.withSchool != nil,
			cq.withTeacher != nil,
			cq.withSubject != nil,
			cq.withGrade != nil,
			cq.withYear != nil,
		}
	)
	if cq.withSchool != nil || cq.withTeacher != nil || cq.withSubject != nil || cq.withGrade != nil || cq.withYear != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, class.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Class).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Class{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withStudents; query != nil {
		if err := cq.loadStudents(ctx, query, nodes,
			func(n *Class) { n.Edges.Students = []*Student{} },
			func(n *Class, e *Student) { n.Edges.Students = append(n.Edges.Students, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withClassPeriods; query != nil {
		if err := cq.loadClassPeriods(ctx, query, nodes,
			func(n *Class) { n.Edges.ClassPeriods = []*ClassPeriod{} },
			func(n *Class, e *ClassPeriod) { n.Edges.ClassPeriods = append(n.Edges.ClassPeriods, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withSchool; query != nil {
		if err := cq.loadSchool(ctx, query, nodes, nil,
			func(n *Class, e *School) { n.Edges.School = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withTeacher; query != nil {
		if err := cq.loadTeacher(ctx, query, nodes, nil,
			func(n *Class, e *Teacher) { n.Edges.Teacher = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withSubject; query != nil {
		if err := cq.loadSubject(ctx, query, nodes, nil,
			func(n *Class, e *Subject) { n.Edges.Subject = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withGrade; query != nil {
		if err := cq.loadGrade(ctx, query, nodes, nil,
			func(n *Class, e *Grade) { n.Edges.Grade = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withYear; query != nil {
		if err := cq.loadYear(ctx, query, nodes, nil,
			func(n *Class, e *Year) { n.Edges.Year = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ClassQuery) loadStudents(ctx context.Context, query *StudentQuery, nodes []*Class, init func(*Class), assign func(*Class, *Student)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Class)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Student(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(class.StudentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.class_students
		if fk == nil {
			return fmt.Errorf(`foreign-key "class_students" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "class_students" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ClassQuery) loadClassPeriods(ctx context.Context, query *ClassPeriodQuery, nodes []*Class, init func(*Class), assign func(*Class, *ClassPeriod)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Class)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ClassPeriod(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(class.ClassPeriodsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.class_class_periods
		if fk == nil {
			return fmt.Errorf(`foreign-key "class_class_periods" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "class_class_periods" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ClassQuery) loadSchool(ctx context.Context, query *SchoolQuery, nodes []*Class, init func(*Class), assign func(*Class, *School)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Class)
	for i := range nodes {
		if nodes[i].school_classes == nil {
			continue
		}
		fk := *nodes[i].school_classes
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(school.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "school_classes" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ClassQuery) loadTeacher(ctx context.Context, query *TeacherQuery, nodes []*Class, init func(*Class), assign func(*Class, *Teacher)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Class)
	for i := range nodes {
		if nodes[i].teacher_classes == nil {
			continue
		}
		fk := *nodes[i].teacher_classes
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(teacher.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "teacher_classes" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ClassQuery) loadSubject(ctx context.Context, query *SubjectQuery, nodes []*Class, init func(*Class), assign func(*Class, *Subject)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Class)
	for i := range nodes {
		if nodes[i].subject_classes == nil {
			continue
		}
		fk := *nodes[i].subject_classes
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(subject.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "subject_classes" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ClassQuery) loadGrade(ctx context.Context, query *GradeQuery, nodes []*Class, init func(*Class), assign func(*Class, *Grade)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Class)
	for i := range nodes {
		if nodes[i].grade_classes == nil {
			continue
		}
		fk := *nodes[i].grade_classes
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(grade.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "grade_classes" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ClassQuery) loadYear(ctx context.Context, query *YearQuery, nodes []*Class, init func(*Class), assign func(*Class, *Year)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Class)
	for i := range nodes {
		if nodes[i].year_classes == nil {
			continue
		}
		fk := *nodes[i].year_classes
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(year.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "year_classes" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *ClassQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ClassQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(class.Table, class.Columns, sqlgraph.NewFieldSpec(class.FieldID, field.TypeString))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, class.FieldID)
		for i := range fields {
			if fields[i] != class.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ClassQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(class.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = class.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClassGroupBy is the group-by builder for Class entities.
type ClassGroupBy struct {
	selector
	build *ClassQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ClassGroupBy) Aggregate(fns ...AggregateFunc) *ClassGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ClassGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassQuery, *ClassGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ClassGroupBy) sqlScan(ctx context.Context, root *ClassQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ClassSelect is the builder for selecting fields of Class entities.
type ClassSelect struct {
	*ClassQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ClassSelect) Aggregate(fns ...AggregateFunc) *ClassSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ClassSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassQuery, *ClassSelect](ctx, cs.ClassQuery, cs, cs.inters, v)
}

func (cs *ClassSelect) sqlScan(ctx context.Context, root *ClassQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
