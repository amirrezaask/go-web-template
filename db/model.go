/*

type User struct {
	Id int
	Friends Users
}
func (u *User) Scan(row Row) error {
	return row.Scan(&u.Id)
}
func (u *User) LoadFriends(rows Rows) error {
	return u.Friends.Scan(rows)
}

type Users []*User
func (us Users) Scan(row Rows) error {
	for row.Next() {
		u := &User{}
		err := row.Scan(&u.Id)
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
)

//Row is the composition of *sql.Row and *sql.Rows
type Row interface {
	Scan(dest ...interface{}) error
}
type Rows interface {
	Row
	Next() bool
}

//RowScanner should be implemented by all of the types who can get their data from Row.
type RowScanner interface {
	Scan(Row) error
}

type RowsScanner interface {
	Scan(Rows) error
}

//Query a single row, use when you are sure that you're query selects only one row of data.
func QueryRow(db *sql.DB, rs RowScanner, query string, params ...interface{}) error {
	row := db.QueryRow(query, params...)
	return rs.Scan(row)
}

//Query multiple rows.
func Query(db *sql.DB, rs RowsScanner, query string, params ...interface{}) error {
	rows, err := db.Query(query, params...)
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
