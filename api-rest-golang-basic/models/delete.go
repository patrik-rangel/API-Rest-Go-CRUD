package models

import "github.com/patrik-rangel/API-Rest-Go_CRUD/tree/main/api-rest-golang-basic/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=?`, id)
	if err != nil{
		return 0, err
	}

	return res.RowsAffected()
}
