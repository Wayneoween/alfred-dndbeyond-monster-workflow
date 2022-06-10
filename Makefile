all: clean build

release: build gopack compress

clean:
	$(RM) alfred-dndbeyond-monster-workflow
	$(RM) alfred-dndbeyond-monster-workflow.upx
	$(RM) "workflow/D&DBeyond\ Monster\ Search.alfredworkflow"
	$(RM) -r testenv/*

build:
	echo "Building alfred-dndbeyond-monster-workflow..."
	go build -v -ldflags="-s -w" ./...

gopack:
	upx --brute alfred-dndbeyond-monster-workflow

compress:
	zip -r workflow/DnDBeyondMonsterSearch.alfredworkflow * -x build -x .git -x demo.gif
