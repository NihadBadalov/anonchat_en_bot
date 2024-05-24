package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	Id             int64
	Uid            string
	Gender         int8
	Age            int8
	Gatekeep_media bool
	Join_date      string
}

// Run a query on the database
//
// Example:
//
//	var version string
//
//	for rows.Next() {
//	  err := rows.Scan(&version)
//	  if err != nil {
//	    log.Fatal(err)
//	  }
//	}
func Query(query string) (*sql.Rows, error) {
	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	return rows, err
}

func AddUser(uid int64, gender int8, age int8, gatekeepMedia bool) error {
	rows, err := Query(
		fmt.Sprintf("INSERT INTO users (uid, gender, age, gatekeep_media) VALUES ('%d', %d, %d, %t) ON CONFLICT DO NOTHING", uid, gender, age, gatekeepMedia),
	)
	defer rows.Close()
	return err
}

func GetUserByUid(uid int64) (User, error) {
	rows, err := Query(
		fmt.Sprintf("SELECT * FROM users WHERE users.uid = '%d'", uid),
	)
	defer rows.Close()
	if err != nil {
		return User{}, err
	}

	var (
		_id             int64
		_uid            string
		_gender         int8
		_age            int8
		_gatekeep_media bool
		_join_date      string
	)
	for rows.Next() {
		err := rows.Scan(&_id, &_uid, &_gender, &_age, &_gatekeep_media, &_join_date)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return User{_id, _uid, _gender, _age, _gatekeep_media, _join_date}, err
}

func SetUserGender(uid int64, gender int) error {
  rows, err := Query(
    fmt.Sprintf("UPDATE users SET gender = %d WHERE users.uid = '%d'", gender, uid),
  )
  defer rows.Close()
  return err
}

func SetUserAge(uid int64, age int) error {
  rows, err := Query(
    fmt.Sprintf("UPDATE users SET age = %d WHERE users.uid = '%d'", age, uid),
  )
  defer rows.Close()
  return err
}

func SetUserGatekeepMedia(uid int64, gatekeepMedia bool) error {
  rows, err := Query(
    fmt.Sprintf("UPDATE users SET gatekeep_media = %t WHERE users.uid = '%d'", gatekeepMedia, uid),
  )
  defer rows.Close()
  return err
}
