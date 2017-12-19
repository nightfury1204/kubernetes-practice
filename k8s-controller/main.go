package main

import (
	"log"
	"kubernetes-practice/k8s-controller/cmd"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

func main() {

	home := homedir.HomeDir()

	rootCmd := &cobra.Command{
		Use: "controller",
		Short:"Create k8s controller",
	}

	rootCmd.PersistentFlags().StringVarP(&cmd.Kubeconfig, "configPath", "c", home+"/.kube/config", "kube config path")
	rootCmd.AddCommand(cmd.CreateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
