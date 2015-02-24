run: deps
	GOPATH=$(CURDIR)/.go go run src/main.go -config=$(CURDIR)/etc/deploy42/config.yml

deps:
	cd src; GOPATH=$(CURDIR)/.go go get -d
