package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "code.byted.org/gopkg/bytedmysql"
	"code.byted.org/gorm/bytedgorm"
	"github.com/bytedance/gopkg/util/logger"
	"gorm.io/gorm"
)

// 通过 bytedmysql 直连MySQL
func connectDirect() {
	// 更多DSN参数见：DSN (Data Source Name)
	dsn := "@sd(toutiao.mysql.rds_mysql_test_write)/?timeout=5s&use_gdpr_auth=true"
	db, err := sql.Open("bytedmysql", dsn)
	if err != nil {
		logger.Info("init db connect error: %s", err.Error())
		panic(err.Error())
	}
	// 执行查询,SQL语句用法可自行搜索，RDS用法与各个第三方库是一致的。
	// .....
	db.Close()
}

// 通过 Gorm 连接MySQL，其中 Gorm 默认使用的driver也是bytedmysql
func connectByGorm() {
	db, err := gorm.Open(
		bytedgorm.MySQL("life.qa.ftf_measure", "lifeqa_infra_cn").With(func(conf *bytedgorm.DBConfig) { //这里注意gorm默认自动补充RDS PSM 后缀，业务正常只需要填写toutiao.mysql.dbname即可，_write/_read会自动补充。这里PSM是RDS数据库的psm，dbname是数据库名，例如psm为toutiao.mysql.dbname_read，其中toutiao.mysql是PSM的固定前缀，test是dbname，read表示这个是读consul，相应的写consul就是toutiao.mysql.test_write。
			conf.ReadTimeout = 2 * time.Second // 配置read超时时间
		}).WithReadReplicas(),
		// WithReadReplicas 开启读写分离, 将使用toutiao.mysql.dbname_write 做为写db, toutiao.mysql.dbname_read 做为读db, 不开启时将使用 toutiao.mysql.dbname_write 进行读写
		bytedgorm.ConnPool{MaxIdleConns: 200, MaxOpenConns: 200}, // 配置连接池信息
	)
	if err != nil {
		logger.Info("init db connect error: %s", err.Error())
		panic(err.Error())
	}
	/// 执行查询,SQL语句用法可自行搜索，RDS用法与各个第三方库是一致的。
	// .....

	fmt.Println(db)
}

func main() {
	connectByGorm()
}
