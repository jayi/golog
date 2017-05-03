# golog
Simply go logging library

## Example

```go
package main

import "github.com/jayi/golog"

func main() {
    // default level is trace
    golog.SetLevel(golog.DebugLevel)
    
    golog.Tracef("%s message\n", "trace")
    golog.Debug("debug message\n")
    golog.Infoln("info message\n")
}
```

output:
```
2017/05/03 15:16:37 main.go:10: [DEBUG] debug message
2017/05/03 15:16:37 main.go:11: [INFO] info message
```

## Levels

Currently supported levels are

- TRACE
- DEBUG
- INFO
- WARN
- ERROR
- PANIC
