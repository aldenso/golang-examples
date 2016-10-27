Gorilla
=======

Just to remember a way to work with gorilla mux and gorilla handlers.


Not found example.
```
$ curl -i 'http://localhost:8080/api/name/aldenso/bad'
HTTP/1.1 404 Not Found
Date: Thu, 27 Oct 2016 19:34:34 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```

Post example.
```
$ curl -X POST -i 'http://localhost:8080/api/name' -d '{"name": "aldenso"}'
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Location: /api/name/aldenso
Date: Thu, 27 Oct 2016 19:36:29 GMT
Content-Length: 0
```

Get example.
```
$ curl -i 'http://localhost:8080/api/name'
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 27 Oct 2016 19:38:11 GMT
Content-Length: 101

[
    {
        "name": "aldenso",
        "created_at": "2016-10-27T15:38:06.27562058-04:00"
    }
]
```

Log output.
```
::1 - - [27/Oct/2016:15:38:02 -0400] "GET /api/name HTTP/1.1" 200 4
::1 - - [27/Oct/2016:15:38:06 -0400] "POST /api/name HTTP/1.1" 201 0
::1 - - [27/Oct/2016:15:38:11 -0400] "GET /api/name HTTP/1.1" 200 101
```