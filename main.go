package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// DESCRIPTION - описание использования
var DESCRIPTION = `
	Access Control (ACL) helper for Squid and [Screen Squid](https://sourceforge.net/projects/screen-squid/) 
	Программа проверяет ip адрес или логин пользователя на вхождение в список модуля Quotas программы Screen Squid.
	Если ip адресу или логину разрешен доступ в интернет то возвращает OK, иначе ERR message="access denied IP-addres not active" 
	-------------------------------------------------
	Additional parameter is "typeid". 
	If your authorization by login, you need to set it in 0 (default).
	If your authorization by IP, you need to set it in 1.

	IMPORTANT: When you configure this script, you need to configure squid.conf.

	This lines you need to add to conf:

	If your authorization by login:

	#acl section
	external_acl_type e_block ttl=10 negative_ttl=10 %LOGIN %SRC /path/to/bin/go_quotablock [-typeid 0] [-typedb 0] -u login -p pass -h host_of_db [-debug 4] [-log /var/log/squid/quoteblock.log]
	acl a_block external e_block

	If your authorization by IP address:

	#acl section
	external_acl_type e_block ttl=10 negative_ttl=10 %SRC /path/to/bin/go_quotablock -typeid 1 -typedb 1 -u login -p pass -h host_of_db [-debug 4] [-log /var/log/squid/quoteblock.log]
	acl a_block external e_block

	For both authorization

	#http rules section
	http_access allow a_block

	Input line from squid:
		ip
	
	Output line send back to squid:
		OK
		or ERR message="xxx"
		or BH message="xxx"
	-------------------------------------------------	
`

var (
	// conf     configType
	config  configType
	user    userType
	rFSquid requestFromSquid
)

func init() {

	// flag.StringVar(&config.typeid, "typeid", "0", "Type of identifity: 0 -by login, 1 - by IP")
	flag.StringVar(&config.typedb, "typedb", "mysql", "Type of DB: 'mysql' - MySQL, 'postgres' - PostgreSQL")
	flag.StringVar(&config.fileLog, "log", "", "File to log ")
	flag.StringVar(&config.userDB, "u", "", "User of DB")
	flag.StringVar(&config.passDB, "p", "", "Password of DB")
	flag.StringVar(&config.hostDB, "h", "localhost", "host of DB")
	flag.StringVar(&config.nameDB, "n", "", "name of DB")
	flag.IntVar(&config.logLevel, "debug", 0, "Level log: 0 - silent, 1 - error, start and end, 2 - '1' + warning, 3 - '2' + read config, parse lines, request from squid 4 - '3' + access granted and denided, 5 - very many logs")
	flag.IntVar(&config.ttl, "ttl", 300, "Defines the time after which data from the database will be updated")
	flag.Parse()
	if config.typedb != "mysql" || config.typedb != "postgres" {
		chkM("Error. typedb must be 'mysql' or 'postgres'.", nil)
	}
	if config.userDB == "" {
		chkM("Error. Username must be specified.", nil)
	}
	if config.fileLog != "" {
		log.SetFlags(log.Ldate | log.Ltime)
		toLog(config.logLevel, 1, "quoteblock | Init started")
		fl, err := os.OpenFile(config.fileLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		chkM(fmt.Sprintf("Error opening to write log file (%v): ", config.fileLog), err)
		defer fl.Close()
		log.SetOutput(fl)
	} // If fileLog not used - silent mode
	// conf.New(typeid, typedb, userDB, passDB, hostDB, fileLog, fileConf, logLevel, ttl)
}

func main() {

	// 	dsn := "user:password@/dbname"
	// db, err := sql.Open("mysql", dsn)
	databaseURL := fmt.Sprintf("%v:%v@%v/%v", config.userDB, config.passDB, config.hostDB, config.nameDB)
	db, err := newDB(config.typedb, databaseURL)
	chk(err)
	defer db.Close()

	store := newStore(db)

	go store.getFromDB(&config)

	for {
		fmt.Scan(&rFSquid.login, &rFSquid.ip)
		s := store.checkUser(rFSquid)
		fmt.Println(s)
		toLog(config.logLevel, 3, "quoteblock | Squid requested:", rFSquid, ". Helper respond:", s)
	}
}
