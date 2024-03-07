package models

import "xpends/webapi/db"

func GetAll() (spends []Spend, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM spends`)
	if err != nil {
		return
	}

	for rows.Next() {
		var spend Spend

		err = rows.Scan(
			&spend.ID,
			&spend.Name,
			&spend.Category,
			&spend.Price,
			&spend.Done,
			&spend.DueDate,
		)
		if err != nil {
			continue
		}
		spends = append(spends, spend)
	}

	return
}
