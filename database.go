package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // ....

	_ "github.com/lib/pq" // ...
)

func (s *storeType) getFromDB(config *configType) {

	rows, err := s.db.Query(`SELECT a.name, l.name,a.userlogin, q.status FROM scsq_alias a 
	JOIN scsq_ipaddress l ON a.tableid=l.id
	JOIN scsq_mod_quotas q ON q.aliasid = a.id;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := new(userType)
		err := rows.Scan(&user.alias, &user.ip, &user.login, &user.active)
		if err != nil {
			log.Fatal(err)
		}
		if user.login == "" {
			s.users[user.login] = *user
		}
		s.users[user.ip] = *user
	}
	chkM("", err)

}

//New ..
func newStore(db *sql.DB) *storeType {
	return &storeType{
		db: db,
	}
}

func newDB(typedb, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open(typedb, databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
