package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var i int

func main() {
	var flname string
	fmt.Print("ファイル名の入力：")
	fmt.Scan(&flname)
	flopen(flname)
}

func flopen(flname string) {
	flop, err := os.OpenFile("tab.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer flop.Close()

	scan := bufio.NewScanner(flop)
	linebox := make([]string, 0, 100)

	for scan.Scan() {
		linebox = append(linebox, scan.Text())
	}
	if serr := scan.Err(); serr != nil {
		log.Fatal(serr)
	}

	flw, err2 := os.OpenFile("tabspace.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err2 != nil {
		log.Fatal(err2)
	}

	defer flw.Close()

	var spacecount int
	for index := 0; index < len(linebox); index++ {
		newlib := linebox[index]
		for i := 0; i < len(string(newlib)); i++ {
			if newlib[i] != 9 && newlib[i] != 32 { //文字
				flw.WriteString(string(newlib[i]))
				spacecount++
			} else if newlib[i] == 9 && spacecount == 0 { //最初の所がTABキーの場合
				continue
			} else if newlib[i] == 9 && spacecount%8 > 0 { //最初の所がTABキーじゃない場合
				flw.WriteString(string(0x09))
				spacecount = 0
			} else if newlib[i] == 32 && spacecount == 0 {
				flw.WriteString(" ")
				spacecount++
			} else {
				flw.WriteString(string(0x09))
				spacecount = 0
			}
		}
		spacecount = 0
		flw.WriteString("\n")
	}
	fmt.Println("処理終わりです。")
}
