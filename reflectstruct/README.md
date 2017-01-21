reflectstruct
=============

Example on how to use reflect to get tags, fields and types from a flat struct.

Data input:
```
{Name: "aldenso", ID: 1, Email: "aldenso@gmail.com", Score: 50.50}
```

Data output:
```
slice of tags: [name id email score]
slice of fields: [aldenso 1 aldenso@gmail.com 50.5]
slice of types: [Name ID Email Score]
```