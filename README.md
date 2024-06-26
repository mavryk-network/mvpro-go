## MvPro-Go – Official Go SDK for the MvPro API

The official Blockwatch Go client library for MvPro. This SDK is free to use under a permissive license and works with the most recent version of the MvPro API v018-2024-01-15. API documentation can be found [here](https://docs.mvpro.io/).

This SDK is based on [MvGo](https://github.com/mavryk-network/mvgo), our open-source Go library for Mavryk.

Blockwatch maintains MvPro-Go on a regular basis to support changes to the Mavryk protocol. We add new API features as they are released. Open-source support is provided through issues in this Github repository at a best effor basis. For commercial support options, please contact us at support@blockwatch.cc.


### MvPro-Go Versioning

As long as MvPro-Go is in beta status we will use major version 0.x. Once interfaces are stable we switch to 1.x. The minor version number expresses compatibility with a Mavryk protocol release.

Supported API and Mavryk versions

- **v0.18**: API release v018-2024-01-15, Mavryk Atlas

Older API versions are no longer supported. Please use an older version tag.


### Installation

```sh
go get -u github.com/mavryk-network/mvpro-go
```

Then import, using

```go
import (
	"github.com/mavryk-network/mvpro-go/mvpro"
)
```

### Initializing the MvPro SDK Client

All functions are exported through a `Client` object. For convenience we have defined two default clients `mvpro.DefaultClient` for mainnet and `mvpro.IpfsClient`for our IPFS gateway. You may construct custom clients for different API URLs like so:

```go
c, err := mvpro.NewClient("https://api.mvpro.io", nil)
```

The default configuration should work just fine, but if you need special timeouts, proxy or TLS settings you may use a custom `http.Client` as second argument.

```go
import (
	"crypto/tls"
	"log"
	"net"
	"net/http"

	"github.com/mavryk-network/mvpro-go/mvpro"
)


func main() {
	hc := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   2 * time.Second,
				KeepAlive: 180 * time.Second,
			}).Dial,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			}
		}
	}

	c, err := mvpro.NewClient("https://my-private-index.local:8000", hc)
	if err != nil {
		log.Fatalln(err)
	}
}
```

### Reading a single Mavryk Account

```go
import (
  "context"
  "github.com/mavryk-network/mvpro-go/mvpro"
  "github.com/mavryk-network/mvgo/mavryk"
)

// use default Mainnet client
client := mvpro.DefaultClient
ctx := context.Background()
addr := mavryk.MustParseAddress("mv3CUdoCngwWGwK3i132dyHsffsmmgy5KZBn")

// get account data and embed metadata if available
params := mvpro.NewAccountParams().WithMeta()
a, err := client.GetAccount(ctx, addr, params)
```

### Reading Smart Contract Storage

```go
import (
  "context"
  "github.com/mavryk-network/mvgo/mavryk"
  "github.com/mavryk-network/mvpro-go/mvpro"
)

// use default Mainnet client
client := mvpro.DefaultClient
ctx := context.Background()
addr := mavryk.MustParseAddress("KT1Puc9St8wdNoGtLiD2WXaHbWU7styaxYhD")

// read storage from API
params := mvpro.NewContractParams()
raw, err := client.GetContractStorage(ctx, addr, params)

// access individual properties with correct type
tokenAddress, ok := raw.GetAddress("tokenAddress")
tokenPool, ok := raw.GetBig("tokenPool")
xtzPool, ok := raw.GetBig("xtzPool")
```

### Listing Account Transactions

```go
import (
  "context"
  "github.com/mavryk-network/mvpro-go/mvpro"
  "github.com/mavryk-network/mvgo/mavryk"
)

// use default Mainnet client
client := mvpro.DefaultClient
ctx := context.Background()
addr := mavryk.MustParseAddress("mv1E2Y8khTrfaRUeErWUBfg6G7zNMKnM4JJL")

// list operations sent and received by this account
params := mvpro.NewOpParams().WithLimit(100).WithOrder(mvpro.OrderDesc)
ops, err := client.GetAccountOps(ctx, addr, params)
```

### Cursoring through results

The SDK has a convenient way for fetching results longer than the default maximum of 500 entries. We use the more efficient cursor method here, but offset would work similar. An empty result means there is no more data available right now. As the chain grows you can obtain fresh data by using the most recent row_id like shown below

```go
import (
  "context"
  "github.com/mavryk-network/mvpro-go/mvpro"
  "github.com/mavryk-network/mvgo/mavryk"
)

client := mvpro.DefaultClient
ctx := context.Background()
addr := mavryk.MustParseAddress("mv1E2Y8khTrfaRUeErWUBfg6G7zNMKnM4JJL")
params := mvpro.NewOpParams()

for {
	// fetch next batch from the explorer API
	ops, err := client.GetAccountOps(ctx, addr, params)
	// handle error if necessary

	// stop when result is empty
	if len(ops) == 0 {
		break
	}

	// handle ops here

	// prepare for next iteration
	params = params.WithCursor(ops[len(ops)-1].RowId)
}
```

### Decoding smart contract data into Go types

```go
import (
  "context"
  "github.com/mavryk-network/mvgo/mavryk"
  "github.com/mavryk-network/mvpro-go/mvpro"
)

// use default Mainnet client
client := mvpro.DefaultClient
ctx := context.Background()
addr := mavryk.MustParseAddress("KT1Puc9St8wdNoGtLiD2WXaHbWU7styaxYhD")

// read storage from API
raw, err := client.GetContractStorage(ctx, addr, mvpro.NewContractParams())

// decode into Go struct
type DexterStorage struct {
	Accounts                int64         `json:"accounts"`
	SelfIsUpdatingTokenPool bool          `json:"selfIsUpdatingTokenPool"`
	FreezeBaker             bool          `json:"freezeBaker"`
	LqtTotal                *big.Int      `json:"lqtTotal"`
	Manager                 mavryk.Address `json:"manager"`
	TokenAddress            mavryk.Address `json:"tokenAddress"`
	TokenPool               *big.Int      `json:"tokenPool"`
	XtzPool                 *big.Int      `json:"xtzPool"`
}

dexterPool := &DexterStorage{}
err := raw.Unmarshal(dexterPool)
```

### Listing bigmap key/value pairs with server-side data unfolding

```go
import (
  "context"
  "github.com/mavryk-network/mvpro-go/mvpro"
)

type HicNFT struct {
	TokenId   int               `json:"token_id,string"`
	TokenInfo map[string]string `json:"token_info"`
}

client := mvpro.DefaultClient
ctx := context.Background()
params := mvpro.NewContractParams().
	WithUnpack().
	WithLimit(500)

for {
	// fetch next batch from the explorer API
	nfts, err := client.GetBigmapValues(ctx, 514, params)
	if err != nil {
		return err
	}

	// stop when result is empty
	if len(nfts) == 0 {
		break
	}
	for _, v := range nfts {
		var nft HicNFT
		if err := v.Unmarshal(&nft); err != nil {
			return err
		}
		// handle the value
	}
}
```

### Building complex Table Queries

The [MvPro Table API](https://mvpro.com/docs/api#table-endpoints) is the fastest way to ingest and process on-chain data in bulk. The SDK defines typed query objects for most tables and allows you to add filter conditions and other configuration to these queries.

```go
import (
  "context"
  "github.com/mavryk-network/mvpro-go/mvpro"
)

client := mvpro.DefaultClient
ctx := context.Background()

// create a new query object
q := client.NewBigmapValueQuery()

// add filters and configure the query to list all active keys
q.WithFilter(mvpro.FilterModeEqual, "bigmap_id", 514).
    WithColumns("row_id", "key_hash", "key", "value").
    WithLimit(1000).
    WithOrder(mvpro.OrderDesc)

// execute the query
list, err := q.Run(ctx)

// walk rows
for _, row := range list.Rows {
	// process data here
}
```

### Listing many Bigmap keys with client-side data decoding

Extending the example above, we now use MvGo's Micheline features to decode annotated bigmap data into native Go structs. For efficiency reasons the API only sends binary (hex-encoded) content for smart contract storage. The SDK lets you decodes this into native Micheline primitives or native Go structs for further processing as shown in the example below.

```go
import (
  "context"
  "github.com/mavryk-network/mvpro-go/mvpro"
)

type HicNFT struct {
	TokenId   int               `json:"token_id,string"`
	TokenInfo map[string]string `json:"token_info"`
}

client := mvpro.DefaultClient
ctx := context.Background()

// fetch bigmap info with prim (required for key/value types)
params := mvpro.NewContractParams().WithPrim()
info, err := client.GetBigmap(ctx, 514, params))
keyType := info.MakeKeyType()
valType := info.MakeValueType()

// create a new query object
q := client.NewBigmapValueQuery()

// add filters and configure the query to list all active keys
q.WithFilter(mvpro.FilterModeEqual, "bigmap_id", 514).
    WithColumns("row_id", "key_hash", "key", "value").
    WithLimit(1000).
    WithOrder(mvpro.OrderDesc)

// execute the query
list, err := q.Run(ctx)

// walk rows
for _, row := range list.Rows {
    // unpack to Micheline primitives and tag types (we ignore errors here)
    key, _ := row.DecodeKey(keyType)
    val, _ := row.DecodeValue(valType)

    // unpack into Go type
    var nft HicNFT
    err = val.Unmarshal(&nft)

    // or access individual named values
    i, ok := val.GetInt64("token_id")
}
```

### Gracefully handle rate-limits

To avoid excessive overload of our API we limit the rate at which we process your requests. This means your program may from time to time run into a rate limit. To let you gracefully handle retries by waiting until a rate limit resets, we expose the deadline and a done channel much like Go's network context does. Here's how you may use this feature:

```go
var acc *mvpro.Account
for {
	var err error
	acc, err = mvpro.GetAccount(ctx, mavryk.MustParseAddress("mv1E2Y8khTrfaRUeErWUBfg6G7zNMKnM4JJL"))
	if err != nil {
		if e, ok := mvpro.IsRateLimited(err); ok {
			fmt.Printf("Rate limited, waiting for %s\n", e.Deadline())
			select {
			case <-ctx.Done():
				// wait until external context is canceled
				err = ctx.Err()
			case <-e.Done():
				// wait until rate limit reset and retry
				continue
			}
		}
	}
	break
}

// handle error and/or result here

```

## License

The MIT License (MIT) Copyright (c) 2021-2024 Blockwatch Data Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is furnished
to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
