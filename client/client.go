package client

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"sync"
	"time"
)

var (
	once     sync.Once
	ESClient elasticsearch.Client
)

func InitESClient() {
	once.Do(func() {
		es, _ := elasticsearch.NewDefaultClient()
		ESClient = *es
	})
}

func NodeInfo(client elasticsearch.Client) esapi.Info {
	return client.Info
}

func ClusterInfo(client elasticsearch.Client) (*esapi.Response, error) {
	res, err := client.Cluster.Health(
		client.Cluster.Health.WithTimeout(time.Duration(500000)),
		client.Cluster.Health.WithWaitForStatus("yellow"),
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}
