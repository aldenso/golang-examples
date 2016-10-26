Gorilla
=======

Just to remember a way to work with gorilla mux and gorilla handlers.


Not found example.
```
$ curl -i 'http://localhost:8080/api/name/aldenso/bad'
HTTP/1.1 404 Not Found
Date: Wed, 26 Oct 2016 21:12:29 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```

OK example.
```
$ curl -i 'http://localhost:8080/api/name/aldenso'
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 26 Oct 2016 21:12:34 GMT
Content-Length: 25

{
    "name": "aldenso"
}
```

Log output.
```
::1 - - [26/Oct/2016:17:12:29 -0400] "GET /api/name/aldenso/bad HTTP/1.1" 404 0
::1 - - [26/Oct/2016:17:12:34 -0400] "GET /api/name/aldenso HTTP/1.1" 200 25
```