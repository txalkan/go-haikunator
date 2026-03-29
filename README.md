# go-haikunator

[![CI](https://github.com/yelinaung/go-haikunator/actions/workflows/ci.yml/badge.svg)](https://github.com/yelinaung/go-haikunator/actions/workflows/ci.yml)

Heroku-like memorable random name generator. Golang port of [haikunator](https://github.com/usmanbashir/haikunator).

By default the generator provides 8645 unique combinations.

```
sparkling-cherry
snowy-brook
bitter-darkness
```

View the [docs](https://pkg.go.dev/github.com/yelinaung/go-haikunator).

## Example

```go
package main

import (
	"fmt"

	haikunator "github.com/yelinaung/go-haikunator"
)

func main() {
	h := haikunator.New()
	fmt.Println(h.Haikunate()) // e.g. "sparkling-cherry"

	// For reproducible output, use a seed:
	h2 := haikunator.NewWithSeed(42)
	fmt.Println(h2.Haikunate())
}
```

## Other Languages

Haikunator is also available in other languages. Check them out:
- Node: https://github.com/AtroxDev/haikunatorjs
- Python: https://github.com/AtroxDev/haikunator
- Ruby: https://github.com/usmanbashir/haikunator

## License

MIT
