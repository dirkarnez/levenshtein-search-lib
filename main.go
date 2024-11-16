import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type User struct {
  ID           uint           // Standard field for the primary key
  Name         string         // A regular string field
  Email        *string        // A pointer to a string, allowing for null values
  Age          uint8          // An unsigned 8-bit integer
}

func Levenshtein () {
	
}
func main() {
	users := []User{}
	
	// // Create a cache with a default expiration time of 5 minutes, and which
	// // purges expired items every 10 minutes
	// c := cache.New(5*time.Minute, 10*time.Minute)

	// // Set the value of the key "foo" to "bar", with the default expiration time
	// c.Set("users", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("users", users, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	usersFound, found := c.Get("users")
	if found {
		Levenshtein(
			usersFound.([]User), 
			func (u *User) {
				return u.Name
			},
			func (u *User) {
				return u.Email
			},
		)
	}

	// // This gets tedious if the value is used several times in the same function.
	// // You might do either of the following instead:
	// if x, found := c.Get("foo"); found {
	// 	foo := x.(string)
	// 	// ...
	// }
	// // or
	// var foo string
	// if x, found := c.Get("foo"); found {
	// 	foo = x.(string)
	// }
	// // ...
	// foo can then be passed around freely as a string

	// // Want performance? Store pointers!
	// c.Set("foo", &MyStruct, cache.DefaultExpiration)
	// if x, found := c.Get("foo"); found {
	// 	foo := x.(*MyStruct)
	// 		// ...
	// }
}
