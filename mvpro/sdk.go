// Copyright (c) 2020-2024 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package mvpro

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/mavryk-network/mvpro-go/internal/client"
	"github.com/mavryk-network/mvpro-go/mvpro/defi"
	"github.com/mavryk-network/mvpro-go/mvpro/identity"
	"github.com/mavryk-network/mvpro-go/mvpro/index"
	"github.com/mavryk-network/mvpro-go/mvpro/ipfs"
	"github.com/mavryk-network/mvpro-go/mvpro/market"
	"github.com/mavryk-network/mvpro-go/mvpro/nft"
	"github.com/mavryk-network/mvpro-go/mvpro/token"
	"github.com/mavryk-network/mvpro-go/mvpro/wallet"

	// "github.com/mavryk-network/mvpro-go/mvpro/zmq"
	"github.com/echa/log"
	lru "github.com/hashicorp/golang-lru/v2"
)

var (
	SdkVersion    = "0.18.0"
	DefaultClient = NewClient("https://api.mvpro.io", nil)
)

type Client struct {
	Account  index.AccountAPI
	Block    index.BlockAPI
	Baker    index.BakerAPI
	Contract index.ContractAPI
	Explorer index.ExplorerAPI
	Metadata index.MetadataAPI
	Op       index.OpAPI
	Stats    index.StatsAPI
	Dex      defi.DexAPI
	Farm     defi.FarmAPI
	Lend     defi.LendingAPI
	Nft      nft.NftAPI
	Token    token.TokenAPI
	Domain   identity.DomainAPI
	Profile  identity.ProfileAPI
	Wallet   wallet.WalletAPI
	Market   market.MarketAPI
	Ipfs     ipfs.IpfsAPI
	// Zmq      zmq.ZmqAPI

	client *client.Client
}

func NewClient(url string, httpClient *http.Client) *Client {
	c := client.NewClient(url, httpClient).
		WithApiKey(os.Getenv("MVPRO_API_KEY")).
		WithUserAgent("mvpro-go/v" + SdkVersion)

	return &Client{
		Account:  index.NewAccountAPI(c),
		Block:    index.NewBlockAPI(c),
		Baker:    index.NewBakerAPI(c),
		Contract: index.NewContractAPI(c),
		Explorer: index.NewExplorerAPI(c),
		Metadata: index.NewMetadataAPI(c),
		Op:       index.NewOpAPI(c),
		Stats:    index.NewStatsAPI(c),
		Dex:      defi.NewDexAPI(c),
		Farm:     defi.NewFarmAPI(c),
		Lend:     defi.NewLendingAPI(c),
		Nft:      nft.NewNftAPI(c),
		Token:    token.NewTokenAPI(c),
		Domain:   identity.NewDomainAPI(c),
		Profile:  identity.NewProfileAPI(c),
		Wallet:   wallet.NewWalletAPI(c),
		Market:   market.NewMarketAPI(c),
		Ipfs: ipfs.NewIpfsAPI(
			client.NewClient("https://ipfs.mvpro.io", httpClient).
				WithApiKey(os.Getenv("MVPRO_API_KEY")).
				WithUserAgent("mvpro-go/v" + SdkVersion).
				WithTimeout(60 * time.Second),
		),
		// Zmq:    zmq.NewZmqAPI(c),
		client: c,
	}
}

func (s *Client) WithHeader(key, value string) *Client {
	s.client.WithHeader(key, value)
	return s
}

func (s *Client) WithUserAgent(agent string) *Client {
	s.client.WithUserAgent(agent)
	return s
}

func (s *Client) WithApiKey(key string) *Client {
	s.client.WithApiKey(key)
	return s
}

func (s *Client) WithMarketUrl(url string) *Client {
	c := client.NewClient(url, nil).
		WithApiKey(os.Getenv("MVPRO_API_KEY")).
		WithUserAgent("mvpro-go/v" + SdkVersion)
	s.Market = market.NewMarketAPI(c)
	return s
}

func (s *Client) WithIpfsUrl(url string) *Client {
	c := client.NewClient(url, nil).
		WithApiKey(os.Getenv("MVPRO_API_KEY")).
		WithUserAgent("mvpro-go/v" + SdkVersion).
		WithTimeout(60 * time.Second)
	s.Ipfs = ipfs.NewIpfsAPI(c)
	return s
}

func (s *Client) WithTLS(tc *tls.Config) *Client {
	s.client.WithTLS(tc)
	return s
}

func (s *Client) WithTimeout(d time.Duration) *Client {
	s.client.WithTimeout(d)
	return s
}

func (s *Client) WithRetry(num int, delay time.Duration) *Client {
	s.client.WithRetry(num, delay)
	return s
}

func (s *Client) WithLogger(log log.Logger) *Client {
	s.client.WithLogger(log)
	return s
}

func (s *Client) WithCacheSize(sz int) *Client {
	s.client.WithCacheSize(sz)
	return s
}

func (s *Client) UseScriptCache(cache *lru.TwoQueueCache[Address, any]) {
	s.client.UseScriptCache(cache)
}

func (s Client) Retries() int {
	return s.client.Retries()
}

func (s Client) RetryDelay() time.Duration {
	return s.client.RetryDelay()
}

func (s Client) CacheGet(key Address) (any, bool) {
	return s.client.CacheGet(key)
}

func (s Client) CacheAdd(key Address, val any) {
	s.client.CacheAdd(key, val)
}
