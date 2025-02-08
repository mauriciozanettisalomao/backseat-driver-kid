// Copyright The Linux Foundation and each contributor to LFX.
// SPDX-License-Identifier: MIT

package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "backseat-driver-kid",
	Short: `
	backseat-driver-kid: 
	
	A CLI tool that’s like having a curious toddler in the back seat of your brain—constantly asking questions so you don’t have to! 
	
	Let it do the thinking while you focus on the driving (or coding).`,
}

// Execute is the main function of the cli tool
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().String("log-level", "debug", "Flag to indicate debug mode.")
	err := viper.BindPFlag("logLevel", rootCmd.PersistentFlags().Lookup("log-level"))
	if err != nil {
		log.Fatalf("error binding flag: %s | error %v", "debug", err)
	}

	rootCmd.PersistentFlags().String("log-format", "json", "Log format")
	err = viper.BindPFlag("logFormat", rootCmd.PersistentFlags().Lookup("log-format"))
	if err != nil {
		log.Fatalf("error binding flag: %s | error %v", "logFormat", err)
	}

}
