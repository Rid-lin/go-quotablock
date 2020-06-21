package main

import (
	"database/sql"
	"sync"
)

type storeType struct {
	db *sql.DB
	sync.Mutex
	users map[string]userType
}

type userType struct {
	alias  string
	ip     string
	login  string
	active string
}

type requestFromSquid struct {
	ip    string
	login string
}

func (s *storeType) checkUser(rFSquid requestFromSquid) string {
	// запросить список пользователей из бд
	// проверить текущего пользователя на вхождение в список
	// если он есть, то ОК
	// если нет то ЕРР
	s.Mutex.Lock()
	if _, ok := s.users[rFSquid.ip]; ok {
		if s.users[rFSquid.ip].active == "0" {
			toLog(config.logLevel, 5, "quoteblock |", rFSquid.ip, "- granted access, IP is active")
			return "OK"
		}
	}
	if _, ok := s.users[rFSquid.login]; ok {
		if s.users[rFSquid.login].active == "0" {
			toLog(config.logLevel, 5, "quoteblock |", rFSquid.login, "- granted access, IP is active")
			return "OK"
		}
	}
	s.Mutex.Unlock()
	toLog(config.logLevel, 5, "quoteblock |", rFSquid, "- access denied, IP-addres not active")
	return `ERR message="access denied, user not active"`
}
