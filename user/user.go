package user

import (
	"database/sql"

	"github.com/belajar/model"
)

//Dbase untuk menjadi pointer connection
type Dbase struct {
	Db *sql.DB
}

//QueryUser digunakan untuk mencari user dengan parameter nama
func (d *Dbase) QueryUser(uname string) []model.UserModel {
	db := d.Db
	defer db.Close()
	usr := model.UserModel{}
	res := []model.UserModel{}

	rows, err := d.Db.Query("SELECT Id,UserName,FirstName,LastName,Password FROM user WHERE UserName =?", uname)
	for rows.Next() {
		var id int
		var username, firstname, lastname, password string

		err = rows.Scan(&id, &username, &firstname, &lastname, &password)
		if err != nil {
			panic(err)
		}

		usr.ID = id
		usr.UserName = username
		usr.FirstName = firstname
		usr.LastName = lastname
		usr.Password = password

		res = append(res, usr)
		//balikan 1 atau banyak

	}
	/*db.QueryRow("SELECT Id,UserName,FirstName,LastName,Password FROM user WHERE UserName =?", uname).Scan(
		&usr.ID, &usr.UserName, &usr.FirstName, &usr.LastName, &usr.Password,
	)*/

	return res
}
