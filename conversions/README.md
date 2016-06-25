conversions
===========

Just a way to remember some usual conversions.

Using vars:

```go
var input = "this is a string"
var numstring = "-50"
var timeInput = time.Now()
```

```
$ go run conversions.go
input: this is a string || type: string
slice: [this is a string] || type: []string
replaced: this,is,a,string
firstStringInSlice: this
nonFirstStringInSlice: [is a string]
quotedSlice: ["this" "is" "a" "string"]
stringToBytes: [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103]
bytesToString: this is a string
timeInput: 2016-06-25 14:02:35.094671669 -0430 VET
timeInputToRFC3339: 2016-06-25T14:02:35-04:30
timeInputToUnix: 1466879555
timePeriodSeconds: 0.000334972
numstring: -50 || type: string
numstringtoint: -50 || type: int
```
