package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/swk67018/bstset/bst"
)

var setCmd = &cobra.Command{
	Use:   "set target...",
	Short: "Set specified options.",
	Long: `Set specified options.
Target must be followed "name:value" format.`,
	Example: `  bstset set 'bst.feature.macros:1'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.PrintErrf("Specify one or more targets.")
			return
		}

		file, _ := cmd.Flags().GetString("file")

		bstOpt := bst.Option{}
		for _, arg := range args {
			token := strings.Split(arg, ":")
			if len(token) != 2 {
				cmd.PrintErrf("\"%s\" is not in the correct format. Follow \"name:value\" format.\n", arg)
				return
			}
			name := token[0]
			value := token[1]
			bstOpt[name] = value
		}

		bstCfg := bst.NewConfig()
		if err := bstCfg.Load(file); err != nil {
			cmd.PrintErrf("Failed to open \"%s\".\n", file)
			return
		}
		bstCfg.Apply(bstOpt)
		if err := bstCfg.Write(file); err != nil {
			cmd.PrintErrf("Failed to output data to \"%s\".", file)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringP("file", "f", bst.DEFAULT_CONFIG_PATH, "BlueStacks configuration file path")
}
