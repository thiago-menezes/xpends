package models

import (
	"xpends/webapi/db"
)

func Update(id int64, spend Spend) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(
		`UPDATE spends SET name=$2, category=$3, price=$4, done=$5, dueDate=$6 WHERE id=$1`,
		id,
		spend.Name,
		spend.Category,
		spend.Price,
		spend.Done,
		spend.DueDate,
	)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
