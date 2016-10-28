Gorilla
=======

Just to remember a way to work with gorilla mux and gorilla handlers.


Not found example.
```
$ curl -i 'http://localhost:8080/api/user/aldenso/bad'
HTTP/1.1 404 Not Found
Date: Fri, 28 Oct 2016 17:22:07 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```

Post example.
```
$ curl -X POST -i 'http://localhost:8080/api/user' -d '{"name": "aldenso"}'
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Location: /api/user/id/0
Date: Fri, 28 Oct 2016 17:22:50 GMT
Content-Length: 0
```

Get example.
```
$ curl -i 'http://localhost:8080/api/user'
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 28 Oct 2016 17:23:16 GMT
Content-Length: 119

[
    {
        "id": 0,
        "name": "aldenso",
        "created_at": "2016-10-28T13:22:50.403090454-04:00"
    }
]
```

Get by username (previously added some users).
```
$ curl -i 'http://localhost:8080/api/user/name/aldenso'
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 28 Oct 2016 17:25:25 GMT
Content-Length: 236

[
    {
        "id": 0,
        "name": "aldenso",
        "created_at": "2016-10-28T13:22:50.403090454-04:00"
    },
    {
        "id": 4,
        "name": "aldenso",
        "created_at": "2016-10-28T13:25:19.739885948-04:00"
    }
]
```

Get by ID.
```
$ curl -i 'http://localhost:8080/api/user/id/4'
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 28 Oct 2016 17:25:54 GMT
Content-Length: 95

{
    "id": 4,
    "name": "aldenso",
    "created_at": "2016-10-28T13:25:19.739885948-04:00"
}
```

Log output.
```
::1 - - [28/Oct/2016:13:24:15 -0400] "POST /api/name HTTP/1.1" 404 0
::1 - - [28/Oct/2016:13:24:48 -0400] "POST /api/user HTTP/1.1" 201 0
::1 - - [28/Oct/2016:13:24:51 -0400] "GET /api/user HTTP/1.1" 200 352
```