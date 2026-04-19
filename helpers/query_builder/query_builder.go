// Package querybuilder provides a type-safe, fluent SQL query builder.
// Each operation (SELECT, INSERT, UPDATE, DELETE) returns its own builder
// type, so invalid method chains (e.g. Values() without Insert()) are caught
// at compile time rather than at runtime.
//
// Placeholder style: PostgreSQL positional ($1, $2, …).
// Write conditions with ? placeholders and pass values as extra arguments;
// the builder rewrites each ? to the correct $n automatically:
//
//	sb.Where("age > ? AND active = ?", 18, true)
package query_builder

import (
	"fmt"
	"strings"
)

// ── Types ─────────────────────────────────────────────────────────────────────

// JoinType represents a SQL JOIN variant.
type JoinType string

const (
	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	RightJoin JoinType = "RIGHT JOIN"
	FullJoin  JoinType = "FULL OUTER JOIN"
)

// OrderDirection represents a sort direction.
type OrderDirection string

const (
	Asc  OrderDirection = "ASC"
	Desc OrderDirection = "DESC"
)

// ── Internal pieces ───────────────────────────────────────────────────────────

type joinClause struct {
	kind      JoinType
	table     string
	condition string
}

type orderClause struct {
	column    string
	direction OrderDirection
}

// argRegistry accumulates bound arguments across a builder's lifetime.
type argRegistry struct {
	args []any
}

// inject rewrites each '?' in expr to the next positional placeholder ($n),
// appends the corresponding values, and returns the rewritten expression.
// Panics if the number of '?' does not match len(values).
func (r *argRegistry) inject(expr string, values []any) string {
	count := strings.Count(expr, "?")
	if count != len(values) {
		panic(fmt.Sprintf(
			"querybuilder: condition %q has %d '?' but %d value(s) were provided",
			expr, count, len(values),
		))
	}
	var b strings.Builder
	vi := 0
	for i := 0; i < len(expr); i++ {
		if expr[i] == '?' {
			r.args = append(r.args, values[vi])
			fmt.Fprintf(&b, "$%d", len(r.args))
			vi++
		} else {
			b.WriteByte(expr[i])
		}
	}
	return b.String()
}

func (r *argRegistry) all() []any { return r.args }

// ── SELECT builder ────────────────────────────────────────────────────────────

// SelectBuilder builds a SELECT statement.
// Obtain one via New(table).Select(…).
type SelectBuilder struct {
	reg     argRegistry
	table   string
	columns []string
	joins   []joinClause
	wheres  []string
	groups  []string
	having  string
	orders  []orderClause
	lim     *int
	off     *int
}

// Where adds an AND condition. Use '?' as the placeholder for each value:
//
//	sb.Where("age > ? AND active = ?", 18, true)
func (s *SelectBuilder) Where(expr string, values ...any) *SelectBuilder {
	s.wheres = append(s.wheres, s.reg.inject(expr, values))
	return s
}

// OrWhere ORs the new condition with the last WHERE clause.
func (s *SelectBuilder) OrWhere(expr string, values ...any) *SelectBuilder {
	rewritten := s.reg.inject(expr, values)
	if len(s.wheres) == 0 {
		s.wheres = append(s.wheres, rewritten)
		return s
	}
	last := s.wheres[len(s.wheres)-1]
	s.wheres[len(s.wheres)-1] = fmt.Sprintf("(%s OR %s)", last, rewritten)
	return s
}

// WhereIn adds a WHERE col IN ($n, …) clause, binding each value.
func (s *SelectBuilder) WhereIn(column string, values ...any) *SelectBuilder {
	placeholders := make([]string, len(values))
	for i, v := range values {
		s.reg.args = append(s.reg.args, v)
		placeholders[i] = fmt.Sprintf("$%d", len(s.reg.args))
	}
	s.wheres = append(s.wheres, fmt.Sprintf("%s IN (%s)", column, strings.Join(placeholders, ", ")))
	return s
}

// WhereNull adds a WHERE col IS NULL clause.
func (s *SelectBuilder) WhereNull(column string) *SelectBuilder {
	s.wheres = append(s.wheres, column+" IS NULL")
	return s
}

// WhereNotNull adds a WHERE col IS NOT NULL clause.
func (s *SelectBuilder) WhereNotNull(column string) *SelectBuilder {
	s.wheres = append(s.wheres, column+" IS NOT NULL")
	return s
}

// Join adds a JOIN clause.
func (s *SelectBuilder) Join(kind JoinType, table, condition string) *SelectBuilder {
	s.joins = append(s.joins, joinClause{kind, table, condition})
	return s
}

// GroupBy adds GROUP BY columns.
func (s *SelectBuilder) GroupBy(columns ...string) *SelectBuilder {
	s.groups = append(s.groups, columns...)
	return s
}

// Having sets a HAVING clause (requires GroupBy).
func (s *SelectBuilder) Having(condition string) *SelectBuilder {
	s.having = condition
	return s
}

// OrderBy adds an ORDER BY expression.
func (s *SelectBuilder) OrderBy(column string, dir OrderDirection) *SelectBuilder {
	s.orders = append(s.orders, orderClause{column, dir})
	return s
}

// Limit sets the row limit.
func (s *SelectBuilder) Limit(n int) *SelectBuilder {
	s.lim = &n
	return s
}

// Offset sets the row offset.
func (s *SelectBuilder) Offset(n int) *SelectBuilder {
	s.off = &n
	return s
}

// Build returns the final SQL and the ordered argument slice.
func (s *SelectBuilder) Build() (string, []any) {
	cols := "*"
	if len(s.columns) > 0 {
		cols = strings.Join(s.columns, ", ")
	}

	var b strings.Builder
	fmt.Fprintf(&b, "SELECT %s FROM %s", cols, s.table)

	for _, j := range s.joins {
		fmt.Fprintf(&b, " %s %s ON %s", j.kind, j.table, j.condition)
	}

	if len(s.wheres) > 0 {
		fmt.Fprintf(&b, " WHERE %s", strings.Join(s.wheres, " AND "))
	}

	if len(s.groups) > 0 {
		fmt.Fprintf(&b, " GROUP BY %s", strings.Join(s.groups, ", "))
	}

	if s.having != "" {
		fmt.Fprintf(&b, " HAVING %s", s.having)
	}

	if len(s.orders) > 0 {
		parts := make([]string, len(s.orders))
		for i, o := range s.orders {
			parts[i] = fmt.Sprintf("%s %s", o.column, o.direction)
		}
		fmt.Fprintf(&b, " ORDER BY %s", strings.Join(parts, ", "))
	}

	if s.lim != nil {
		fmt.Fprintf(&b, " LIMIT %d", *s.lim)
	}

	if s.off != nil {
		fmt.Fprintf(&b, " OFFSET %d", *s.off)
	}

	return b.String(), s.reg.all()
}

// ── INSERT builder ────────────────────────────────────────────────────────────

// InsertBuilder builds an INSERT statement.
// Obtain one via New(table).Insert(…).
// Call Values() to bind the row data and get a build-ready InsertValuesBuilder.
type InsertBuilder struct {
	table   string
	columns []string
}

// Values binds each value positionally and returns an InsertValuesBuilder
// ready to call Build(). The number of values should match the columns
// provided to Insert().
func (i *InsertBuilder) Values(values ...any) *InsertValuesBuilder {
	var binder argBinder
	placeholders := make([]string, len(values))
	for idx, v := range values {
		placeholders[idx] = binder.bind(v)
	}
	return &InsertValuesBuilder{
		binder:       binder,
		table:        i.table,
		columns:      i.columns,
		placeholders: placeholders,
	}
}

// InsertValuesBuilder is the terminal stage of an INSERT chain.
// It only exposes Build() — no further mutations are possible.
type InsertValuesBuilder struct {
	binder       argBinder
	table        string
	columns      []string
	placeholders []string
}

// Build returns the final INSERT SQL and its arguments.
func (i *InsertValuesBuilder) Build() (string, []any) {
	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		i.table,
		strings.Join(i.columns, ", "),
		strings.Join(i.placeholders, ", "),
	)
	return sql, i.binder.all()
}

// ── UPDATE builder ────────────────────────────────────────────────────────────

// UpdateBuilder builds an UPDATE … SET … statement.
// Obtain one via New(table).Update().
// Add SET clauses with Set(), then optionally narrow with Where().
type UpdateBuilder struct {
	reg        argRegistry
	table      string
	setClauses []string
}

// Set adds a SET col = $n clause, injecting value immediately.
func (u *UpdateBuilder) Set(column string, value any) *UpdateBuilder {
	u.reg.args = append(u.reg.args, value)
	placeholder := fmt.Sprintf("$%d", len(u.reg.args))
	u.setClauses = append(u.setClauses, fmt.Sprintf("%s = %s", column, placeholder))
	return u
}

// Where transitions to an UpdateWhereBuilder so WHERE conditions can be added.
// Use '?' placeholders for values:
//
//	ub.Set("name", "Bob").Where("id = ?", 42)
func (u *UpdateBuilder) Where(expr string, values ...any) *UpdateWhereBuilder {
	w := &UpdateWhereBuilder{
		reg:        u.reg,
		table:      u.table,
		setClauses: u.setClauses,
	}
	w.wheres = append(w.wheres, w.reg.inject(expr, values))
	return w
}

// Build returns an UPDATE with no WHERE clause (updates all rows).
func (u *UpdateBuilder) Build() (string, []any) {
	sql := fmt.Sprintf("UPDATE %s SET %s", u.table, strings.Join(u.setClauses, ", "))
	return sql, u.reg.all()
}

// UpdateWhereBuilder is returned after Where() is called on an UpdateBuilder.
// It carries the same arg registry so placeholder indices stay consistent.
type UpdateWhereBuilder struct {
	reg        argRegistry
	table      string
	setClauses []string
	wheres     []string
}

// Where adds another AND condition. Use '?' placeholders for values:
//
//	uw.Where("deleted_at IS NULL").Where("role = ?", "admin")
func (u *UpdateWhereBuilder) Where(expr string, values ...any) *UpdateWhereBuilder {
	u.wheres = append(u.wheres, u.reg.inject(expr, values))
	return u
}

// Build returns the final UPDATE … SET … WHERE … SQL and its arguments.
func (u *UpdateWhereBuilder) Build() (string, []any) {
	sql := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		u.table,
		strings.Join(u.setClauses, ", "),
		strings.Join(u.wheres, " AND "),
	)
	return sql, u.reg.all()
}

// ── DELETE builder ────────────────────────────────────────────────────────────

// DeleteBuilder builds a DELETE statement.
// Obtain one via New(table).Delete().
type DeleteBuilder struct {
	reg    argRegistry
	table  string
	wheres []string
}

// Where adds an AND condition. Use '?' placeholders for values:
//
//	db.Where("id = ?", 7)
func (d *DeleteBuilder) Where(expr string, values ...any) *DeleteBuilder {
	d.wheres = append(d.wheres, d.reg.inject(expr, values))
	return d
}

// Build returns the final DELETE SQL and its arguments.
func (d *DeleteBuilder) Build() (string, []any) {
	var b strings.Builder
	fmt.Fprintf(&b, "DELETE FROM %s", d.table)
	if len(d.wheres) > 0 {
		fmt.Fprintf(&b, " WHERE %s", strings.Join(d.wheres, " AND "))
	}
	return b.String(), d.reg.all()
}

// ── Entry point ───────────────────────────────────────────────────────────────

// Table is the single entry point for all query builders.
type Table struct{ name string }

// New creates a Table entry point for the given table name.
func New(table string) *Table { return &Table{name: table} }

// Select starts a SELECT builder for the given columns (empty = "*").
func (t *Table) Select(columns ...string) *SelectBuilder {
	return &SelectBuilder{table: t.name, columns: columns}
}

// Insert starts an INSERT builder with the specified column names.
// Must be followed by Values(…) before calling Build().
func (t *Table) Insert(columns ...string) *InsertBuilder {
	return &InsertBuilder{table: t.name, columns: columns}
}

// Update starts an UPDATE builder. Call Set(…) to add SET clauses.
func (t *Table) Update() *UpdateBuilder {
	return &UpdateBuilder{table: t.name}
}

// Delete starts a DELETE builder.
func (t *Table) Delete() *DeleteBuilder {
	return &DeleteBuilder{table: t.name}
}

type argBinder struct {
	args []any
}

func (b *argBinder) bind(value any) string {
	b.args = append(b.args, value)
	return fmt.Sprintf("$%d", len(b.args))
}

func (b *argBinder) all() []any { return b.args }
