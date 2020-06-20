package main

import "database/sql"

type storeType struct {
	db    *sql.DB
	users map[string]userType
}

type userType struct {
	alias  string
	ip     string
	login  string
	active bool
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

	// return "OK"
	var index string

	if rFSquid.login != "" {
		index = rFSquid.login
	} else {
		index = rFSquid.ip
	}
	if _, ok := s.users[index]; ok {
		if s.users[index].active {
			toLog(config.logLevel, 5, "quoteblock |", index, "- granted access, IP is active")
			return "OK"
		}
	}
	toLog(config.logLevel, 5, "quoteblock |", index, "- access denied, IP-addres not active")
	return `ERR message="access denied, IP-addres not active"`
}
