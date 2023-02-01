package g

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"redrock-test/model/config"
)

var (
	Config *config.Config
	Mdb    *sql.DB
	Rdb    *redis.Client
	Logger *zap.Logger
)
