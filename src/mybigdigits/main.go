// Copyright © 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var barType bool

func init() {
	const (
		defaultBar = false
		usage      = "draw an underbar and an overbar"
	)

	flag.BoolVar(&barType, "bar", defaultBar, usage)
	flag.BoolVar(&barType, "b", defaultBar, usage+"(shorthand)")
}

func main() {

	// if len(os.Args) == 1 {
	// 	fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
	// 	os.Exit(1)
	// }
	flag.Parse()
	fmt.Println("flag: ", barType)
	// for i := range os.Args {
	// 	fmt.Println("flg: ", os.Args[i])
	// }
	var stringOfDigits string
	if barType && len(os.Args) > 2 {
		stringOfDigits = os.Args[2]
	} else if len(os.Args) == 2 {
		stringOfDigits = os.Args[1]
	}

	for row := range bigDigits[0] {
		line := ""
		//fmt.Println("row: ", row)
		for column := range stringOfDigits {
			//fmt.Printf("string of digit: %T, %v\n", stringOfDigits, stringOfDigits)
			digit := stringOfDigits[column] - '0'
			if digit >= 0 && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number")
			}
			//fmt.Println("line after column: \n", line, "\n")
		}
		fmt.Println(line)
	}
}

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{"  1 ",
		" 11 ",
		"1 1 ",
		"  1 ",
		"  1 ",
		"  1 ",
		"1111"},
	{" 222 ",
		"2   2",
		"   2 ",
		"  2  ",
		" 2   ",
		"2    ",
		"22222"},
	{" 333 ",
		"3   3",
		"    3",
		"  33 ",
		"    3",
		"3   3",
		" 333 "},
	{"   4  ",
		"  44  ",
		" 4 4  ",
		"4  4  ",
		"444444",
		"   4  ",
		"   4  "},
	{"55555",
		"5    ",
		"5    ",
		" 555 ",
		"    5",
		"5   5",
		" 555 "},
	{" 666 ",
		"6    ",
		"6    ",
		"6666 ",
		"6   6",
		"6   6",
		" 666 "},
	{"77777",
		"    7",
		"   7 ",
		"  7  ",
		" 7   ",
		"7    ",
		"7    "},
	{" 888 ",
		"8   8",
		"8   8",
		" 888 ",
		"8   8",
		"8   8",
		" 888 "},
	{" 9999",
		"9   9",
		"9   9",
		" 9999",
		"    9",
		"    9",
		" 999 "},
}
