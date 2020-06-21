.PHONY: build

build:
	go build -o build/quoteblock.exe -v ./

run: build
	build/quoteblock.exe
	
.DUFAULT_GOAL := build

pack:
	d:\Apps\upx.exe --ultra-brute build\quoteblock*

deploy_win: deploy_for_win pack

deploy_for_win: 
	go build --ldflags "-w -s" -o build/quoteblock.exe -v .

deploy_nix: deploy_for_nix pack_nix

pack_nix:
	/cygdrive/d/apps/upx.exe --ultra-brute build/quoteblock

deploy_for_nix:
	powershell '$$env:GOOS = "linux"';	'go build --ldflags "-w -s" -o build/quoteblock -v .'

send_to_remote_pc:
	sftp -b sftp.cfg root@192.168.65.155

deploy_all: deploy_for_win deploy_for_nix pack

vendor:
	go mod tidy
	go mod download
	go mod vendor