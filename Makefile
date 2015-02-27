run: deps
	GOPATH=$(CURDIR)/.go go run third/main.go \
		 -c $(CURDIR)/third/etc/deploy42/config.yml

deps:
	cd third; GOPATH=$(CURDIR)/.go go get -d
	rm -Rf $(CURDIR)/.go/src/github.com/andrerocker/deploy42
	ln -s $(CURDIR) $(CURDIR)/.go/src/github.com/andrerocker/deploy42
