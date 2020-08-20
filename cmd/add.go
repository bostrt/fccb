/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/bostrt/fccb/pkg/util"
	"github.com/coreos/fcct/config/v1_1"
	"github.com/spf13/cobra"
)

var inPlace bool
var inputFCCFile string
var fcc *v1_1.Config

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		inputFCCFile := args[0]

		if !util.FileExists(inputFCCFile) {
			Leave()
		}

		fcc, err = util.UnmarshalFCC(inputFCCFile)
		if err != nil {
			Leave()
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// TODO: Add output configurations
	// TODO: Add overwrite file configurations

	addCmd.PersistentFlags().BoolVarP(&inPlace, "in-place", "i", false, "Edit FCC in place (overwrite input FCC with updates)")
}

