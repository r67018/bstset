package cmd

import (
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/swk67018/bstset/bst"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load file",
	Short: "Load configuration file and apply.",
	Long: `Load configuration file and apply options written in the file.
The file format is as follows.
  bstConfigPath: string
    Path to BlueStacks configuration file
  targets: {name: string, value: string}[]
    Array of option you want to set.
`,
	Example: `  bstset load ./config.json`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Println("It is allowed only one file to load.")
			return
		}

		file := args[0]

		cfg, err := unmarshalConfig(file)
		if err != nil {
			cmd.Printf("Failed to load %s.\nCheck if %s exists and is correct format.\n", file, file)
			return
		}

		bstCfg := bst.NewConfig()
		if err := bstCfg.Load(cfg.BstConfigPath); err != nil {
			cmd.PrintErrf("Failed to open \"%s\".\n", file)
			return
		}
		bstCfg.Apply(convertTargetsToBstOption(cfg.Targets))
		if err := bstCfg.Write(cfg.BstConfigPath); err != nil {
			cmd.PrintErrf("Failed to output data to \"%s\".", file)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}

type Target struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Config struct {
	BstConfigPath string   `json:"bstConfigPath"`
	Targets       []Target `json:"targets"`
}

func unmarshalConfig(path string) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func convertTargetsToBstOption(ts []Target) bst.Option {
	o := bst.Option{}
	for _, t := range ts {
		o[t.Name] = t.Value
	}
	return o
}
