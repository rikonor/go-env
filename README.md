go-env
---

_Mostly ported over from [github.com/TV4/env](https://github.com/TV4/env)_

### Installation

```bash
go get github.com/rikonor/go-env
```

### Usage

```go
package main

import env "github.com/rikonor/go-env"

func main() {
    // String
    name, err := env.String("NAME")
    if err != nil {
        ...
    }

    // Int
    age, _ := env.Int("AGE")
    
    // Float64
    height, _ := env.Float64("HEIGHT")

    // Bool
    likesIceCream, _ := env.Bool("LIKES_ICE_CREAM")

    // Duration
    timeToBed, _ := env.Duration("TIME_TO_BED")
}
```

### LICENSE

MIT
