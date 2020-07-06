/*
type User struct {
	Id int
	Friends Users
}
func (u *User) Scan(row *sql.Row) error {
	return row.Scan(&u.Id)
}
func (u *User) LoadFriends(rows *sql.Rows) error {
	return u.Friends.Scan(rows)
}

type Users []*User
func (us Users) Scan(rows *sql.Rows) error {
	for rows.Next() {
		u := &User{}
		err := rows.Scan(&u.Id)
		if err != nil {
			return err
		}
		us = append(us, u)
	}
	return nil
}

func main() {
	var db *sql.DB
	users := Users{}
	err := Query(db, users, `SELECT * FROM users`)
	if err != nil {
		panic(err)
	}
	user := &User{}
	err = QueryRow(db, user, `SELECT * FROM users WHERE id=?`, 1)
	friends := Users{}
	err =  Query(db, friends, `SELECT * FROM users WHERE friend_id=?`, 1)
}
*/

package db

import (
	"database/sql"
	"fmt"
	"reflect"
)

//RowScanner should be implemented by all of the types who can get their data from Row.
type RowScanner interface {
	Scan(*sql.Row) error
}

type RowsScanner interface {
	Scan(*sql.Rows) error
}

//Query a single row, use when you are sure that you're query selects only one row of data.
func QueryRow(db *sql.DB, rs RowScanner, query string, args ...interface{}) error {
	row := db.QueryRow(query, args...)
	return rs.Scan(row)
}

//Query multiple rows.
func Query(db *sql.DB, rs RowsScanner, query string, args ...interface{}) error {
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}
	for rows.Next() {
		err = rs.Scan(rows)
		if err != nil {
			return err
		}
	}
	return nil
}
//Fields generate a list of pointers to the fields of given struct, this function use reflection so it should have
//some overhead.
func Fields(m interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(m)
	if v.Type().Kind() != reflect.Struct {
		if v.Type().Kind() == reflect.Ptr {
			return Fields(v.Elem().Interface())
		}
		return nil, fmt.Errorf("cant get fields of %s", v.Type().String())
	}
	var ptrs []interface{}
	for i:=0;i<v.NumField();i++ {
		f := v.Field(i).Interface()
		ptrs = append(ptrs, &f)
	}
	return ptrs, nil
}