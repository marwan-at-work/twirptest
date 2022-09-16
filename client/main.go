package main

import (
	"context"
	"fmt"
	"net/http"

	"marwan.io/twirptest/protoold"
)

func main() {
	c := protoold.NewHaberdasherProtobufClient("http://localhost:5959", http.DefaultClient)
	resp, err := c.MakeHat(context.Background(), &protoold.Size{
		Inches: 456,
	})
	must(err)
	fmt.Println("GOT", resp.Color, resp.Name, resp.Size)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
