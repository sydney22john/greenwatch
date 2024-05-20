package main

import (
	_ "github.com/mattn/go-sqlite3"
)

/*
requests:
- GET api/greenwashers: body {"domain": "value"}
 -> returns {"offenses": []}
*/

func main() {
	runServer(":8080")
}
