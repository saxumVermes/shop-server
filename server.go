package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/saxumVermes/shop-server/pkg/shop"
)

type shoeForm struct {
	Brand  string  `form:"brand" json:"brand"`
	Model  string  `form:"model" json:"model"`
	Price  float32 `form:"price" json:"price"`
	Colors string  `form:"colors[]" json:"colors[]"`
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

func parseDBCred() {
	parts := strings.Split(strings.TrimSpace(os.Getenv("DB_URL")), "://")
	if len(parts) != 2 {
		fmt.Fprintln(os.Stderr, "invalid db uri, check README.md or set SHOE_TEST_ENV to true for in memory storage")
		os.Exit(1)
	}
	dialect := parts[0]
	uri := parts[1]
	if match, err := regexp.MatchString("^(sqlite3|postgres|mysql)$", dialect); !match || err != nil {
		panic(fmt.Sprintf("connection failed! might be wrong input dialect: %v", err))
	}

	shop.Dialect, shop.DBUri = dialect, uri
}
