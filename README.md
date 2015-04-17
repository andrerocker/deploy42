deploy42
=========
Read a YAML and export this as a 'Command as a Service' :p 

https://github.com/andrerocker/deploy42

##### Example

Based on a simple yaml
```yaml
daemon:
  bind: 127.0.0.1
  port: 8888
  http:
    pipe: true
    vars: true
  load:
    - /etc/deploy42/config.d/*.yml
    - /var/www/*/config/deploy42.yml

namespaces:
  - endpoint: admin-ops
    chaining:
      - a_audit_filter
      - a_security_filter
    commands:
      process:
        - get: ps -ef | grep {process}
          put: kill {process}
          delete: kill -9 {process}

  - endpoint: free-path
    commands:
      log:
        - get: tail -f {log}
        
      echo:
        - put: cat -
```

You can do this
```console
$ curl http://server:8888/admin-ops/process/ruby
andrero+ 1337 42   5  11:18 pts/25   00:00:01 ruby bin/rails s
andrero+ 1338 42   29 11:18 pts/26   00:00:01 ruby bin/rails c

$ curl -X PUT http://server:8888/process/1338
$

$ curl http://server:8888/admin-ops/process/ruby
andrero+ 1337 42   5  11:18 pts/25   00:00:01 ruby bin/rails s
```

```console
$ curl http://server:8888/free-path/log/log/development.log
Started GET "/document/42" for 127.0.0.1 at 2014-11-30 03:20:32 -0200
Processing by DocumentController#show as HTML
  Parameters: {"id"=>"42"}
  User Load (0.4ms)  SELECT  "users".* FROM "users"  WHERE "users"."id" = 1337
  Doc Load (0.6ms)  SELECT  "docs".* FROM "docs"  WHERE "docs"."id" = 42 LIMIT 1
  Rendered document/show.html.erb within layouts/application (0.2ms)
Completed 200 OK in 97ms (Views: 94.6ms | ActiveRecord: 1.0ms)
```

```console
$ echo "Andre Master of Universe" | curl -s -T - http://server:8888/free-path/echo/yeah
Andre Master of Universe
```

##### TODO List

- [x] New yaml model
- [x] Command namespaces
- [x] Base authetication filters
- [ ] Base audit filters
- [ ] Routes reloader SIG USR2 and HTTP ROUTE
- [ ] Package with fpm
- [ ] Logger

##### Running locally

- dependencies: make and go1.4+
- run: just run "make"
- have a fun!!! \o/
