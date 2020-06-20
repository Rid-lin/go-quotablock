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
	Usage in squid.conf:
	external_acl_type quota_ip cache=30 children-max=10 ipv4 %SRC %LOGIN /usr/local/bin/quoteblock -debug 4 -log /var/log/squid/quoteblock.log
	acl allow_to_inet external quota_ip
	http_access allow allow_to_inet
	Input line from squid:
		ip
	Output line send back to squid:
		OK
		or ERR message="xxx"
		or BH message="xxx"
	-------------------------------------------------
`

var (
	conf     configType
	user     userType
	fileLog  string
	fileConf string
	userDB   string
	passDB   string
	hostDB   string
	logLevel int
	ttl      int
	rFSquid  requestFromSquid
)

func init() {

	flag.StringVar(&fileLog, "log", "", "File to log ")
	flag.StringVar(&userDB, "u", "", "User of DB")
	flag.StringVar(&passDB, "p", "", "Password of DB")
	flag.StringVar(&hostDB, "h", "localhost", "host of DB")
	flag.IntVar(&logLevel, "debug", 0, "Level log: 0 - silent, 1 - error, start and end, 2 - '1' + warning, 3 - '2' + read config, parse lines, request from squid 4 - '3' + access granted and denided, 5 - very many logs")
	flag.IntVar(&ttl, "ttl", 300, "Defines the time after which data from the database will be updated")
	flag.Parse()
	if userDB == "" {
		chkM("Error. Username must be specified.", nil)
	}
	if fileLog != "" {
		log.SetFlags(log.Ldate | log.Ltime)
		toLog(logLevel, 1, "quoteblock | Init started")
		fl, err := os.OpenFile(fileLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		chkM(fmt.Sprintf("Error opening to write log file (%v): ", fileLog), err)
		defer fl.Close()
		log.SetOutput(fl)
	} // If fileLog not used - silent mode
	conf.New(fileLog, fileConf, logLevel, ttl)
}

func main() {

	// go getFromDB(ttl)

	for {
		fmt.Scan(&rFSquid.IP)
		s := conf.CheckUser(rFSquid.IP)
		fmt.Println(s)
		toLog(logLevel, 3, "quoteblock | Squid requested:", rFSquid.IP, ". Helper respond:", s)
	}
}
