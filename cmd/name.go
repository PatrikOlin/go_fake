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
	"fmt"
	"bufio"			 
	"os"
	"os/exec"
	"math/rand"
	"time"
	"strings"
	"runtime"
	"log"			 
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("copy")
		fmt.Println(getFullName(fstatus))
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	nameCmd.Flags().BoolP("copy", "c", false, "Copy name to clipboard")
}

func getFullName(copyFlag bool) string {
	var fullName strings.Builder
	fullName.WriteString(getFirstName())
	fullName.WriteString(" ")
	fullName.WriteString(getSurname())

	if copyFlag == true {
		clipboard.WriteAll(fullName.String())
	}
	return fullName.String()
}

func getRandomLine(file *os.File) string {
					 
	scanner := bufio.NewScanner(file)
					 
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)

	lineNum := 1
	var pick string
	for scanner.Scan() {
		line := scanner.Text()

		roll := randGenerator.Intn(lineNum)
		if roll == 0 {
			pick = line
		}

		lineNum += 1
	}

	file.Close()	 
	return pick		 
	
}


func toClipboard(input []byte) {
	var copyCmd *exec.Cmd
	arch := runtime.GOOS

	// Mac OS
	if arch == "darwin" {
		copyCmd = exec.Command("pbcopy")
	}

	// Linux
	if arch == "linux" {
		copyCmd = exec.Command("xclip", "-selection", "c")
	}

	in, err := copyCmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := copyCmd.Start(); err != nil {
		log.Fatal(err)
	}

	if _, err := in.Write([]byte(input)); err != nil {
		log.Fatal(err)
	}

	copyCmd.Wait()

}

func getFirstName() string {
	file, err := os.Open("fname")
	check(err)		 
	var firstName string
	firstName = getRandomLine(file)

	return firstName
}

func getSurname() string {
	file, err := os.Open("lname")
	check(err)		 
	var surname string
	surname = getRandomLine(file)

	return surname
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
