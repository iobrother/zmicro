package client

import (
	"github.com/iobrother/zmicro/core/log"
	etcd_client "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"go.opentelemetry.io/otel"
)

type Client struct {
	opts    Options
	xClient client.XClient
}

func NewClient(opts ...Option) *Client {
	options := newOptions(opts...)

	c := &Client{opts: options}

	if len(c.opts.EtcdAddr) > 0 {
		d, err := etcd_client.NewEtcdV3Discovery(
			c.opts.BasePath,
			c.opts.ServiceName,
			c.opts.EtcdAddr,
			false,
			nil,
		)
		if err != nil {
			log.Fatal(err.Error())
		}
		opt := client.DefaultOption
		opt.SerializeType = protocol.ProtoBuffer
		c.xClient = client.NewXClient(
			c.opts.ServiceName,
			client.Failtry,
			client.RoundRobin,
			d,
			opt,
		)
	} else {
		d, err := client.NewPeer2PeerDiscovery("tcp@"+c.opts.ServiceAddr, "")
		if err != nil {
			log.Fatal(err.Error())
		}

		opt := client.DefaultOption
		opt.SerializeType = protocol.ProtoBuffer
		c.xClient = client.NewXClient(c.opts.ServiceName, client.Failtry, client.RoundRobin, d, opt)
	}

	if c.opts.Tracing {
		tracer := otel.Tracer("rpcx")
		p := client.NewOpenTelemetryPlugin(tracer, nil)
		pc := client.NewPluginContainer()
		pc.Add(p)
		c.xClient.SetPlugins(pc)
	}

	return c
}

func (c *Client) GetXClient() client.XClient {
	return c.xClient
}
