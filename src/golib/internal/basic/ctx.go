package basic

// ContextKey implementation is referenced from go stdlib:
// https://github.com/golang/go/blob/2184a394777ccc9ce9625932b2ad773e6e626be0/src/net/http/http.go#L42
type ContextKey struct {
	Name string
}

// String provides the string representation of ContextKey.
func (k ContextKey) String() string { return "golib:context_key:" + k.Name }
