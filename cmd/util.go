package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func GetStringOrNil(cmd *cobra.Command, name string) (*string, bool) {
	changed := false
	if cmd.Flags().Changed(name) {
		changed = true
	}
	v,err := cmd.Flags().GetString(name)
	if err != nil {
		fmt.Println(err)
		return nil, changed
	}

	return &v, changed
}

func GetStringOrNilNoDefault(cmd *cobra.Command, name string) *string {
	v, changed := GetStringOrNil(cmd, name)
	if !changed {
		return nil
	}
	return v
}

// GetIntOrNil returns an *int if command line argument was passed
// or nil if no argument was passed. A second return value of type
// bool indicates whether the value was specified at CLI or not.
func GetIntOrNil(cmd *cobra.Command, name string) (*int, bool) {
	changed := false
	if 	cmd.Flags().Changed(name) {
		changed = true
	}
	cmd.Flags().Lookup("asdf")
	v, err := cmd.Flags().GetInt(name)
	if err != nil {
		fmt.Println(err)
		return nil, changed
	}
	return &v, changed
}

func GetIntOrNilNoDefault(cmd *cobra.Command, name string) *int {
	v, changed := GetIntOrNil(cmd, name)
	if !changed {
		return nil
	}
	return v
}

func Leave() {
	os.Exit(1)
}

