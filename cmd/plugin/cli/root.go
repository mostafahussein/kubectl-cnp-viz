package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mostafahussein/kubectl-cnp-viz/pkg/plugin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
	opts                  plugin.DiagramCaptureOptions
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "kubectl-cnp-viz",
		Short:         "Visualize Cilium Network Policies",
		Long:          `Visualize Cilium Network Policies`,
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.MaximumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			cnpName := args[0]
			if err := plugin.RunPlugin(cnpName, KubernetesConfigFlags, opts); err != nil {
				return errors.Unwrap(err)
			}

			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	KubernetesConfigFlags = genericclioptions.NewConfigFlags(false)
	KubernetesConfigFlags.AddFlags(cmd.Flags())

	cmd.Flags().Float64Var(&opts.Scale, "scale", 1.00, "Scaling factor (default: 1.00)")
	cmd.Flags().Float64Var(&opts.TranslateX, "x", 340, "Reposition the diagram along the horizontal axis (default: 340)")
	cmd.Flags().Float64Var(&opts.TranslateY, "y", 30, "Reposition the diagram along the vertical axis (default: 30)")
	cmd.Flags().StringVar(&opts.OutputDir, "outputdir", os.TempDir(), "Directory to store output files (default: system temp dir)")

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cmd
}

func InitAndExecute() {
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.AutomaticEnv()
}
