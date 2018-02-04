// Basic key value access functions to be used in server_impl

package p0

import "fmt"

/*
Go doesn't have the concept of class, use struct instead
access modifier: there is public/private/protected keyword. use*/
/*small/big initials: if the first letter is big case, it is exported*/
/*otherwise it is only package access

Presence of interface: to enable polymorphism
*/

type KVStore struct{
	/* can struct fields be initialized in definition */
	kvstore map[string][]byte
}


// this function instantiates the database
// if the initial is small case, then access is package level
func (kvstore *KVStore) init_db() {
	/*
		difference between make and new?
		1. make(T) always return type T; new(T) returns type *T
		2. new works for all types,
	and dynamically allocates space for a variable of that type,
	initialized to the zero value of that type, and returns a pointer to it;
		 make works as a kind of "constructor" for certain bult-in types(
	slice, map, channel)
	*/
	fmt.Println("initializing kvstore")
	kvstore.kvstore = make(map[string][]byte)
}

// put inserts a new key value pair or updates the value for a
// given key in the store
func (kvstore *KVStore) put(key string, value []byte) {
	/* kvstore is a pointer, can it be used in this way? yes*/
	kvstore.kvstore[key] = value
}

// get fetches the value associated with the key
func (kvstore *KVStore) get(key string) []byte {
	v, _ := kvstore.kvstore[key]
	return v
}
