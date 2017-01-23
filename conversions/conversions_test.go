package main

func Example1() {
	runGeneralConversions()
	//Output:
	//input: this is a string || type: string
	//slice: [this is a string] || type: []string
	//newstring: this-is-a-string || type: string
	//replaced: this,is,a,string
	//firstStringInSlice: this
	//nonFirstStringInSlice: [is a string]
	//quotedSlice: ["this" "is" "a" "string"]
	//stringToBytes: [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103]
	//bytesToString: this is a string
	//numstring: -50 || type: string
	//numstringtoint: -50 || type: int
}

func Example2() {
	runTimeConversions()
	//Output:
	//timeInput: 2017-01-21 03:30:00 -0400 VET
	//timeInputToRFC3339: 2017-01-21T03:30:00-04:00
	//timeInputToUnix: 1484983800
}
