package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/saxumVermes/shop_orm_inmem/pkg/shop"
)

type shoeForm struct {
	Brand  string  `form:"brand"`
	Model  string  `form:"model"`
	Price  float32 `form:"price"`
	Colors string  `form:"colors[]"`
}

func init() {
	shop.Dialect, shop.DBUri = parseDBCred()
}

func port() string {
	port := strings.TrimSpace(os.Getenv("SHOE_SERVER_PORT"))
	match, err := regexp.MatchString("^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$", port)
	if err != nil {
		panic(fmt.Sprintf("regular expression: %v", err))
	}
	if port != "" && match {
		return ":" + port
	}
	return ":8080"
}

func parseDBCred() (dialect string, uri string) {
	parts := strings.Split(strings.TrimSpace(os.Getenv("DB_URL")), "://")
	dialect = parts[0]
	uri = parts[1]
	if match, err := regexp.MatchString("^(sqlite3|postgres|mysql)$", dialect); !match || err != nil {
		panic("connection failed! might be wrong input dialect")
	}

	return dialect, uri
}
