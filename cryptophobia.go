package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func binary(g string) string { // conversion to binary
	var s1 int
	for i := 0; i < len(g); i++ {
		if int(g[i]) != 1 {
			s1 = 2 * int(g[i])
		} else {
			s1 = 1
		}
	}
	res := (strconv.FormatInt(int64(s1), 2))
	return res
}

func hex(g string) string { // conversion to hexadecimal
	i, _ := strconv.ParseInt(g, 10, 64)
	return strings.ToUpper(fmt.Sprint(strconv.FormatInt(int64(i), 16)))
}

func convchar(z byte) string { // conversion input to numbers with the ascii
	var w int
	s := []byte{97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122}
	s1 := []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	for h := 0; h < 26; h++ {
		if z == s[h] || z == s[h]-32 {
			w = h + 1
		}
		for h := 0; h <= 9; h++ {
			if z == s1[h] {
				w = h + 26 // +26 for the diffrence between 123 like in list of (abc) and 123 like ('123') in other list of numbers because they look the same in binary
				//F DECRYPTAGE if w <26 then a number else it is a char
			}
		}
	}
	return fmt.Sprint(w)

}

func funds() string {

	xstr := ""
	x := os.Args[1:] // take input from the terminal
	//step1 := ""
	for i := 0; i < len(x); i++ {
		xstr = xstr + x[i] + " "
	}
	tab := make([]string, len(xstr))
	for i := 0; i < len(xstr)-1; i++ {

		//step1 = step1 +  convchar(xstr[i]) + "\n"
		tab[i] = convchar(xstr[i])

	}
	i := 0
	res := ""
	for i < len(xstr)-1 {
		if i == 0 {
			res = res + fmt.Sprint(hex(binary(tab[i])))
			i++
		} else if i%2 == 0 {
			res = res + fmt.Sprint("!", hex(binary(tab[i]))) //put ! between every char from input if the position pair
			i++
		} else {
			res = res + fmt.Sprint("?", hex(binary(tab[i])))
			i++
		}
		if len(res) == 0 {
			res = res + fmt.Sprint("?")
		}

	}
	return res
}
func Crypted_Result() {
	file, err := os.OpenFile("Crypted_Result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// os.O_WRONLY tells the computer you are only going to writo to the file, not read
	// os.O_CREATE tells the computer to create the file if it doesn't exist
	// os.O_APPEND tells the computer to append to the end of the file instead of overwritting or truncating it

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)
	tta := funds()
	for _, data := range tta {
		_, _ = datawriter.WriteString(string(data))
	}

	datawriter.Flush()
	file.Close()

}

func main() {

	Crypted_Result()
}
