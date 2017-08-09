# vitess-sqlparser

simply sqlparser powered by vitess  
this library inspired by https://github.com/xwb1989/sqlparser

(original source : https://github.com/youtube/vitess/tree/master/go/vt/sqlparser)

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
