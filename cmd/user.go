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

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Add user to FCC",
	Long: `TODO: Long desc with examples`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var uidP *int

		name,_ := cmd.Flags().GetString("name")
		homeDir,_ := GetStringOrNil(cmd, "home-dir")
		password,_ := GetStringOrNil(cmd, "password")
		uid,_ := cmd.Flags().GetInt("uid")
		gecos,_ := GetStringOrNil(cmd, "gecos")
		noCreateHome,_ := cmd.Flags().GetBool("no-create-home")
		primaryGroup,_ := GetStringOrNil(cmd, "primary-group")
		noUserGroup,_ := cmd.Flags().GetBool("no-user-group")
		noLogInit,_ := cmd.Flags().GetBool("no-log-init")
		shell,_ := GetStringOrNil(cmd, "shell")
		system,_ := cmd.Flags().GetBool("system")
		groups,_ := cmd.Flags().GetStringArray("groups")
		sshAuthKeys,_ := cmd.Flags().GetStringArray("ssh-authorized-key")

		if uid == -1 {
			uidP = nil
		} else {
			uidP = &uid
		}

		fccbUser := passwd.NewUser(
			name,
			password,
			uidP,
			gecos,
			homeDir,
			noCreateHome,
			primaryGroup,
			noUserGroup,
			noLogInit,
			shell,
			system)

		if len(groups) > 0 {
			for _,g := range groups {
				fccbUser.AddGroup(g)
			}
		}

		if len(sshAuthKeys) > 0 {
			for _,s := range sshAuthKeys {
				fccbUser.AddSSHAuthKey(s)
			}
		}

		fccbUser.Add(fcc)

		// Marshal and print it
		util.MarshalAndPrintFCC(fcc, inPlace)
	},
}

func init() {
	addCmd.AddCommand(userCmd)

	userCmd.Flags().StringP("name", "n", "", "User name")
	userCmd.Flags().StringP("home-dir", "d", "", "User home directory")
	userCmd.Flags().String("password", "", "Password hash for user")
	userCmd.Flags().IntP("uid", "u", -1, "UID for user")
	userCmd.Flags().String("gecos", "", "GECOS field for user")
	userCmd.Flags().BoolP("no-create-home", "M", false, "Do not create user home directory")
	userCmd.Flags().StringP("primary-group", "g", "", "Primary group for user")
	userCmd.Flags().BoolP("no-user-group", "N", false, "Do not create group with same name as user")
	userCmd.Flags().BoolP("no-log-init", "l", false, "Do not add user to lastlog and faillog db")
	userCmd.Flags().StringP("shell", "s", "", "Login shell for user")
	userCmd.Flags().BoolP("system", "r", false, "Create user as system account")
	userCmd.Flags().StringArrayP("groups", "G", []string{}, "Supplementary group for user. Can be specified multiple times.")
	userCmd.Flags().StringArrayP("ssh-authorized-key", "a", []string{}, "SSH authorized key for user. Can be specified multiple times.")

	userCmd.MarkFlagRequired("name")
}
