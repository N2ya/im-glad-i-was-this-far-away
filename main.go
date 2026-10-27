package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	commands := map[string]string{
		"よかった":   "+",
		"お祭り男！":  "-",
		"アカーン":   ">",
		"この距離で":  "<",
		"兄さん":    ".",
		"最後の餃子":  ",",
		"クイヤ":    "[",
		"浜田の嫁":   "]",
	}

	if len(os.Args) != 2 {
		fmt.Println("使い方: go run . ファイル名.igfa")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("ファイルを開けません:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result strings.Builder

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		
command, ok := commands[line]
if !ok {
    fmt.Fprintf(os.Stderr, "%d行目: 不明な命令 %q\n", lineNumber, line)
    os.Exit(1)
}


		result.WriteString(command)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("読み込みエラー:", err)
		return
	}

	
if err := runBF(result.String()); err != nil {
    fmt.Fprintln(os.Stderr, "実行エラー:", err)
    os.Exit(1)
}

}
