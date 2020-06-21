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
	Если ip адресу или логину разрешен доступ в интернет то возвращает OK, иначе ERR message="access denied user not active" 
	-------------------------------------------------
	IMPORTANT: When you configure this script, you need to configure squid.conf.

	This lines you need to add to conf:

	If your authorization by login:

	#qouta acl section
	external_acl_type e_block ttl=10 negative_ttl=10 %LOGIN /path/to/bin/go_quotablock [-typedb mysql] -u login -p pass -h host_of_db -n name_of_db [-debug 4] [-log /var/log/squid/quoteblock.log] [-ttl 300]
	acl a_block external e_block

	If your authorization by IP address:

	#qouta acl section
	external_acl_type e_block ttl=10 negative_ttl=10 %SRC /path/to/bin/go_quotablock -typedb postgres -u login -p pass -h host_of_db -n name_of_db [-debug 4] [-log /var/log/squid/quoteblock.log] [-ttl 300]
	acl a_block external e_block

	Input line from squid:
		ip
		login
	
	Output line send back to squid:
		OK
		or ERR message="xxx"
	-------------------------------------------------	
`

type configType struct {
	fileLog  string
	fileConf string
	userDB   string
	passDB   string
	hostDB   string
	nameDB   string
	typeid   string
	typedb   string
	logLevel int
	ttl      int
}

var (
	config  configType
	user    userType
	rFSquid requestFromSquid
)

func init() {

	flag.StringVar(&config.typedb, "typedb", "mysql", `Type of DB: 
		'mysql' - MySQL, 
		'postgres' - PostgreSQL`)
	flag.StringVar(&config.fileLog, "log", "/var/log/squid/access.log", "File to log ")
	flag.StringVar(&config.userDB, "u", "", "User of DB")
	flag.StringVar(&config.passDB, "p", "", "Password of DB")
	flag.StringVar(&config.hostDB, "h", "localhost", "host of DB")
	flag.StringVar(&config.nameDB, "n", "", "name of DB")
	flag.IntVar(&config.logLevel, "debug", 0, `Level log: 
		0 - silent, 
		1 - error, start and end, 
		2 - error, start and end, warning, 
		3 - error, start and end, warning, access granted and denided,
		4 - error, start and end, warning, access granted and denided, request from squid `)
	flag.IntVar(&config.ttl, "ttl", 300, "Defines the time after which data from the database will be updated in seconds")
	flag.Parse()
	if config.typedb != "mysql" || config.typedb != "postgres" {
		chkM("Error. typedb must be 'mysql' or 'postgres'.", nil)
	}
	if config.userDB == "" {
		chkM("Error. Username must be specified.", nil)
	}
	if config.logLevel > 4 {
		config.logLevel = 4
	}
	if config.logLevel != 0 {
		log.SetFlags(log.Ldate | log.Ltime)
		toLog(config.logLevel, 1, "quoteblock | Init started")
		fl, err := os.OpenFile(config.fileLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		chkM(fmt.Sprintf("Error opening to write log file (%v): ", config.fileLog), err)
		defer fl.Close()
		log.SetOutput(fl)
	} // If logLevel not specified - silent mode
}

func main() {

	// dsn := "user:password@(host_bd)/dbname"
	// db, err := sql.Open("mysql", dsn)
	databaseURL := fmt.Sprintf("%v:%v@(%v)/%v", config.userDB, config.passDB, config.hostDB, config.nameDB)
	db, err := newDB(config.typedb, databaseURL)
	chk(err)
	defer db.Close()

	store := newStore(db)

	go store.getFromDB(&config)

	for {
		fmt.Scan(&rFSquid.user)
		s := store.checkUser(rFSquid)
		fmt.Println(s)
		toLog(config.logLevel, 4, "quoteblock | Squid requested:", rFSquid, ". Helper respond:", s)
	}
}
