# mgo-oid
A unique id generator that produces ids like mongodb object id. This generator implementation refers the source code of [mgo project](https://github.com/go-mgo/mgo), that's why it has mgo as its name prefix.

## ObjectId
Object id is a globally unique identifier for object in mongodb. It consists of 12 bytes, divided as follows:
* a 4-byte value representing the seconds since the Unix epoch
* a 3-byte machine identifier
* a 2-byte process id
* a 3-byte counter, starting with a random value

## Installation
Use the `go` command:

	$ go get github.com/coolbed/mgo-oid

## Example
```go
package main

import (
	"fmt"

	"github.com/coolbed/mgo-oid"
)

func main() {
	objectID := oid.NewOID()
	fmt.Println("object id:", objectID.String())
	fmt.Println("object timestamp", objectID.Timestamp())
}
``` 