package models

import "xpends/webapi/db"

func Get(id int64) (spend Spend, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM spends WHERE id=$1`, id)
	err = row.Scan(
		&spend.ID,
		&spend.Name,
		&spend.Category,
		&spend.Price,
		&spend.Done,
		&spend.DueDate,
	)
	return
}
