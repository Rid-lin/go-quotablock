# quoteblock
Quota Squid Helper

Access Control (ACL) helper for squid 3.4+

Программа проверяет ip адрес пользователя на вхождение в список не превысивших лимит трафика.
Если ip адрес присутсвует, то возвращает OK

	-------------------------------------------------
	Usage in squid.conf:
	external_acl_type quota_ip cache=0 children-max=1 ipv4 %SRC /usr/local/bin/quoteblock -l 4 -c /etc/quoteblock/quoteblock.json -f /var/log/squid/quoteblock.log
	acl allow_to_inet external quota_ip
	http_access allow allow_to_inet
	Input line from squid:
		ip
	Output line send back to squid:
		OK
		or ERR message="xxx"
		or BH message="xxx"
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
