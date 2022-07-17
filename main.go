package main

import (
	match_service "github.com/ZeroAC/rpc_match/kitex_gen/match_service/match"
	"log"
)

func main() {
	svr := match_service.NewServer(new(MatchImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
