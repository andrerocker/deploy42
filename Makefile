run: deps
	GOPATH=$(CURDIR)/.go go run examples/main.go \
		 -c $(CURDIR)/examples/etc/deploy42/base.yml \
		 -a $(CURDIR)/examples/etc/deploy42/auth.yml 

deps:
	cd examples; GOPATH=$(CURDIR)/.go go get -d
	rm -Rf $(CURDIR)/.go/src/github.com/andrerocker/deploy42
	ln -s $(CURDIR) $(CURDIR)/.go/src/github.com/andrerocker/deploy42
