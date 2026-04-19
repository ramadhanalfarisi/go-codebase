package query_builder

import (
	"testing"

)

func check(t *testing.T, label, wantSQL, gotSQL string, wantArgs, gotArgs []any) {
	t.Helper()
	if wantSQL != gotSQL {
		t.Errorf("[%s] SQL\n  want: %s\n   got: %s", label, wantSQL, gotSQL)
	}
	if len(wantArgs) != len(gotArgs) {
		t.Errorf("[%s] arg count: want %d got %d", label, len(wantArgs), len(gotArgs))
		return
	}
	for i := range wantArgs {
		if wantArgs[i] != gotArgs[i] {
			t.Errorf("[%s] args[%d]: want %v got %v", label, i, wantArgs[i], gotArgs[i])
		}
	}
}

// ── SELECT ────────────────────────────────────────────────────────────────────

func TestSelectAll(t *testing.T) {
	sql, args := New("users").Select().Build()
	check(t, "select *", "SELECT * FROM users", sql, nil, args)
}

func TestSelectColumns(t *testing.T) {
	sql, args := New("users").Select("id", "name", "email").Build()
	check(t, "select cols", "SELECT id, name, email FROM users", sql, nil, args)
}

func TestSelectWhere(t *testing.T) {
	sql, args := New("users").Select("id", "name").
		Where("age > ?", 18).
		Where("active = ?", true).
		Build()

	check(t, "where AND",
		"SELECT id, name FROM users WHERE age > $1 AND active = $2",
		sql, []any{18, true}, args)
}

func TestSelectWhereMultiArg(t *testing.T) {
	sql, args := New("users").Select().
		Where("age > ? AND active = ?", 18, true).
		Build()

	check(t, "where multi-arg",
		"SELECT * FROM users WHERE age > $1 AND active = $2",
		sql, []any{18, true}, args)
}

func TestSelectOrWhere(t *testing.T) {
	sql, args := New("users").Select().
		Where("role = ?", "admin").
		OrWhere("role = ?", "moderator").
		Build()

	check(t, "or where",
		"SELECT * FROM users WHERE (role = $1 OR role = $2)",
		sql, []any{"admin", "moderator"}, args)
}

func TestSelectWhereIn(t *testing.T) {
	sql, args := New("users").Select().WhereIn("id", 1, 2, 3).Build()
	check(t, "where in",
		"SELECT * FROM users WHERE id IN ($1, $2, $3)",
		sql, []any{1, 2, 3}, args)
}

func TestSelectWhereNull(t *testing.T) {
	sql, args := New("users").Select().WhereNull("deleted_at").Build()
	check(t, "where null", "SELECT * FROM users WHERE deleted_at IS NULL", sql, nil, args)
}

func TestSelectWhereNotNull(t *testing.T) {
	sql, args := New("users").Select().WhereNotNull("email").Build()
	check(t, "where not null", "SELECT * FROM users WHERE email IS NOT NULL", sql, nil, args)
}

func TestSelectJoin(t *testing.T) {
	sql, args := New("orders").
		Select("orders.id", "users.name").
		Join(InnerJoin, "users", "orders.user_id = users.id").
		Build()

	check(t, "inner join",
		"SELECT orders.id, users.name FROM orders INNER JOIN users ON orders.user_id = users.id",
		sql, nil, args)
}

func TestSelectGroupByHaving(t *testing.T) {
	sql, args := New("orders").
		Select("user_id", "COUNT(*) AS total").
		GroupBy("user_id").
		Having("COUNT(*) > 5").
		Build()

	check(t, "group by having",
		"SELECT user_id, COUNT(*) AS total FROM orders GROUP BY user_id HAVING COUNT(*) > 5",
		sql, nil, args)
}

func TestSelectOrderLimitOffset(t *testing.T) {
	sql, args := New("products").
		Select("id", "name", "price").
		OrderBy("price", Desc).
		OrderBy("name", Asc).
		Limit(10).
		Offset(20).
		Build()

	check(t, "order limit offset",
		"SELECT id, name, price FROM products ORDER BY price DESC, name ASC LIMIT 10 OFFSET 20",
		sql, nil, args)
}

func TestSelectComplex(t *testing.T) {
	sql, args := New("orders o").
		Select("o.id", "u.name", "SUM(oi.price) AS total").
		Join(LeftJoin, "users u", "o.user_id = u.id").
		Join(InnerJoin, "order_items oi", "oi.order_id = o.id").
		Where("o.status = ?", "completed").
		WhereNotNull("o.shipped_at").
		GroupBy("o.id", "u.name").
		Having("SUM(oi.price) > 100").
		OrderBy("total", Desc).
		Limit(5).
		Build()

	want := "SELECT o.id, u.name, SUM(oi.price) AS total FROM orders o" +
		" LEFT JOIN users u ON o.user_id = u.id" +
		" INNER JOIN order_items oi ON oi.order_id = o.id" +
		" WHERE o.status = $1 AND o.shipped_at IS NOT NULL" +
		" GROUP BY o.id, u.name" +
		" HAVING SUM(oi.price) > 100" +
		" ORDER BY total DESC" +
		" LIMIT 5"

	check(t, "complex select", want, sql, []any{"completed"}, args)
}

func TestSelectMixedWhere(t *testing.T) {
	// WhereIn followed by a regular Where — indices must continue correctly.
	sql, args := New("products").Select("id", "name").
		WhereIn("category_id", 1, 2).
		Where("price > ?", 50.0).
		Build()

	check(t, "where in + where",
		"SELECT id, name FROM products WHERE category_id IN ($1, $2) AND price > $3",
		sql, []any{1, 2, 50.0}, args)
}

// ── INSERT ────────────────────────────────────────────────────────────────────

func TestInsert(t *testing.T) {
	sql, args := New("users").
		Insert("name", "email", "age").
		Values("Alice", "alice@example.com", 30).
		Build()

	check(t, "insert",
		"INSERT INTO users (name, email, age) VALUES ($1, $2, $3)",
		sql, []any{"Alice", "alice@example.com", 30}, args)
}

// ── UPDATE ────────────────────────────────────────────────────────────────────

func TestUpdateWithWhere(t *testing.T) {
	sql, args := New("users").Update().
		Set("name", "Bob").
		Set("active", false).
		Where("id = ?", 42).
		Build()

	check(t, "update where",
		"UPDATE users SET name = $1, active = $2 WHERE id = $3",
		sql, []any{"Bob", false, 42}, args)
}

func TestUpdateNoWhere(t *testing.T) {
	sql, args := New("sessions").Update().Set("active", false).Build()
	check(t, "update no where",
		"UPDATE sessions SET active = $1",
		sql, []any{false}, args)
}

func TestUpdateMultiWhere(t *testing.T) {
	sql, args := New("users").Update().
		Set("role", "admin").
		Where("id = ?", 1).
		Where("deleted_at IS NULL").
		Build()

	check(t, "update multi where",
		"UPDATE users SET role = $1 WHERE id = $2 AND deleted_at IS NULL",
		sql, []any{"admin", 1}, args)
}

// ── DELETE ────────────────────────────────────────────────────────────────────

func TestDelete(t *testing.T) {
	sql, args := New("users").Delete().Where("id = ?", 7).Build()
	check(t, "delete where",
		"DELETE FROM users WHERE id = $1",
		sql, []any{7}, args)
}

func TestDeleteAll(t *testing.T) {
	sql, args := New("sessions").Delete().Build()
	check(t, "delete all", "DELETE FROM sessions", sql, nil, args)
}

func TestDeleteMultiWhere(t *testing.T) {
	sql, args := New("tokens").Delete().
		Where("user_id = ?", 99).
		Where("expired = ?", true).
		Build()

	check(t, "delete multi where",
		"DELETE FROM tokens WHERE user_id = $1 AND expired = $2",
		sql, []any{99, true}, args)
}

// ── Panic on mismatch ─────────────────────────────────────────────────────────

func TestWhereMismatchPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for ? / value count mismatch, got none")
		}
	}()
	New("users").Select().Where("id = ? AND role = ?", 1 /* missing second value */)
}
