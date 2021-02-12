package command

import "github.com/spf13/cobra"

type root struct{}

func newRoot() *cobra.Command {
	return &cobra.Command{
		Use:     "Reavers",
		Short:   "Core business logic of Reavers",
		Example: "reavers",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
}
