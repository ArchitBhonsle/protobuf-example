package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/ArchitBhonsle/protobuf-example/pinger/tick"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	ticker := time.Tick(time.Second * 10)
	for time := range ticker {
		data := rand.Int31()
		t := tick.Tick{
			Time: timestamppb.New(time),
			Data: data,
		}

		out, err := proto.Marshal(&t)
		if err != nil {
			panic(err)
		}

		http.Post("http://localhost:8080", "text/plain", bytes.NewReader(out))

		fmt.Printf("ping %v at %v\n", data, time.Format("01/02/2006, 3:04:05 PM"))
	}
}
