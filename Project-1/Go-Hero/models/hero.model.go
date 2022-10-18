package models

import (
	"net/http"
	"go-hero/db"
	validator "github.com/go-playground/validator/v10"
)

type Hero struct {
	Id      		int    		`json:"id"`
	Nama_Hero    	string 		`json:"nama_hero" validate:"required"`
	Role  			string 		`json:"role" validate:"required"`
	Emblem 			string 		`json:"emblem" validate:"required"`
	Battle_Spell 	string 		`json:"battle_spell" validate:"required"`
}

func GetAllHero() (Response, error) {

	var obj Hero
	var arrobj []Hero
	var res Response

	con := db.CreateCon()

	// Execute the query
	sqlStatement := "SELECT * FROM Hero"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama_Hero, &obj.Role, &obj.Emblem, &obj.Battle_Spell)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	//setup response
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func PostHero(nama_hero string , role string, emblem string , battle_spell string) (Response, error) {
	var res Response

	v := validator.New()

	peg := Hero{
		Nama_Hero		: nama_hero,
		Role			: role,
		Emblem			: emblem,
		Battle_Spell	: battle_spell,
	}

	err := v.Struct(peg)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT hero (nama_hero, role, emblem, battle_spell) VALUES (?, ?, ?,?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama_hero, role, emblem, battle_spell)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	//setup response
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateHero(Id int, nama_hero string, role string, emblem string , battle_spell string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE hero SET nama_hero = ?, role = ?, emblem = ?, battle_spell = ? WHERE id = ?"

	// Prepare statement
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	// Execute the SQL statement
	result, err := stmt.Exec(nama_hero, role, emblem, battle_spell, Id)
	if err != nil {
		return res, err
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	//setup response
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}

func DeleteHero(id int) (Response, error){
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM hero WHERE id = ?"

	// Prepare statement
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	// Execute the SQL statement
	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	//setup response
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}