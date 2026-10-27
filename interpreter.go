
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func runBF(code string) error {
	// 30000個のメモリを用意する
	memory := make([]byte, 30000)

	// 現在のメモリ位置
	ptr := 0

	// ループの対応関係を調べる
	jumps := make(map[int]int)
	stack := []int{}

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '[':
			stack = append(stack, i)

		case ']':
			if len(stack) == 0 {
				return fmt.Errorf("ループ終了 ] に対応する [ がありません")
			}

			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			jumps[start] = i
			jumps[i] = start
		}
	}

	if len(stack) != 0 {
		return fmt.Errorf("ループ開始 [ に対応する ] がありません")
	}

	input := bufio.NewReader(os.Stdin)

	// 命令を順番に実行する
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '+':
			memory[ptr]++

		case '-':
			memory[ptr]--

		case '>':
			ptr++
			if ptr >= len(memory) {
				return fmt.Errorf("メモリの右端を超えました")
			}

		case '<':
			if ptr == 0 {
				return fmt.Errorf("メモリの左端を超えました")
			}
			ptr--

		case '.':
			_, err := os.Stdout.Write([]byte{memory[ptr]})
			if err != nil {
				return err
			}

		case ',':
			value, err := input.ReadByte()

			if err == io.EOF {
				memory[ptr] = 0
			} else if err != nil {
				return err
			} else {
				memory[ptr] = value
			}

		case '[':
			if memory[ptr] == 0 {
				i = jumps[i]
			}

		case ']':
			if memory[ptr] != 0 {
				i = jumps[i]
			}
		}
	}

	return nil
}
