package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ZeroAC/rpc_match/kitex_gen/match_service"
	"github.com/ZeroAC/rpc_match/kitex_gen/match_service/match"
	"github.com/cloudwego/kitex/client"
)

// getInputUser parses args, If successful, the formatted user information is returned.
// where args[0] is the operation name and args[1] is the user ID, args[2] is user Name, args[3] is user Score
func getInputUser(args []string) (*match_service.User, error) {
	if args[0] == "add" {
		id, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			return nil, err
		}
		name := args[2]
		score, err := strconv.ParseInt(args[3], 10, 32)
		if err != nil {
			return nil, err
		}
		return &match_service.User{Id: int32(id), Name: name, Score: int32(score)}, nil
	}
	if args[0] == "remove" {
		id, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			return nil, err
		}
		return &match_service.User{Id: int32(id)}, nil
	}
	return &match_service.User{}, errors.New("operator name not found")
}

// argsCheck check whether the number of parameters of the operation is correct
func argsCheck(args []string) error {
	if args[0] == "add" && len(args) < 4 || args[0] == "remove" && len(args) < 2 {
		return errors.New("too few parameters")
	}
	if args[0] == "add" && len(args) > 4 || args[0] == "remove" && len(args) > 2 {
		return errors.New("too many parameters")
	}
	return nil
}
func main() {
	c, err := match.NewClient("match_service", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	bf := bufio.NewScanner(os.Stdin)
	for bf.Scan() {
		line := bf.Text()
		args := strings.Split(line, " ")
		if err := argsCheck(args); err != nil { //输入参数规范检查
			log.Println(err)
			continue
		}
		user, err := getInputUser(args) //获取输入的用户信息
		if err != nil {
			log.Println(err)
			continue
		}
		if args[0] == "add" { //rpc调用
			resp, err := c.AddUser(context.Background(), user, "")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("resp: %v\n", resp)
		} else if args[0] == "remove" {
			resp, err := c.RemoveUser(context.Background(), user, "")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("resp: %v\n", resp)
		}
	}
}
