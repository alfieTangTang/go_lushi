package dal

import (
	"github.com/alfieTangTang/go_lushi/app/product/biz/dal/mysql"
	"github.com/alfieTangTang/go_lushi/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
