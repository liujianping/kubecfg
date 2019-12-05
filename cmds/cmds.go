package cmds

import (
	"github.com/liujianping/kubecfg/ops"
	"github.com/x-mod/cmd"
)

func init() {
	cmd.Add(
		cmd.Path("/ls"),
		cmd.Short("list all contexts"),
		cmd.Main(ops.MainLS),
	)
	cmd.Add(
		cmd.Path("/cd"),
		cmd.Short("change context"),
		cmd.Main(ops.MainCD),
	)
	cmd.Add(
		cmd.Path("/cat"),
		cmd.Short("view ~/.kube/config"),
		cmd.Main(ops.MainCat),
	)
	cmd.Add(
		cmd.Path("/rm"),
		cmd.Short("remove context"),
		cmd.Main(ops.MainRM),
	)
	cmd.Add(
		cmd.Path("/ncd"),
		cmd.Short("change current context namespace"),
		cmd.Main(ops.MainNCD),
	)
	cmd.Add(
		cmd.Path("/rename"),
		cmd.Short("rename context"),
		cmd.Main(ops.MainRename),
	)
	cmd.Add(
		cmd.Path("/merge"),
		cmd.Short("merge a new kube config"),
		cmd.Main(ops.MainMerge),
	)
}
