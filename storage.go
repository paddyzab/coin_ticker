package main

import (
	"fmt"
	"sort"
)

type Cache struct {
	items map[int64]Entry
}

type Keys []int64

func (a Keys) Len() int           { return len(a) }
func (a Keys) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Keys) Less(i, j int) bool { return a[i] < a[j] }

//Entry grouping bitcoin and ether price, storeable in Cache
type Entry struct {
	bitcoinPrice string
	etherPrice string
	ratio float64
}

//Adds entry to Cache on a given key
func (c *Cache) AddEntry(key int64, b string, e string, r float64) {
	c.items[key] = Entry{
		bitcoinPrice: b,
		etherPrice: e,
		ratio: r,
	}

	// just for testing leave me alone :-)
	for k, v := range c.items {
		fmt.Printf("key[%f] value[%s]\n", k, v)
	}
}

//Retrieves Entry from Cache based on a provided key
// Returns nil if key is not present in Cache
func (c *Cache) GetEntry(key int64) (Entry) {
	item, found := c.items[key]

	if !found {
		fmt.Printf("We cannot find under key: [%s]", key)
	}

	return item
}

//Clears Cache, by replacing data with empty map
func (c *Cache) Clear() {
	c.items = map[int64]Entry{}
}

func (c *Cache) Size() int {
	return len(c.items)
}

//Returns last element of the
func (c *Cache) GetLast() (Entry) {
	var keys Keys
	for k := range c.items {
		keys = append(keys, k)
	}

	sort.Sort(keys)
	lElm := keys[len(keys) -1]

	return c.items[lElm]
}

//Returns clear Cache
func New() *Cache {
	i := make(map[int64]Entry)
	return &Cache{
		items: i,
	}
}
