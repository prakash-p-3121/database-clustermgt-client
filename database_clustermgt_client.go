package database_clustermgt_client

import (
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/errorlib"
)

type DatabaseClusterMgtClient interface {
	FindShard(tableName string, resourceID string) (*model.DatabaseShard, errorlib.AppError)
	FindAllShardsByTable(tableName string) ([]*model.DatabaseShard, errorlib.AppError)
}
