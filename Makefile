all: clean build copy

ZIP := $(shell which 7za)

release: build gopack copy compress

clean:
	$(RM) alfred-dndbeyond-monster-workflow
	$(RM) alfred-dndbeyond-monster-workflow.upx
	$(RM) "workflow/D&DBeyond\ Monster\ Search.alfredworkflow"

build:
	echo "Building alfred-dndbeyond-monster-workflow..."
	go build -ldflags="-s -w" alfred-dndbeyond-monster-workflow.go icons.go types.go

gopack:
	upx --brute alfred-dndbeyond-monster-workflow

copy:
	python2 workflow-install.py

compress:
	$(ZIP) a "workflow/D&DBeyond\ Monster\ Search.alfredworkflow" * -xr'!build' -xr'!.git' -xr'!demo.gif'
