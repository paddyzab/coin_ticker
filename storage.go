package main

import "fmt"

type Cache struct {
	items map[string]Entry
}

//Entry grouping bitcoin and ether price, storeable in Cache
type Entry struct {
	bitcoinPrice string
	etherPrice string
}

//Adds entry to Cache on a given key
func (c *Cache) AddEntry(key string, b string, e string) {
	c.items[key] = Entry{
		bitcoinPrice: b,
		etherPrice: e,
	}

	// just for testing leave me alone :-)
	for k, v := range c.items {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
}

//Retrieves Entry from Cache based on a provided key
// Returns nil if key is not present in Cache
func (c *Cache) GetEntry(key string) (Entry) {
	item, found := c.items[key]

	if !found {
		fmt.Printf("We cannot find under key: [%s]", key)
	}

	return item
}

//Clears Cache, by replacing data with empty map
func (c *Cache) Clear() {
	c.items = map[string]Entry{}
}

func (c *Cache) Size() int {
	return len(c.items)
}

//Returns clear Cache
func New() *Cache {
	i := make(map[string]Entry)
	return &Cache{
		items: i,
	}
}
