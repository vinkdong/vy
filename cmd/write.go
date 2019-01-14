// Copyright © 2019 VinkDong <dong@wenqi.us>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vinkdong/vy/pkg/vy"
	"github.com/pkg/errors"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			return err
		}
		if len(args) < 2 {
			return errors.New("must specify a tag and value")
		}
		w := vy.Write{
			Path: input,
			Tag: args[0],
			Value: args[1],
		}
		return w.Write()
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	writeCmd.Flags().StringP("input", "i", "", "Input file")
}
