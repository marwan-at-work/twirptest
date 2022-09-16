package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"marwan.io/twirptest/protonew"
)

func main() {
	log.Fatal(http.ListenAndServe(":5959", protonew.NewHaberdasherServer(&server{})))
}

type server struct {
}

// MakeHat implements protonew.Haberdasher
func (*server) MakeHat(ctx context.Context, req *protonew.Size) (*protonew.Hat, error) {
	fmt.Println("GOT", req.Inches)
	return &protonew.Hat{
		Size:  123,
		Color: "abc",
		Name:  "qwer",
	}, nil
}
