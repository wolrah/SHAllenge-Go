package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log/slog"
	"strconv"
)

const username string = "wolrah"
const data string = "R93900X+1thread+LearningGo-"
const start int = 48000000000
const target int = 100000000000

var lowest [32]byte = [32]byte{0x00, 0x00, 0x00, 0x00, 0x0b, 0x49, 0x0d, 0x34, 0x24, 0x87, 0x79, 0xfa, 0x66, 0x88, 0x7f, 0xf0, 0x52, 0xc3, 0xe7, 0x54, 0x90, 0xb2, 0x69, 0xf3, 0xf9, 0x36, 0x7b, 0x13, 0x94, 0x4c, 0xf7, 0x73}

//var lowest [32]byte = [32]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

func main() {
	//slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("Starting search...", "sequence", start, "target", target, "lowest", fmt.Sprintf("%x", lowest))
	lowest_slice := lowest[:]
	for sequence := start; sequence <= target; sequence++ {
		prefix := username + "/" + data
		entry := prefix + strconv.Itoa(sequence)
		sum := sha256.Sum256([]byte(entry))
		sum_slice := sum[:]
		if bytes.Compare(lowest_slice, sum_slice) > 0 {
			fmt.Printf("Winner!: %s %x\n", entry, sum)
			slog.Info("New Lowest Found", "sequence", sequence, "string", entry, "hash", sum)
			lowest = sum
			lowest_slice = lowest[:]
		}
		if sequence%1000000000 == 0 && sequence != start {
			slog.Info("1B Strings Tested", "sequence", sequence, "target", target, "lowest", fmt.Sprintf("%x", lowest), "lastentry", entry, "lasthash", fmt.Sprintf("%x", sum))
		}
	}
}
