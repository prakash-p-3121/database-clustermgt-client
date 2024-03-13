package impl

import (
	"fmt"
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/restclientlib"
	"log"
	"net/url"
)

type DatabaseClusterMgtClientImpl struct {
	Host string
	Port uint
}

func (client *DatabaseClusterMgtClientImpl) HostPort() string {
	return fmt.Sprintf("https://%s:%d", client.Host, client.Port)
}

func (client *DatabaseClusterMgtClientImpl) FindCurrentWriteShard(tableName, resourceID string) (*model.DatabaseShard, errorlib.AppError) {
	restClient := restclientlib.NewRestClient()
	hostPort := client.HostPort()

	baseUrl := hostPort + "/" + findCurrentWriteShardUrl
	log.Println("url =", baseUrl)

	params := url.Values{}
	params.Add("table-name", tableName)
	params.Add("resource-id", resourceID)
	encodedParams := params.Encode()
	log.Println("encodedParams=" + encodedParams)
	finalUrl := baseUrl + "?" + encodedParams
	log.Println("finalUrl=" + finalUrl)

	var resp model.DatabaseShard
	appErr := restClient.Get(finalUrl, &resp)
	return &resp, appErr
}

func (client *DatabaseClusterMgtClientImpl) FindAllShardsByTable(tableName string) ([]*model.DatabaseShard, errorlib.AppError) {
	restClient := restclientlib.NewRestClient()
	hostPort := client.HostPort()

	baseUrl := hostPort + "/" + findAllShardsByTableUrl
	log.Println("url =", baseUrl)
	params := url.Values{}
	params.Add("table-name", tableName)
	encodedParams := params.Encode()
	log.Println("encodedParams=" + encodedParams)
	finalUrl := baseUrl + "?" + encodedParams
	log.Println("finalUrl=" + finalUrl)

	var resp []*model.DatabaseShard
	appErr := restClient.Get(finalUrl, &resp)
	log.Println(resp)
	return resp, appErr
}
