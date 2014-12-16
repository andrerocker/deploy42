run:
	GOPATH=$(CURDIR)/.go go run src/main.go


deps:
	cd src; GOPATH=$(CURDIR)/.go go get -d
