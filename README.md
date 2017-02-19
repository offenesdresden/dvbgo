## 🚡 dvbgo

[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/kiliankoe/dvbgo) [![Travis](https://img.shields.io/travis/kiliankoe/dvbgo.svg?style=flat-square)](https://travis-ci.org/kiliankoe/dvbgo) [![Codecov](https://img.shields.io/codecov/c/github/kiliankoe/dvbgo.svg?style=flat-square)](https://codecov.io/gh/kiliankoe/dvbgo)

An unofficial go package giving you a few options to query a collection of publicly accessible API methods for Dresden's public transport system.

Want something like this for another language, look [no further](https://github.com/kiliankoe/vvo#libraries) 🙂

### Quick Start

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

Reading the [documentation on godoc](https://godoc.org/github.com/kiliankoe/dvbgo) will probably be of more use to see what this package can be used for than this README for now 😉

### Known Usages

- [alfred_dvb](https://github.com/kiliankoe/alfred_dvb)
- [fsr/matthias](https://github.com/fsr/matthias)
