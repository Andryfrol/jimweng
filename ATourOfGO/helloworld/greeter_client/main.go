/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/goPractice/ATourOfGO/helloworld/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var Input_flag_num = flag.Int("NumOut", 10, "Input number")

const (
	address           = "localhost:50051"
	defaultName       = "world"
	defaultNum  int32 = 123
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	number := defaultNum
	if len(os.Args) > 1 {
		name = os.Args[1]
		number = int32(*Input_flag_num)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &helloworld.HelloRequest{
		Name:    name,
		TestNum: number,
		// convert int32 into helloworld.HelloRequest_EnumTest
		EnumTest: helloworld.HelloRequest_EnumTest(number),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	r, err = c.SayHelloAgain(context.Background(), &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
