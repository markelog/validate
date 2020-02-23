// +build tools

package main

import (
	_ "github.com/canthefason/go-watcher/cmd/watcher"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/mgechev/revive"
	_ "github.com/securego/gosec"
)
