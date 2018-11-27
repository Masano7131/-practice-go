
package main

import (
	"log"
	"fmt"
	"os/exec"
)

func main() {
	log.Print("これはコマンドを実行するプログラム\n")
	command := exec.Command("sleep", "2").Output()

	if error != nil{
		log.Print("error: %s", error)
		return
	}

	log.Printf("コマンド開始")

	error = command.Wait()

	if error != nil{
		panic(fmt.Sprintf("エラー２: %v", error))
	}

	log.Print("コマンドが終了")
}