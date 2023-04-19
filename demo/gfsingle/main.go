package main

import (
	_ "gfsingle/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gfsingle/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
