package storage

import "time"

//Cache stores Entry items.
type Cache struct {
	items []Entry
}

//Entry groups bitcoin and ether price, storeable in Cache.
type Entry struct {
	CoinData  Results
	Timestamp time.Time
}

// Results contain raw results in form of a hash map and possible errors.
type Results struct {
	Result map[string]float64
	Errors []error
}

//AddEntry adds to Cache underlining slice.
func (c *Cache) AddEntry(res Results, t time.Time) {
	c.items = append(c.items, Entry{
		CoinData:  res,
		Timestamp: t,
	})
}

//GetEntry retrieves Entry from Cache based on a provided position in the slice.
func (c *Cache) GetEntry(position int) Entry {
	item := c.items[position]

	return item
}

//Clear clears Cache, by replacing data with empty map.
func (c *Cache) Clear() {
	c.items = []Entry{}
}

//Size returns current size of Cache.
func (c *Cache) Size() int {
	return len(c.items)
}

//GetLast returns last element of the Cached array.
func (c *Cache) GetLast() Entry {
	if len(c.items) == 0 {
		return Entry{}
	}
	return c.items[len(c.items)-1]
}

//NewCache returns clear Cache.
func NewCache() *Cache {
	i := make([]Entry, 0)
	return &Cache{
		items: i,
	}
}
