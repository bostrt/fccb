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
	"github.com/bostrt/fccb/pkg/systemd"
	"github.com/bostrt/fccb/pkg/util"
	"io/ioutil"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// unitCmd represents the unit command
var unitCmd = &cobra.Command{
	Use:   "unit",
	Short: "Add systemd unit to FCC",
	Long: `TODO: Long desc with examples`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: Use basename for unit name! This way someone can specify path to file 
		unitFile,_ := cmd.Flags().GetString("unit-file")
		enabled,_ := cmd.Flags().GetBool("enabled")
		masked,_ := cmd.Flags().GetBool("mask")
		dropinsDir,_ := cmd.Flags().GetString("drop-ins")

		if !util.FileExists(unitFile) {
			return
		}
		if dropinsDir != "" {
			if !util.FileExists(dropinsDir) {
				return
			}
		}

		// Read in file contents
		unitFileContents, _ := ioutil.ReadFile(unitFile)

		// Build Systemd unit and add to FCC
		fccbUnit := systemd.NewUnit(unitFile, unitFileContents, enabled, masked)

		// Add drop-ins
		files, _ := ioutil.ReadDir(dropinsDir)
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".conf") {
				dropContents,_ := ioutil.ReadFile(path.Join(dropinsDir, file.Name()))
				fccbUnit.AddDropin(file.Name(), dropContents)
			}
		}

		// Insert Unit into Systemd section
		fccbUnit.Add(fcc)

		// Marshal and print it
		util.MarshalAndPrintFCC(fcc, inPlace)
	},
}

func init() {
	addCmd.AddCommand(unitCmd)

	unitCmd.Flags().StringP("unit-file", "f", "", "File containing systemd unit")
	unitCmd.Flags().Bool("mask", false, "Maks or unmask systemd unit")
	unitCmd.Flags().Bool("enabled", true, "Enable or disable systemd unit")
	unitCmd.Flags().StringP("drop-ins", "d", "", "Systemd drop-ins directory")

	unitCmd.MarkFlagRequired("unit-file")
}
