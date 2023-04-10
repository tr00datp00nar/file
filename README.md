# :deciduous_tree: file

This is `file`. I have put it into a command branch for inclusion into my c Bonzai stateful command tree.

## Installation

If you just want to try it out, grab the release binary with curl and put into your PATH:

```
curl -L https://github.com/tr00datp00nar/file/releases/latest/download/file-linux-amd64 -o ~/.local/bin/file
curl -L https://github.com/tr00datp00nar/file/releases/latest/download/file-darwin-amd64 -o ~/.local/bin/file
curl -L https://github.com/tr00datp00nar/file/releases/latest/download/file-darwin-arm64 -o ~/.local/bin/file
curl -L https://github.com/tr00datp00nar/file/releases/latest/download/file-windows-amd64 -o ~/.local/bin/file
```

Or with `go`:

```shell
go install github.com/tr00datp00nar/file/cmd/file@latest
```

Composed

```go
package c

import (
	Z "github.com/rwxrob/bonzai/z"
    "github.com/tr00datp00nar/file"
)

var Cmd = &Z.Cmd{
	Name:     'c',
    Commands: []*Z.Cmd{help.Cmd, file.Cmd},
}
```

## Resources

To learn more about Bonzai stateful command trees: https://github.com/rwxrob/bonzai

To see my personal Bonzai stateful command tree: https://github.com/tr00datp00nar/c

To see the original Bonzai stateful command tree z: https://github.com/rwxrob/z
