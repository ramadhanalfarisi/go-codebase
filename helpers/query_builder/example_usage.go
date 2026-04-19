package query_builder

import (
	"fmt"
)

func main() {

	// ── SELECT ────────────────────────────────────────────────────────────────
	// Each '?' in the condition is rewritten to $1, $2, … automatically.
	sql, args := New("users").
		Select("id", "name", "email").
		Join(LeftJoin, "orders", "orders.user_id = users.id").
		Where("users.active = ?", true).
		WhereNotNull("users.email").
		GroupBy("users.id").
		Having("COUNT(orders.id) > 0").
		OrderBy("users.name", Asc).
		Limit(20).
		Offset(0).
		Build()

	fmt.Println("── SELECT ──")
	fmt.Println(sql)
	fmt.Println(args)
	// SELECT id, name, email FROM users
	//   LEFT JOIN orders ON orders.user_id = users.id
	//   WHERE users.active = $1 AND users.email IS NOT NULL
	//   GROUP BY users.id HAVING COUNT(orders.id) > 0
	//   ORDER BY users.name ASC LIMIT 20 OFFSET 0
	// [true]

	// Multiple '?' in a single Where call is also fine:
	sql, args = New("events").Select().
		Where("starts_at >= ? AND ends_at <= ?", "2024-01-01", "2024-12-31").
		Build()

	fmt.Println("\n── SELECT multi-arg where ──")
	fmt.Println(sql)
	fmt.Println(args)
	// SELECT * FROM events WHERE starts_at >= $1 AND ends_at <= $2
	// [2024-01-01 2024-12-31]

	// ── WHERE IN ─────────────────────────────────────────────────────────────
	sql, args = New("products").
		Select("id", "name", "price").
		WhereIn("category_id", 1, 2, 3).
		OrderBy("price", Desc).
		Build()

	fmt.Println("\n── WHERE IN ──")
	fmt.Println(sql)
	fmt.Println(args)
	// SELECT id, name, price FROM products WHERE category_id IN ($1, $2, $3) ORDER BY price DESC
	// [1 2 3]

	// ── OR WHERE ─────────────────────────────────────────────────────────────
	sql, args = New("users").Select("id", "name").
		Where("role = ?", "admin").
		OrWhere("role = ?", "moderator").
		Build()

	fmt.Println("\n── OR WHERE ──")
	fmt.Println(sql)
	fmt.Println(args)
	// SELECT id, name FROM users WHERE (role = $1 OR role = $2)
	// [admin moderator]

	// ── INSERT ────────────────────────────────────────────────────────────────
	sql, args = New("users").
		Insert("name", "email", "role").
		Values("Alice", "alice@example.com", "admin").
		Build()

	fmt.Println("\n── INSERT ──")
	fmt.Println(sql)
	fmt.Println(args)
	// INSERT INTO users (name, email, role) VALUES ($1, $2, $3)
	// [Alice alice@example.com admin]

	// ── UPDATE ───────────────────────────────────────────────────────────────
	// Set() injects values; Where() continues the same $n sequence.
	sql, args = New("users").Update().
		Set("name", "Bob").
		Set("active", true).
		Where("id = ?", 42).
		Build()

	fmt.Println("\n── UPDATE ──")
	fmt.Println(sql)
	fmt.Println(args)
	// UPDATE users SET name = $1, active = $2 WHERE id = $3
	// [Bob true 42]

	// ── DELETE ───────────────────────────────────────────────────────────────
	sql, args = New("sessions").Delete().
		Where("user_id = ?", 99).
		Build()

	fmt.Println("\n── DELETE ──")
	fmt.Println(sql)
	fmt.Println(args)
	// DELETE FROM sessions WHERE user_id = $1
	// [99]
}
