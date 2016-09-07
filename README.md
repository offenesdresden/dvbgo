## 🚡 dvbgo

[![Travis](https://img.shields.io/travis/kiliankoe/dvbgo.svg?style=flat-square)]()

An unofficial go package giving you a few options to query a collection of publicly accessible API methods for Dresden's public transport system.

Similar libs also exist for [Node](https://github.com/kiliankoe/dvbjs), [Python](https://github.com/kiliankoe/dvbpy), [Swift](https://github.com/kiliankoe/DVB) and [Ruby](https://github.com/kiliankoe/dvbrb) 😊

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

Rest of the docs (this README) will follow as this package starts to form 🙃
