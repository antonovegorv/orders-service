package cache

import (
	"time"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/antonovegorv/orders-service/database"
	"github.com/antonovegorv/orders-service/internal/model"
)

// A cache with orders.
var Orders ttlcache.SimpleCache

// Init initializes an Orders global package variable, setups cache properties
// and fills cache up with data from the Database.
func Init() error {
	// Initialize a new cache.
	Orders = ttlcache.NewCache()

	// Set the expiration time for the cache. (No expiration in this case).
	Orders.SetTTL(0 * time.Second)

	// Read all values from the Database.
	rows, err := database.DB.Query("SELECT data FROM orders")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Iterate through the rows.
	for rows.Next() {
		// An order instance to scan the result in.
		var order model.Order
		if err := rows.Scan(&order); err != nil {
			return err
		}

		// Set the value to cache.
		Orders.Set(order.UUID.String(), order)
	}

	return nil
}
