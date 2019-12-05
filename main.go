package main

import (
	"github.com/x-mod/build"
	"github.com/x-mod/cmd"

	_ "github.com/liujianping/kubecfg/cmds"
)

func main() {
	cmd.Version(build.String())
	cmd.Add(
		cmd.Name("kubecfg"),
		cmd.Short("kubecfg, more convenient tool for kube config"),
	)
	cmd.Execute()
}
