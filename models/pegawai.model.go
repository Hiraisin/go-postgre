package models

import (
	"net/http"

	"github.com/hiraisin/go-postgre/config"
)

func GetAll() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := config.CreateCon()

	sqlStatement := "select * from pegawai"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)

		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StorePegawai(nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := config.CreateCon()

	sqlStatement := "insert pegawai (nama,alamat,telepon) values (?,?,?) "

	stmt, err := con.Prepare(sqlStatement)

	result, err := stmt.Exec(nama, alamat, telepon)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}
