package sqlparser

import (
	"reflect"
	"testing"
)

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("%+v", err)
	}
}

func checkEqual(t *testing.T, src interface{}, dst interface{}) {
	if !reflect.DeepEqual(src, dst) {
		t.Fatalf("not equal %v and %v", src, dst)
	}
}

func TestDropTableParsing(t *testing.T) {
	_, err := Parse(`DROP table if exists users`)
	checkErr(t, err)
}

func TestSelectParsing(t *testing.T) {
	_, err := Parse(`SELECT * from users where id = 1 order by created_at limit 1 offset 3`)
	checkErr(t, err)
}

func TestCreateTableParsing(t *testing.T) {
	_, err := Parse(`
		CREATE TABLE users (
		  id bigint(20) unsigned NOT NULL,
		  other_id bigint(20) unsigned NOT NULL,
		  enum_column enum('a','b','c','d') DEFAULT NULL,
		  int_column int(10) DEFAULT '0',
		  PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	checkErr(t, err)
}

func TestCreateTableWithPartition(t *testing.T) {
	_, err := Parse(`
CREATE TABLE histories (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  user_id bigint(20) unsigned NOT NULL,
  note text NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  PRIMARY KEY (id,created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
/*!50500 PARTITION BY RANGE  COLUMNS(created_at)
(PARTITION p201812 VALUES LESS THAN ('2019-01-01') ENGINE = InnoDB,
 PARTITION p201901 VALUES LESS THAN ('2019-02-01') ENGINE = InnoDB,
 PARTITION p201902 VALUES LESS THAN ('2019-03-01') ENGINE = InnoDB,
 PARTITION p201903 VALUES LESS THAN ('2019-04-01') ENGINE = InnoDB) */;
`)
	checkErr(t, err)
}

func TestShowCreateTableParsing(t *testing.T) {
	ast, err := Parse(`SHOW CREATE TABLE users`)
	checkErr(t, err)
	switch stmt := ast.(type) {
	case *Show:
		checkEqual(t, "users", stmt.TableName)
	default:
		t.Fatalf("%+v", "type mismatch")
	}
}
