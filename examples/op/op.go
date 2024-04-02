package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mavryk-network/mvpro-go/mvpro"
)

func main() {
	// use default Mainnet client
	c := mvpro.NewClient("https://api.mvpro.io", nil)
	ctx := context.Background()

	q := c.Op.NewQuery().AndEqual("hash", os.Args[1])
	res, err := q.Run(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		buf, _ := json.MarshalIndent(res.Rows()[0], "", "  ")
		fmt.Println(string(buf))
	}
}
