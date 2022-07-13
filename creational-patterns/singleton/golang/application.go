package main

import (
	"golang/db"
)

/*
	CONTEXT:
 So, in this Singleton pattern example, for some unknown reason, we're calling
 the creation method of the DB Provider twice, and to successfully apply the pattern,
 we're required to resolve the same instance of the DB Provider.
*/

/*
 SOLUTION:
 Go to file > db/main.go for further explanation
*/

func main() {
	database, err := db.NewDBInstance()
	if err != nil {
		panic(err)
	}

	/* some unknown logic... */

	database2, err := db.NewDBInstance()
	if err != nil {
		panic(err)
	}
}
