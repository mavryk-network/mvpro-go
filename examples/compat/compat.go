package main

import (
	"context"
	"fmt"

	"github.com/echa/log"
	"github.com/mavryk-network/mvpro-go/mvpro"
)

func main() {
	log.SetLevel(log.LevelDebug)
	c := mvpro.NewClient("https://api.mvpro.io", nil).WithLogger(log.Log)
	_, err := c.Contract.NewQuery().Run(context.Background())
	fmt.Println(err)
}
