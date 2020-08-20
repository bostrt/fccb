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
	"github.com/bostrt/fccb/pkg/storage"
	"github.com/bostrt/fccb/pkg/util"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Add file to FCC",
	Args: cobra.ExactArgs(1),
	Long: `TODO: Longer desc with examples`,
	Run: func(cmd *cobra.Command, args []string) {

		path,_ := cmd.Flags().GetString("path")
		overwrite,_ := cmd.Flags().GetBool("overwrite")
		file,_ := cmd.Flags().GetString("file")
		append,_ := cmd.Flags().GetBool("append")
		mode,_ := cmd.Flags().GetInt("mode")
		owner := GetStringOrNilNoDefault(cmd, "owner")
		group := GetStringOrNilNoDefault(cmd, "group")

		if !util.FileExists(file) {
			return
		}

		fileContents, _ := ioutil.ReadFile(file)

		f := storage.NewFile(
			path,
			overwrite,
			fileContents,
			append,
			mode,
			owner,
			group,
			)

		f.Add(fcc)

		// Marshal and print it
		util.MarshalAndPrintFCC(fcc, inPlace)
	},
}

func init() {
	addCmd.AddCommand(fileCmd)

	fileCmd.Flags().StringP("path", "p", "", "The absolute path to file to add via FCC and Ignition")
	fileCmd.Flags().Bool("overwrite", false, "Whether to overwrite existing file")
	fileCmd.Flags().StringP("file", "f", "", "Local path to file with contents to add into FCC")
	fileCmd.Flags().Bool("append", false, "Whether to append file contents to file")
	fileCmd.Flags().Int("mode", 644, "File permissions mode")
	fileCmd.Flags().String("owner", "", "Username or UID for owner")
	fileCmd.Flags().String("group", "", "Group name or GID for owner")

	fileCmd.MarkFlagRequired("path")
	fileCmd.MarkFlagRequired("file")
}
