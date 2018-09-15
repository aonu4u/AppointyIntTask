package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ars25198"
	dbname   = "postgres"
)

type list struct {
	ID    int    `json:"id", db:"id"`
	Items []int  `json:"items", db:"items"`
	Name  string `json:"name", db:"name"`
}
type Lists struct {
	Lists []list `json:"lists",db:"lists"`
}
type items struct {
	ID        int    `json:"id", db:"id"`
	Value     string `json:"value", db:"value"`
	Completed string `json:"completed", db:"completed"`
}
type Items struct {
	Items []items `json:"items", db:"items"`
}

var db *sql.DB

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//http.HandleFunc("/todolist", ToDoList)
	http.HandleFunc("/todolist/", Dellist)
	//http.HandleFunc("/todolist/",patchup)*/
	//http.HandleFunc("/todolist:addItem",add)
	http.HandleFunc("/todolist:deleteItem/", Delitem)
	http.HandleFunc("/todolist:getItem/", getItm)
	//http.HandleFunc("/todolist:updateItem",UpItem)
	log.Fatal(http.ListenAndServe(":9090", nil))

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func getItm(w http.ResponseWriter, r *http.Request) {
	i := Items{}
	id := r.URL.Path[len("/todolist:getItem/"):]
	e1 := querytodolist(&i, id)
	if e1 != nil {
		http.Error(w, e1.Error(), 500)
		return
	}
	fmt.Println(i.Items)
	z := i.Items
	json.NewEncoder(w).Encode(z)
}
func querytodolist(i *Items, id string) error {

	rows, err := db.Query(`SELECT *
	FROM items where id=$1`, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		ll := items{}
		err = rows.Scan(
			&ll.ID,
			&ll.Value,
			&ll.Completed,
		)
		if err != nil {
			return err
		}
		i.Items = append(i.Items, ll)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
func Delitem(w http.ResponseWriter, r *http.Request) {
	i := Items{}
	id := r.URL.Path[len("/todolist:deleteItem/"):]
	e1 := querytodo1(&i, id)
	if e1 != nil {
		http.Error(w, e1.Error(), 500)
		return
	}
	fmt.Println(i.Items)
	z := i.Items
	json.NewEncoder(w).Encode(z)
}
func querytodo1(i *Items, id string) error {
	rows, err := db.Query(`DELETE
	FROM items WHERE id=$1`, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		ll := items{}
		err = rows.Scan(
			&ll.ID,
			&ll.Value,
			&ll.Completed,
		)
		if err != nil {
			return err
		}
		i.Items = append(i.Items, ll)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
func Dellist(w http.ResponseWriter, r *http.Request) {
	l := Lists{}
	id := r.URL.Path[len("/todolist/"):]
	e1 := querytodo2(&l, id)
	if e1 != nil {
		http.Error(w, e1.Error(), 500)
		return
	}
	fmt.Println(l.Lists)
	z := l.Lists
	json.NewEncoder(w).Encode(z)
}
func querytodo2(l *Lists, id string) error {
	rows, err := db.Query(`DELETE
	FROM list WHERE id=$1`, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		ll := list{}
		err = rows.Scan(
			&ll.ID,
			&ll.Items,
			&ll.Name,
		)
		if err != nil {
			return err
		}
		l.Lists = append(l.Lists, ll)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
