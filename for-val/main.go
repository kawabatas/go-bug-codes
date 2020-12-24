package main

import "fmt"

func main() {
	// forループ変数のバグ
	var printFuncList []func()
	for _, s := range []string{"a", "b", "c"} {
		// s := s // これが必要
		printFuncList = append(printFuncList, func() {
			fmt.Printf("s = \"%s\" (%p)\n", s, &s)
		})
	}
	for _, printFunc := range printFuncList {
		printFunc()
	}

	// deferの意図せぬ挙動
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // OK: 9 ... 0 を出力します
		// 	defer func() { fmt.Println(i) }()       // NG: "10" を10回出力します
		// 	defer func(i int) { fmt.Println(i) }(i) // OK: 9 ... 0 を出力します
		// 	defer print(&i)                         // NG: "10" を10回出力します
		// 	go fmt.Println(i)                       // OK: 0 ... 9 をランダムな順に出力します
		// 	go func() { fmt.Println(i) }()          // NG: 完全に予測不可能
	}
}
