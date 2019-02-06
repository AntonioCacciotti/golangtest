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
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	resty "gopkg.in/resty.v1"
)

type checkAnswer struct {
	QID int    `json:"questionID""`
	AID int    `json:"answerID"`
	NID string `json:"nicknameID"`
}

// answerCmd represents the answer command
var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "type your answer",
	RunE: func(cmd *cobra.Command, args []string) error {
		//copy(s, args)
		j := strings.Join(args, ",")
		//fmt.Println("string join", j)
		s := strings.Split(j, ",")
		fmt.Println("copy args to new slice:", s, "len:", len(s))
		if len(s) < 3 {
			return fmt.Errorf("Invalir argument! args format is questionID,answerID,nickname")
		}

		qID := convertStringToInt(s[0])
		answerID := convertStringToInt(s[1])
		nickanme := s[2]
		answer := checkAnswer{qID, answerID, nickanme}
		log.Println("answer q s n :", qID, answerID, nickanme)
		data, err := json.Marshal(answer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Json print %s\n", data)
		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(answer).
			Put("http://localhost:8000/v1/quiz/checkanswer")

		fmt.Printf("\nError: %v", err)
		fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
		fmt.Printf("\nResponse Body: %v", resp)
		return nil
	},
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		//Fatalln print and call os.Exit
		log.Fatalln("error converting from string to int", err)
	}
	return i
}
func init() {
	rootCmd.AddCommand(answerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// answerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// answerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
