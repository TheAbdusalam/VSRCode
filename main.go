package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"strconv"
)

var dimenstion int = 64 // INFO: should be the power of 2^n
var padding int = 0
var MESSAGE_TEXT string = "a Computer Engineer that loves low-level stuff, i pray everyday to not have to use JS (don’t get accepted mostly) For the last ~8 years I’ve been lucky enough to learn, work on and teach (for a brief period of time) stuff that i am truly passionate about. Here i post my thoughts on them, things i find interesting and most importantly to articulate my thoughts to understand myself better. You will find me bragging about stuff i did, occasionally."

var BIT_SIZE int = 8

type blockType [64][64]int // INFO: should be the same as the number of dimenstions

var block blockType

func main() {

	// Error handling

	encodingString := encodeString(MESSAGE_TEXT)
	if dimenstion*dimenstion < len(encodingString) {
		fmt.Println("Data will overflow when using,", dimenstion, "please use a higher dimenstion matrix or reduce the message size to fit", dimenstion*dimenstion)
		fmt.Println("Current message size:", len(encodingString))
		fmt.Println("Recommended dimenstion size:", calculateDimenstion(len(encodingString)))
		os.Exit(0)
	}

	newBlock := parsePadding(&block)

	col, row, str_num := 0, 0, 0
outerLoop:
	for row < dimenstion {
		col = 0
		for col < dimenstion {
			if str_num >= len(encodingString) {
				break outerLoop
			}

			num, _ := strconv.Atoi(string(encodingString[str_num]))
			newBlock[row][col] = int(num)
			col++
			str_num++
		}
		row++
	}

	draw(newBlock)

	fmt.Print("\n\n\n")
	// fmt.Printf("%s: %d/%d = %d%% used\n", encodingString, len(encodingString), dimenstion*dimenstion, (len(encodingString)*100)/(dimenstion*dimenstion)) // Debug the length and size of the binary message

	fmt.Println(decode(encodingString))

}

func calculateDimenstion(messageSize int) int {
	for {
		if dimenstion*dimenstion > messageSize {
			break
		}
		dimenstion *= 2
	}

	return dimenstion
}

func draw(b blockType) {
	clearScreen()
	display(b)
}

func display(b blockType) {
	for i := range b {
		for j := range b[i] {
			if b[i][j] == 1 {
				fg()
			} else {
				bg()
			}
		}
		fmt.Println()
	}
}

func fg() {
	fmt.Print("\033[48;5;120m  \033[0m")
}

func bg() {
	fmt.Print("\033[48;5;0m  \033[0m")
}

func clearScreen() {
	cmd, _ := exec.Command("clear").Output()
	fmt.Print(string(cmd))
}

func parsePadding(b *blockType) blockType {
	for i := range *b {
		for x := 0; x <= padding; x++ {
			for k := range b[x] {
				b[x][k] = 0
			}
			for k := range b[dimenstion-x] {
				b[dimenstion-(x+1)][k] = 0
			}
			b[i][x] = 0
			b[i][dimenstion-(x+1)] = 0
		}
	}

	return *b
}

func encodeString(data string) string {
	if BIT_SIZE == 8 {
		data = fmt.Sprintf("%08b", []byte(data))
	} else {
		data = fmt.Sprintf("%b", []byte(data))
	}

	data = strings.ReplaceAll(data, " ", "")
	data = strings.TrimPrefix(data, "[")
	data = strings.TrimSuffix(data, "]")
	return data
}

func decode(message string) string {
	written := strings.Split(message, "")
	data := ""
	for i := 0; i < len(written); i += BIT_SIZE {
		joinStr := strings.Join(written[i:i+BIT_SIZE], "")
		num, _ := strconv.ParseInt(joinStr, 2, 64)
		data += string(num)
	}

	return data
}
