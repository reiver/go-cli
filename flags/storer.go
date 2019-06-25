package cliflags

// Storer represents something that can store a key-value pair.
//
// This makes it so you can pass custom types (that are able to store a key-value pair) to cliflags.Parse()
// as the target.
type Storer interface {
	Store(string, string) error
}
