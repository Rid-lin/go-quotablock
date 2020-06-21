
# DIR - The name of the folder in which go-quotablock will be installed
DIR = /modules/Quotas
# PREFIX - The path where ScreenSquid will be installed
PREFIX = /var/www/html/screensquid

# The user and group that is required to start the http server
PROXY-USER = squid
PROXY-GROUP = squid

.PHONY: build

build:
	go build -o build/go-quotablock -v ./

run: build
	build/go-quotablock
	
.DUFAULT_GOAL := build

install: sayhello install-copy fix-permission sayOK

install-move:
	mv build/go-quotablock $(PREFIX)/$(DIR)/

fix-permission:
	chown PROXY-USER:PROXY-GROUP $(PREFIX)/$(DIR)/go-quotablock
	chmod +x $(PREFIX)/$(DIR)/go-quotablock

clear:
	rm build/go-quotablock

uninstall:
	rm $(PREFIX)/$(DIR)/go-quotablock

sayhello:
	@echo ''
	@echo 'Hello! Starting install go-quotablock for Screen Squid'
	@echo ''

sayOK:
	@echo ''
	@echo 'go-quotablock for Screen Squid installed successfully!'
	@echo 'Please, go to browser and type http://yourserverip/screensquid'
