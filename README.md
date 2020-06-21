# quoteblock
Quota Squid Helper

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




# How To Build



If you don't have golang, you must first do

	wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz

or any other distribution kit from [https://golang.org/dl/](https://golang.org/dl/) suitable for your operation system

	tar -xf go1.12.7.linux-amd64.tar.gz

and next

	git clone https://github.com/Rid-lin/quoteblock.git
	cd quoteblock
	go1.12.7.linux-amd64/bin/go build 

or the other patch where golang is

	make
	make install

where *squid* is the name of the group under which squid runs

# Thanks

- Идея реализации была взята отсюда https://github.com/funway/squid-helper
- Так же пригодился этот материал http://freesoftwaremagazine.com/articles/authentication_with_squid/
- Без перевода директив для Squid мне понадобилось бы больше времени  http://break-people.ru/cmsmade/index.php?page=translate_squid_reference_tag_external_acl_type
- Очень простое описание формата логов https://wiki.enchtex.info/doc/squidlogformat
