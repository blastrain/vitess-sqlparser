package sqlparser

import (
	"testing"
)

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("%+v", err)
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
