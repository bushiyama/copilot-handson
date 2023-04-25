// じゃんけんゲームを作成する
// 0: グー, 1: チョキ, 2: パー
// あなたの手: 0
// あなた: グー, CPU: チョキ
// あなたの勝ち
// 0: グー, 1: チョキ, 2: パー
// あなたの手: 1
// あなた: チョキ, CPU: チョキ
// あいこ
// 0: グー, 1: チョキ, 2: パー
// あなたの手: 2
// あなた: パー, CPU: チョキ
// あなたの負け
// 0: グー, 1: チョキ, 2: パー
// あなたの手: 3
// 0, 1, 2のどれかを入力してください
//
// このようなじゃんけんゲームを作成してください。
// ただし、CPUの手は乱数で生成してください。
// また、ユーザーが0, 1, 2以外の数値を入力した場合は、
// 「0, 1, 2のどれかを入力してください」と表示してください。
//
// ヒント: 乱数を生成するには、math/randパッケージを使います。
// ヒント: 乱数のseedを設定するには、math/rand.Seed関数を使います。
// ヒント: 乱数のseedには、time.Now().UnixNano()を使います。
// ヒント: 乱数を生成するには、math/rand.Intn関数を使います。
// ヒント: ユーザーの入力を受け取るには、fmt.Scan関数を使います。
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ユーザーの入力を受け取る関数
func input() int {
	var input int
	fmt.Print("あなたの手: ")
	fmt.Scan(&input)
	return input
}

// CPUの手を生成する関数
func cpu() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}

// じゃんけんの結果を表示する関数
func result(user, cpu int) {
	hands := []string{"グー", "チョキ", "パー"}
	fmt.Printf("あなた: %s, CPU: %s\n", hands[user], hands[cpu])
}

func main() {
	for {
		fmt.Println("0: グー, 1: チョキ, 2: パー")
		user := input()
		if user < 0 || user > 2 {
			fmt.Println("0, 1, 2のどれかを入力してください")
			continue
		}

		cpu := cpu()
		result(user, cpu)

		switch {
		case user == cpu:
			fmt.Println("あいこ")
		case user == 0 && cpu == 1:
			fmt.Println("あなたの勝ち")
		case user == 1 && cpu == 2:
			fmt.Println("あなたの勝ち")
		case user == 2 && cpu == 0:
			fmt.Println("あなたの勝ち")
		default:
			fmt.Println("あなたの負け")
		}
	}
}
