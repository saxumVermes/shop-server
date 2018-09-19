* Run `export DB_URL="mysql://user:mysql@tcp(localhost:[MYSQL_PORT])/shop?charset=utf8&parseTime=True&loc=Local"`
* Docker mysql: run `./mysql.sh <port>`

```
Additional ENVs:

  SHOE_SERVER_PORT: defaults to 8080
  SHOE_TEST_ENV: store data in memory, boolean value
  DB_URL: accepts valid database uri: valid postgres, mysql, sqlite3
```
[Check gorm documentation for valid db uri syntax](http://doc.gorm.io/database.html#connecting-to-a-database)

```
Examples:

MYSQL: mysql://user:password@protocol(host:port)/database?charset=utf8&parseTime=True&loc=Local
SQLITE3: sqlite3://database.db
```

* cd to frontend, run `./run.sh setup <port>` and build the app
