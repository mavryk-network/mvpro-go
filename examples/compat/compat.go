package main

import (
	"context"
	"fmt"

	"github.com/mavryk-network/tzpro-go/tzpro"

	"github.com/echa/log"
)

func main() {
	log.SetLevel(log.LevelDebug)
	c := tzpro.NewClient("https://api.tzpro.io", nil).WithLogger(log.Log)
	_, err := c.Contract.NewQuery().Run(context.Background())
	fmt.Println(err)
}
