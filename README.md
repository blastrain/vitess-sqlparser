# vitess-sqlparser

Simply SQL and DDL parser for Go (powered by vitess and TiDB )
this library inspired by https://github.com/xwb1989/sqlparser

(original source : https://github.com/youtube/vitess/tree/master/go/vt/sqlparser)

# Why

[xwb1989/sqlparser](https://github.com/xwb1989/sqlparser) is famous sql parser in Go.  
But it cannot parse some query (like offset or bulk insert...) because it customizes vitess's sql parser.  

Also, some libraries use from vitess sql parser directly. But vitess's sql parser only partial supports DDL parsing.  

We want to perfectly support parsing for SQL and DDL.  
Therefore we use vitess sql parser directly and also use TiDB parser for DDL parsing. 

# Compare SQL parser libraries in Go

| library | supports offset (or other complexity) query | supports DDL |
|:---:|:---:|:---:|
|xwb1989/sqlparser |✗ | △|
|zhenjl/sqlparser | ○|△ |
|knocknote/vitess-sqlparser|○|○|

# Installation

```
go get -u github.com/knocknote/vitess-sqlparser
```

# Examples

```go
package main

import (
 	"fmt"
	"github.com/knocknote/vitess-sqlparser/sqlparser"
)

func main() {
	stmt, err := sqlparser.Parse("select * from user_items where user_id=1 order by created_at limit 3 offset 10")
	if err != nil {
		panic(err)
	}
	fmt.Printf("stmt = %+v\n", stmt)
}

```
