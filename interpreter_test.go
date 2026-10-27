
package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "hello_world.igfa")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("実行に失敗: %v\n%s", err, output)
	}

	result := strings.TrimSpace(string(output))

	if result != "Hello World!" {
		t.Errorf("期待値: Hello World!, 実際: %s", result)
	}
}

func TestEcho(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "echo.igfa")

	cmd.Stdin = strings.NewReader("Z")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("実行に失敗: %v\n%s", err, output)
	}

	if string(output) != "Z" {
		t.Errorf("期待値: Z, 実際: %q", string(output))
	}
}


func TestPointerUnderflow(t *testing.T) {
	err := runBF("<")

	if err == nil {
		t.Fatal("メモリの左端を超えたのにエラーが発生しませんでした")
	}

	if !strings.Contains(err.Error(), "左端") {
		t.Errorf("予期しないエラー: %v", err)
	}
}


func TestUnmatchedBrackets(t *testing.T) {
	tests := []struct {
		name string
		code string
	}{
		{"ループ開始だけ", "["},
		{"ループ終了だけ", "]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := runBF(tt.code)

			if err == nil {
				t.Fatal("不正なループなのにエラーが発生しませんでした")
			}

			if !strings.Contains(err.Error(), "対応する") {
				t.Errorf("予期しないエラー: %v", err)
			}
		})
	}
}


func TestPointerOverflow(t *testing.T) {
    code := strings.Repeat(">", 30000)

    err := runBF(code)

    if err == nil {
        t.Fatal("メモリの右端を超えたのにエラーが発生しませんでした")
    }

    if !strings.Contains(err.Error(), "右端") {
        t.Errorf("予期しないエラー: %v", err)
    }
}


func TestInvalidInstruction(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "invalid.igfa")

	output, err := cmd.CombinedOutput()

	if err == nil {
		t.Fatal("不正な命令なのに正常終了しました")
	}

	result := string(output)

	if !strings.Contains(result, "3行目") {
		t.Errorf("エラーの行番号が正しくありません: %s", result)
	}

	if !strings.Contains(result, "不明な命令") {
		t.Errorf("不正な命令のエラーが表示されません: %s", result)
	}
}
