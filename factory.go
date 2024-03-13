package database_clustermgt_client

import "github.com/prakash-p-3121/database-clustermgt-client/impl"

func NewDatabaseClusterMgtClient(host string, port uint) DatabaseClusterMgtClient {
	return &impl.DatabaseClusterMgtClientImpl{Host: host, Port: port}
}
