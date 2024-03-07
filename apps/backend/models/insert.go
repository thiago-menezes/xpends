package models

import (
	"xpends/webapi/db"
)

func Insert(spend Spend) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO spends (name, category, price, done, dueDate) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = conn.QueryRow(
		sql,
		spend.Name,
		spend.Category,
		spend.Price,
		spend.Done,
		spend.DueDate,
	).Scan(&id)

	return
}
