package ops

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/x-mod/cmd"
	"github.com/x-mod/routine"
)

func MainCD(c *cmd.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("context name required")
	}
	return routine.Main(
		context.TODO(),
		routine.Command("kubectl",
			routine.ARG("config"),
			routine.ARG("use-context"),
			routine.ARG(args[0]),
		),
	)
}

func MainLS(c *cmd.Command, args []string) error {
	return routine.Main(
		context.TODO(),
		routine.Command("kubectl",
			routine.ARG("config"),
			routine.ARG("get-contexts"),
		),
	)
}

func MainRename(c *cmd.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("args required: oldname, newname")
	}
	return routine.Main(
		context.TODO(),
		routine.Command("kubectl",
			routine.ARG("config"),
			routine.ARG("rename-context"),
			routine.ARG(args[0]),
			routine.ARG(args[1]),
		),
	)
}

func MainRM(c *cmd.Command, args []string) error {
	for _, arg := range args {
		if err := routine.Main(
			context.TODO(),
			routine.Command("kubectl",
				routine.ARG("config"),
				routine.ARG("delete-context"),
				routine.ARG(arg),
			),
		); err != nil {
			return err
		}
	}
	return nil
}

func MainNCD(c *cmd.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("one namespace required")
	}
	scripts := fmt.Sprintf("kubectl config set-context $(kubectl config current-context) --namespace %s", args[0])
	return routine.Main(
		context.TODO(),
		routine.Command("sh",
			routine.ARG("-c"),
			routine.ARG(scripts),
		),
	)
}

func MainCat(c *cmd.Command, args []string) error {
	return routine.Main(
		context.TODO(),
		routine.Command("kubectl",
			routine.ARG("config"),
			routine.ARG("view"),
		),
	)
}

func MainMerge(c *cmd.Command, args []string) error {
	if len(args) == 0 {
		return nil
	}
	fns := []string{}
	for _, arg := range args {
		fn, err := filepath.Abs(arg)
		if err != nil {
			return err
		}
		fns = append(fns, fn)
	}
	scriptfmt := `
	if [ "$KUBECONFIG" == "" ]; then
		if [ -f ~/.zshrc ]; then
			echo 'export KUBECONFIG=~/.kube/config' >> ~/.zshrc
		fi
		if [ -f ~/.bashrc ]; then
			echo 'export KUBECONFIG=~/.kube/config' >> ~/.bashrc
		fi
	fi
	if [ -f ~/.zshrc ]; then
		echo 'export KUBECONFIG=$KUBECONFIG:%s' >> ~/.zshrc
	fi
	if [ -f ~/.bashrc ]; then
		echo 'export KUBECONFIG=$KUBECONFIG:%s' >> ~/.bashrc
	fi
	export KUBECONFIG=$KUBECONFIG:%s
	`
	scripts := fmt.Sprintf(scriptfmt,
		strings.Join(fns, ":"),
		strings.Join(fns, ":"),
		strings.Join(fns, ":"),
	)
	return routine.Main(
		context.TODO(),
		routine.Command("sh",
			routine.ARG("-c"),
			routine.ARG(scripts),
		),
	)
}
