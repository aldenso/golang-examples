# test

Example to remember how to test common scenarios.

```sh
go test -coverprofile=/tmp/go-code-cover -timeout 10s github.com/aldenso/golang-examples/test
```

```txt
ok  	github.com/aldenso/golang-examples/test	0.018s	coverage: 65.4% of statements
```