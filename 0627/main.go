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

// Connect directly to MySQL via bytedmysql
func connectDirect() {
	// More DSN parameters: DSN (Data Source Name)
	dsn := "@sd(toutiao.mysql.rds_mysql_test_write)/?timeout=5s&use_gdpr_auth=true"
	db, err := sql.Open("bytedmysql", dsn)
	if err != nil {
		logger.Info("init db connect error: %s", err.Error())
		panic(err.Error())
	}
	// Execute queries; SQL usage can be searched online, RDS usage is consistent with third-party libraries.
	// .....
	db.Close()
}

// Connect to MySQL via Gorm, whose default driver is also bytedmysql
func connectByGorm() {
	db, err := gorm.Open(
		bytedgorm.MySQL("life.qa.ftf_measure", "lifeqa_infra_cn").With(func(conf *bytedgorm.DBConfig) { //Note: Gorm auto-appends RDS PSM suffix. Just provide toutiao.mysql.dbname; _write/_read are auto-appended. PSM is the RDS database psm, dbname is the database name. E.g. psm=toutiao.mysql.dbname_read where toutiao.mysql is the fixed PSM prefix, test is dbname, read means read consul; corresponding write consul is toutiao.mysql.test_write.
			conf.ReadTimeout = 2 * time.Second // Configure read timeout
		}).WithReadReplicas(),
		// WithReadReplicas enables read/write separation, uses toutiao.mysql.dbname_write as write db, toutiao.mysql.dbname_read as read db. When disabled, toutiao.mysql.dbname_write handles both reads and writes
		bytedgorm.ConnPool{MaxIdleConns: 200, MaxOpenConns: 200}, // Configure connection pool
	)
	if err != nil {
		logger.Info("init db connect error: %s", err.Error())
		panic(err.Error())
	}
	/// Execute queries; SQL usage can be searched online, RDS usage is consistent with third-party libraries.
	// .....

	fmt.Println(db)
}

func main() {
	connectByGorm()
}
