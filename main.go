package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"goji.io"
	"goji.io/pat"

	//"golang.org/x/crypto/bcrypt"

	"github.com/belajar/user"
	"github.com/belajar/model"
	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//DbConn untuk koneksi, nantinya mungkin akan dipisah filenya
func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "myblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Println("error db")
		panic(err.Error())
	}
	return db
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Println("welcome home")
}

func test(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	db := user.Dbase{
		Db : DbConn(),
	}

	name := pat.Param(r,"name")
	users := db.QueryUser(name)
	var response model.ResponseUser
	
	if len(users) > 0 {
		response.Status = 1
		response.Message = "Success"
		response.Data = users
		
		res, err := json.Marshal(response)

		if err != nil {
			panic(err.Error())
		}

		w.Write(res)
		return
	}
	
	return
}

func routers(){
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), home)
	mux.HandleFunc(pat.Get("/test/:name"), test)

	http.ListenAndServe(":8081", mux)
}

func main() {
	routers()
}