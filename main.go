package main

import (
	"os"
	"os/exec"
	"strings"
)

const dictBaseURL = "https://dictionary.cambridge.org/zht/詞典/英語-漢語-繁體/"

func main() {
	args := os.Args[1:]
	word := strings.Join(args, "-")
	exec.Command("open", dictBaseURL+word).Run()
}
