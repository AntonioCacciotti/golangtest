// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	resty "gopkg.in/resty.v1"
)

// actionsCmd represents the actions command
var actionsCmd = &cobra.Command{
	Use:   "actions",
	Short: "possible actions are result or score",

	Run: func(cmd *cobra.Command, args []string) {
		action := strings.Join(args, "")
		if action == "result" {
			resp, err := resty.R().Get("http://localhost:8000/v1/quiz/end?action=result")
			fmt.Printf("\nResponse Status Code: %v", err)
			fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
			fmt.Printf("\nResponse Body: %v", resp)
		}
		if action == "scores" {
			resp, err := resty.R().Get("http://localhost:8000/v1/quiz/end?action=score")
			fmt.Printf("\nResponse Status Code: %v", err)
			fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
			fmt.Printf("\nResponse Body: %v", resp)
		}
	},
}

func init() {
	rootCmd.AddCommand(actionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// actionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// actionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
