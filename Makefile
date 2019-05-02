all: clean build copy

release: clean build pack copy

clean:
	$(RM) alfred-dndbeyond-monster-workflow
	$(RM) alfred-dndbeyond-monster-workflow.upx

build:
	go build -ldflags="-s -w" alfred-dndbeyond-monster-workflow.go icons.go types.go

copy:
	python2 workflow-install.py

pack:
	upx --brute alfred-dndbeyond-monster-workflow
