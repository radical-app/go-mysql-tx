package toolbox_test

import (
	_ "github.com/go-sql-driver/mysql"
)

const TEST_SCHEME_CREATE=`CREATE TABLE test_db.db_table_1 (
    id INT NOT NULL AUTO_INCREMENT ,
    name VARCHAR(45) NULL,
    PRIMARY KEY (id));`

const TEST_TABLE_CREATE=`CREATE TABLE IF NOT EXISTS db_table_1 (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(145) DEFAULT NULL,
   PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

const TEST_TABLE_INSERT=`INSERT INTO test_db.db_table_1 (id, name) VALUES (null, ?);`

const TEST_TABLE_SELECT=`SELECT id, name FROM test_db.db_table_1 where id = ?;`

const TEST_TABLE_DROP=`DROP TABLE IF EXISTS db_table_1;`


