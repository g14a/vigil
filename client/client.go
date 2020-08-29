package client

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"sync"
	"time"
	"vigil/utils"
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

func NodeInfo(esClient elasticsearch.Client) NodeInfoResponse {
	resp, err := esClient.Info(func(request *esapi.InfoRequest) {
		request.Human = true
	})

	utils.CheckError(err)

	var nodeInfo NodeInfoResponse

	if err := json.NewDecoder(resp.Body).Decode(&nodeInfo); err != nil {
		utils.CheckError(err)
	}

	return nodeInfo
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
