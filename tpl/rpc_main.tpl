/**
* @file: main.go ==>
* @package: __
* @author:
* @since: 2022/12/18
* @desc: //TODO
 */

package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"{{ServerUpper}}Rpc/{{ServerLower}}"
)

var config = flag.String("f", "./{{ServerLower}}.yaml", "{{ServerLower}} rpc config.")

func main() {
	// read config...

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	server := grpc.NewServer()
	study.RegisterStudyServer(server, &StudyServer{})
	reflection.Register(server)

	if err = server.Serve(lis); err != nil {
		fmt.Println(err.Error())
		server.Stop()
		return
	}
	defer server.Stop()
}
