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

	"github.com/spf13/cobra"
	resty "gopkg.in/resty.v1"
)

// questionsCmd represents the questions command
var questionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "Are u ready to play our quiz?",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\nargs: %v", args)
		if len(args) == 0 {
			resp, err := resty.R().Get("http://localhost:8000/v1/quiz/questions")
			fmt.Printf("\nError: %v", err)
			fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
			fmt.Printf("\nResponse Body: %v", resp)
		} else {
			resp, err := resty.R().Get("http://localhost:8000/v1/quiz/questions?next=" + args[0])
			fmt.Printf("\nResponse Body: %v", resp)
			fmt.Printf("\nError: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(questionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// questionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// questionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
