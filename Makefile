all: clean build pack

clean:
	$(RM) alfred_ddb

build:
	go build -ldflags="-s -w" alfred_ddb.go

pack:
	upx --brute alfred_ddb
