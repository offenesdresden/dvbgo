## :aerial_tramway: dvbgo

An unofficial go package giving you a few options to query a collection of publicly accessible API methods for Dresden's public transport system.

Similar libs also exist for [Node](https://github.com/kiliankoe/dvbjs), [Python](https://github.com/kiliankoe/dvbpy), [Swift](https://github.com/kiliankoe/DVB) and [Ruby](https://github.com/kiliankoe/dvbrb) :blush:

**dvbgo is WIP and not even close to something usable at the moment. I'm also just playing around with go using this as my advanced hello-world example and learning as I go. No pun intended.**

```go
import "github.com/kiliankoe/dvbgo"
```

You can now do the following for example:

```go
func main() {
	deps, _ := Monitor("Helmholtzstraße", 0, "")
	fmt.Println(deps)
}
```

Rest of the docs (this README) will follow as this package starts to form :upside_down:
