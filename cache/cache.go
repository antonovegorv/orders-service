package cache

import (
	"time"

	"github.com/ReneKroon/ttlcache/v2"
)

// A cache with orders.
var Orders ttlcache.SimpleCache

// Init initializes an Orders global package variable, setups cache properties
// and fills cache up with data from the Database.
func Init() {
	// Initialize a new cache.
	Orders = ttlcache.NewCache()

	// Set the expiration time for the cache. (No expiration in this case).
	Orders.SetTTL(0 * time.Second)

	// Read all values from the Database.
}
