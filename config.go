package main

type configType struct {
	fileLog  string
	fileConf string
	logLevel int
	ttl      int
	Users    map[string]userType
}

type userType struct {
	IP     string
	Name   string
	Active bool
}

type requestFromSquid struct {
	IP   string
	Name string
}

func (conf *configType) New(fileLog, fileConf string, logLevel, ttl int) {
	conf.fileLog = fileLog
	conf.fileConf = fileConf
	conf.logLevel = logLevel
	conf.ttl = ttl
}

func (conf *configType) CheckUser(ip string) string {
	// запросить список пользователей из бд
	// проверить текущего пользователя на вхождение в список
	// если он есть, то ОК
	// если нет то ЕРР

	return "OK"

	// if conf.Users[ip].Active {
	// 	toLog(logLevel, 5, "quoteblock |", ip, "- granted access, IP is active")
	// 	return "OK"
	// }
	// toLog(logLevel, 5, "quoteblock |", ip, "- access denied, IP-addres not active")
	// return `ERR message="access denied, IP-addres not active"`
}
