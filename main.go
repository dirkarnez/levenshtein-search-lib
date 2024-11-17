package main

import (
	"fmt"
	"github.com/agnivade/levenshtein"
	"github.com/patrickmn/go-cache"
	"github.com/samber/lo"
	// "math/rand"
	"math"
	"strings"
	"time"
)

type User struct {
  ID           uint           // Standard field for the primary key
  Name         string         // A regular string field
  Email        *string        // A pointer to a string, allowing for null values
  Age          uint8          // An unsigned 8-bit integer
}

func LevenshteinRatio (searchString, data string) float64 {
	return 1.0 - (float64(levenshtein.ComputeDistance(searchString, data)) / float64(math.Max(float64(len(searchString)), float64(len(data)))))
}

func Levenshtein[T any, Slice ~[]T](searchString string, collection Slice, getters ...func(item T) string) Slice {
	for _, getter := range getters {
	
	}

	lo.Map(collection, func(_, index int) bool {
		return strconv.FormatInt(x, 10)
	})


	for _, getter := range getters {
		data := getter(x)
		ratio := LevenshteinRatio(searchString, data)
		fmt.Println("searchString", searchString, "data", data, "LevenshteinRatio", ratio)
	}
	
	return lo.Filter(collection, func(x T, index int) bool {

		return true
	})
}

// func GetRandomNumber(min, max int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(max - min) + min
// }

// func Shuffle[L interface{}](array []L) {
// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
// }

// func GetRandomName() string {
// 	a := []string{"1213"}
// 	Shuffle(&a)
// 	return a[0]
// }

var names = []string{
	"Isaac Thomson","Oliver Henderson","Bella Harris","Olivia Slater","Heather Skinner","Joe Cornish","Una Paige","Alexandra Turner","Jason Sharp","Heather Jones","Samantha Vaughan","Rebecca Sanderson","Kevin Bond","Gordon Gibson","Yvonne Scott","Dylan Brown","Simon Churchill","Irene Coleman","Peter Miller","Nicola Hughes","Fiona Forsyth","Virginia Mackay","Anne Parr","Gordon Marshall","Sophie Young","Anne Hart","Cameron James","Colin Davidson","Frank Graham","Joseph Murray","Audrey Campbell","Samantha Mackenzie","Felicity Manning","Bernadette Mackenzie","Felicity Lambert","Carl Burgess","Natalie Wilkins","Joseph Wilson","Nicholas Scott","Caroline Hill","Keith MacDonald","Andrew Payne","Dan Fisher","Blake Russell","Leonard Kelly","Stephen Bond","Victor MacDonald","Jacob Underwood","Jonathan Rees","Megan Robertson","Warren Fisher","Robert Campbell","Olivia Metcalfe","Nicola Brown","Sue Peake","Joanne Avery","Jacob Powell","Lillian Bailey","Emily Howard","Rachel Sharp","Peter Peters","Lisa Gray","Ryan Cameron","Rose Bond","Frank Taylor","Sophie Paige","Austin Lyman","Maria Ross","Victoria Thomson","Ella Murray","Andrea Davidson","Lucas Lambert","Katherine Gray","David Davies","Sarah Walker","Robert Quinn","Jessica McLean","Amy Burgess","David MacDonald","Samantha Ball","Amanda Chapman","Wanda Brown","Diana Wright","Sebastian Ogden","Audrey Wilson","Elizabeth MacLeod","Wanda Hughes","Anna Russell","Claire Cameron","Tracey Duncan","Kevin Hodges","Benjamin Gibson","Matt Piper","Theresa Paige","Austin Henderson","Michelle Brown","Anna Berry","Isaac Gill","Evan Reid","Yvonne Metcalfe",
}

func main() {
	users := []User{}
	for i, name := range names {
		id := uint(i + 1)
		email := strings.ToLower(strings.ReplaceAll(name, " ", "")) + "@gmail.com"
		fmt.Println("id", id, "name", name, "email", email)
		users = append(users, User{ID: id, Name: name, Email: &email})
	}
	
	
	// // Create a cache with a default expiration time of 5 minutes, and which
	// // purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// // Set the value of the key "foo" to "bar", with the default expiration time
	// c.Set("users", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("users", users, cache.NoExpiration)

	// // Get the string associated with the key "foo" from the cache
	usersFound, found := c.Get("users")
	if found {
		usersFoundCasted := usersFound.([]User)
		fmt.Println("found", len(usersFoundCasted))
		Levenshtein(
			"olv",
			usersFoundCasted, 
			func (u User) string {
				return u.Name
			},
			func (u User) string {
				return *u.Email
			},
		)

	} else {
		fmt.Println("not found")
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
