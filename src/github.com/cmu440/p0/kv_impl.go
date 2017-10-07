// Basic key value access functions to be used in server_impl

package p0

/*
Go doesn't have the concept of class, use struct instead
accee modifier: there is public/private/protected keyword. use*/
/*small/big initials: if the first letter is big case, it is exported*/
/*otherwise it is only package access

Presence of interface: to enable polymorphism
*/

type KVStore struct{
	/* can struct fields be initialized in definition */
	kvstore map[string][]byte
}


// this function instantiates the database
func (kvstore *KVStore) init_db() {
	/* []byte is a byte array which is type as well*/
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
