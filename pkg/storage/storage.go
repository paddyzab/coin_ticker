package storage

import "time"

//Stores Entry items in array.
type Cache struct {
	items []Entry
}

//Entry groups bitcoin and ether price, storeable in Cache.
type Entry struct {
	BitcoinPrice string
	EtherPrice   string
	Ratio        float64
	Timestamp    time.Time
}

//Adds entry to Cache underlining slice.
func (c *Cache) AddEntry(b string, e string, r float64, t time.Time) {
	c.items = append(c.items, Entry{
		BitcoinPrice: b,
		EtherPrice:   e,
		Ratio:        r,
		Timestamp:    t,
	})
}

//Retrieves Entry from Cache based on a provided position in the slice.
func (c *Cache) GetEntry(position int) Entry {
	item := c.items[position]

	return item
}

//Clears Cache, by replacing data with empty map.
func (c *Cache) Clear() {
	c.items = []Entry{}
}

//Returns current size of Cache.
func (c *Cache) Size() int {
	return len(c.items)
}

//Returns last element of the Cached array.
func (c *Cache) GetLast() Entry {
	if len(c.items) == 0 {
		return Entry{}
	}
	return c.items[len(c.items)-1]
}

//Returns clear Cache.
func NewCache() *Cache {
	i := make([]Entry, 0)
	return &Cache{
		items: i,
	}
}