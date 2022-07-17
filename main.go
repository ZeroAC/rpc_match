package main

import (
	"fmt"
	"log"

	match_service "github.com/ZeroAC/rpc_match/kitex_gen/match_service/match"
)

func main() {
	go ConsumeTask()
	svr := match_service.NewServer(new(MatchImpl))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("loading...")

}
