package db

/*
	HOW-TO: Singleton Pattern
*/

/*
	CONTEXT:
	If you're looking at this code, chances are that you know how Golang works, but in case you
	don't, basically the folder "db" that we're currently in, is considered a "package", that's
	why we're declaring on line 1 `package db`.

	Anything created under the folder "db" is considered a part of this package.
	Any variable or function declared under this package, that starts with a lowercase letter,
	is considered a private variable or function, and if it starts with uppercase, then it's public.
	Anything variable declared outside of a function in this package, it's considered a global variable,
	to the package.
*/

// This definition of a struct "DBInstance" is considered a public definition. We'll be applying
// the Singleton pattern to this struct.
type DBInstance struct {
	conn string
}

// In order for us to have control wether the Instance has already been created, we need some kind
// of storage that allows us to see if we already have one or not.
// In this case, the private variable "currentInstance" is the storage for the instance.
// It is initialized as nil in order to establish that there's no active instance yet.
var currentInstance *DBInstance = nil

// This is the part where the magic happens.
// If someone calls the public function GetDBInstance, we'll check if we already have an instance
// if not, then we're going to initialize one; else we're going to return the one that we already
// have, thus providing the Singleton pattern.
func GetDBInstance() (*DBInstance, error) {

	if currentInstance == nil {
		currentInstance = &DBInstance{
			conn: "db connection",
		}
	}

	return currentInstance, nil
}
