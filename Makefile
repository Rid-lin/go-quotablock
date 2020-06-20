.PHONY :build

build:
	go build -o build/quoteblock.exe -v . 

.PHONY :test

# test:
# 	go test -race -timeout 30s ./...

# .PHONY :test_log

# test_log:
# 	go test -v -race -timeout 30s ./...

# .PHONY :run

# run: test build
run: build
	.\quoteblock.exe
	
.DUFAULT_GOAL := build

.PHONY :pack

pack:
	d:\Apps\upx.exe --ultra-brute build\quoteblock.exe

.PHONY :deploy_win

deploy_win: test 
	go build --ldflags "-w -s" -o build/quoteblock.exe -v ./cmd
	d:\Apps\upx.exe --ultra-brute build\quoteblock.exe

build_for_deploy:
	go build --ldflags "-w -s" -o build/quoteblock.exe -v ./cmd 

vendor:
	go mod tidy
	go mod download
	go mod vendor