package utils

import (
	"testing"
)

func TestAdaptivePostgresqlDsn(t *testing.T) {
	pgDsns := []string{
		"postgres://root:123456@192.168.3.37:5432/account",
		"postgres://root:123456@192.168.3.37:5432/account?sslmode=disable",
		"postgres://root:123456@192.168.3.37:5432/account?TimeZone=Asia/Shanghai",
		"root:123456@192.168.3.37:5432/account",
		"host=192.168.3.37 port=5432 user=root password=123456 dbname=account sslmode=disable",

		"root:123456@(192.168.3.37:5432)/account",
		"postgres://root:123456@(192.168.3.37:5432)/account",
		"postgres://root:123456@(192.168.3.37:5432)/account?TimeZone=Asia/Shanghai",
	}

	for _, v := range pgDsns {
		dsn := AdaptivePostgresqlDsn(v)
		t.Log(dsn)
	}
}

func TestAdaptiveMysqlDsn(t *testing.T) {
	mysqlDsns := []string{
		"root:123456@(192.168.3.37:3306)/account",
		"mysql://root:123456@(192.168.3.37:3306)/account",
	}

	for _, v := range mysqlDsns {
		dsn := AdaptiveMysqlDsn(v)
		t.Log(dsn)
	}
}
