# easy-timeout
easy timeout in golang

## Usage
### Install
```
go get github.com/flybikeGx/easy-timeout
```
### Use
```
import "github.com/flybikeGx/easy-timeout/timelimit"

    ok := timelimit.Run(1*time.Second, func() {
        // your closure, main thread will be blocked
    })
    if !ok {
        // handle time out
    }
```
## Guarantee
1. Thread safe
2. Stop at the exact time as golang timer

## Notice
Closure safety should take care of