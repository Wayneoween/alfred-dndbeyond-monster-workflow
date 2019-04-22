all: clean build copy

release: clean build pack copy

clean:
	$(RM) alfred-dndbeyond-monster-workflow

build:
	go build -ldflags="-s -w" alfred-dndbeyond-monster-workflow.go

copy:
	python2 workflow-install.py

pack:
	upx --brute alfred-dndbeyond-monster-workflow
