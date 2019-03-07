package ht

import (
	"errors"
	"hash/fnv"
)

//HashTable represent hash-table structure
type HashTable struct {
	keys   []uint32
	values []interface{}
}

// New constructs hash-table
func New() *HashTable {
	return &HashTable{}
}

// Add adds values to hash-table
func (ht *HashTable) Add(key string, value interface{}) {
	hash := ht.hash(key)

	ht.keys = append(ht.keys, hash)
	ht.values = append(ht.values, value)
}

// Get returns value by key
func (ht *HashTable) Get(key string) (interface{}, error) {
	hash := ht.hash(key)

	for i := range ht.keys {
		key := ht.keys[i]

		if key == hash {
			return ht.values[i], nil
		}
	}

	return nil, errors.New("Value not found")
}

// Remove deletes element from hash-table
func (ht *HashTable) Remove(key string) error {
	hash := ht.hash(key)

	for i := range ht.keys {
		key := ht.keys[i]

		if key == hash {
			copy(ht.keys[i:], ht.keys[i+1:])
			ht.keys[len(ht.keys)-1] = 0
			ht.keys = ht.keys[:len(ht.keys)-1]

			copy(ht.values[i:], ht.values[i+1:])
			ht.values[len(ht.values)-1] = nil
			ht.values = ht.values[:len(ht.values)-1]

			return nil
		}
	}

	return errors.New("Value not found")
}

func (ht *HashTable) hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
