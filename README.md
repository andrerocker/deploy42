deploy.go
=========

https://github.com/andrerocker/deploy.go

### Description

- Read a YAML and export this as a 'Command as a Service' :p


##### Example

- Based on a simple yaml

```
daemon:
  bind: 127.0.0.1
  port: 8888

commands:
  process:
    - get: ps -ef | grep {process}
      put: kill {process}
      delete: kill -9 {process}

  log:
    - get: tail -f {log}
```

- You can do this

```
$ curl http://server:8888/process/ruby
andrero+ 1337 45   5  11:18 pts/25   00:00:01 ruby bin/rails s
andrero+ 1337 45   29 11:18 pts/26   00:00:01 ruby bin/rails c

```


