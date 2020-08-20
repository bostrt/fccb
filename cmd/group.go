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
	"github.com/bostrt/fccb/pkg/passwd"
	"github.com/bostrt/fccb/pkg/util"
	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Add group to FCC",
	Long: `TODO: Long desc with examples`,
	Run: func(cmd *cobra.Command, args []string) {
		name,_ := cmd.Flags().GetString("name")
		gid,_ := cmd.Flags().GetInt("gid")
		password,_ := cmd.Flags().GetString("password")
		system,_ := cmd.Flags().GetBool("system")

		fccbGroup := passwd.NewGroup(
			name,
			&gid,
			&password,
			&system)

		fccbGroup.Add(fcc)

		// Marshal and print it
		// TODO: Have option for STDOUT vs overwrite file
		util.MarshalAndPrintFCC(fcc, inPlace)
	},
}

func init() {
	addCmd.AddCommand(groupCmd)

	groupCmd.Flags().StringP("name", "n", "", "Group name")
	groupCmd.Flags().IntP("gid", "g", -1, "GID for group")
	groupCmd.Flags().String("password", "", "Password hash for group")
	groupCmd.Flags().BoolP("system", "r", false, "Create a system group")

	groupCmd.MarkFlagRequired("name")
}
