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
	"github.com/spf13/cobra"
	"github.com/atotto/clipboard"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"fmt"
	"os"
)

// addressCmd represents the address command
var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("copy")
		fmt.Println(getFullAddress(fstatus))
	},
}

func init() {
	rootCmd.AddCommand(addressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addressCmd.Flags().BoolP("copy", "c", false, "Copy person to clipboard")
}

func getFullAddress(copyFlag bool) string {
	rand.Seed(time.Now().UnixNano())
	var fullAddress strings.Builder
	fullAddress.WriteString(getStreet())
	fullAddress.WriteString("\n")
	fullAddress.WriteString(getPostalCode())

	if copyFlag == true {
		clipboard.WriteAll(fullAddress.String())
	}
	return fullAddress.String()
}

func getStreetPrefix() string {
	file, err := os.Open("stpre")
	check(err)
	stPrefix := getRandomLine(file)

	return stPrefix
}

func getStreet() string {
	suffixes := []string{"vägen", "gatan", "gränden", "gränd", "stigen", "branten"}
	var street strings.Builder
	street.WriteString(getStreetPrefix())
	street.WriteString(suffixes[rand.Intn(len(suffixes))])
	street.WriteString(" ")
	street.WriteString(strconv.Itoa(rand.Intn(198) + 1))
	if rand.Intn(8) < 2 {
		street.WriteString(randomLetter())
	}

	return street.String()
}

func randomLetter() string {
	var letters = []rune("ABCDEFGHIJK")
	letter := make([]rune, 1)
	for i := range letter {
		letter[i] = letters[rand.Intn(len(letters))]
	}

	return string(letter)
}

func getPostalCode() string {
	file, err := os.Open("pnumort")
	check(err)
	postal := getRandomLine(file)

	return postal
}
