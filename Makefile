.PHONY: build

build:
	go build -o build/quoteblock.exe -v ./

run: build
	build/quoteblock.exe
	
.DUFAULT_GOAL := build

pack:
	d:\Apps\upx.exe --ultra-brute build\quoteblock*

pack_nix:
	/cygdrive/d/apps/upx.exe --ultra-brute build\quoteblock


deploy_win: deploy_for_win pack

deploy_for_win: 
	go build --ldflags "-w -s" -o build/quoteblock.exe -v .

deploy_nix: deploy_for_nix pack_nix

deploy_for_nix:
	set GOOS=linux
	go build --ldflags "-w -s" -o build/quoteblock -v .

deploy_all: deploy_for_win deploy_nix

vendor:
	go mod tidy
	go mod download
	go mod vendor