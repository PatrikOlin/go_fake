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
)

// pinCmd represents the pin command
var pinCmd = &cobra.Command{
	Use:   "pin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("copy")
		fmt.Println(getPIN(fstatus))
	},
}

func init() {
	rootCmd.AddCommand(pinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	pinCmd.Flags().BoolP("copy", "c", false, "Copy name to clipboard")
}

func getPIN(copyFlag bool) string {
	rand.Seed(time.Now().UnixNano())
	pin := toString(complete(partial()))

	if copyFlag == true {
		clipboard.WriteAll(pin)
	}
	return pin
}

func randInts(min, max int) (int, int, int) {
	num := min + rand.Intn(max-min)
	if num > 99 {
		return num / 100 % 10, num / 10 % 10, num % 10
	}
	if num < 10 {
		return 0, num, 0
	}
	return num / 10 % 10, num % 10, 0
}

func partial() []int {
	y, yy, _ := randInts(0, 100)
	m, mm, _ := randInts(1, 13)
	d, dd, _ := randInts(1, 29)
	x, xx, xxx := randInts(111, 1000)
	// x, xx, xxx := randInts(980, 1000) // safe range

	return []int{y, yy, m, mm, d, dd, x, xx, xxx}
}

func complete(ints []int) []int {
	var sum int
	for i, v := range ints {
		v = v * (2 - i%2)
		if v >= 10 {
			v -= 9
		}
		sum += v
	}
	lastDigit := (100 - sum) % 10
	return append(ints, lastDigit)
}

func toString(a []int) string {
	b := make([]string, len(a)+1)

	for i, v := range a {
		if i == 6 {
			b = append(b, "-")
		}
		b = append(b, strconv.Itoa(v))
	}
	return strings.Join(b, "")
}
