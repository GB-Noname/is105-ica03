package main

import (
	"fmt"
	"strings"

	"./fileutils"
)

func main() {

	//hexPath, _ := filepath.Abs("/GitHub/org/is105-ica03/treasure/treasure.txt")
	//fmt.Printf("%s", fileutils.FileToByteslice(hexPath))
	//treasure.PrintTreasureUTF8(hexPath)
	byteSlice := fileutils.FileToByteslice("/GitHub/org/is105-ica03/treasure/treasure.txt")
	//fmt.Println(byteSlice)
	//fmt.Printf("%s", byteSlice)
	//fmt.Printf("%q", byteSlice)
	s := fmt.Sprintln(byteSlice)
	fmt.Println(s)
	strings.Trim(s, "\\x")
	fmt.Printf("%c", 120)

	//bs, _ := hex.DecodeString(string(byteSlice))
	//fmt.Println(bs)
	//treasure.PrintTreasureUTF8(hexPath)

}
