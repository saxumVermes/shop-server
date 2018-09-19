package main

// Endpoints:
//  /shoes
//  /shoes/add
//  /shoes/{id}
//
// Example request:
//    curl -X POST localhost:{port}/shoes/add --form brand=adidas --form model=air --form price=43.55 --form colors="black, blue"
//    curl localhost:{port}/shoes/{id}
//    curl -X DELETE localhost:{port}/shoes/{id}/delete
//    curl localhost:{port}/shoes
//
// ENV:
// 	SHOE_SERVER_PORT=8080
// 	DB_URL=sqlite3://test.db
//  SHOE_TEST_ENV=true  -> sets storage to in memory

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/saxumVermes/shop_orm_inmem/pkg/shop"
)

const VERSION = "1.0.0"

var GITCOMMIT = "???"

func addHandler(ss shop.ShoeServer) func(*gin.Context) {
	return func(c *gin.Context) {
		var sf shoeForm
		if c.ShouldBind(&sf) == nil {
			if sf.Brand != "" && sf.Model != "" {
				shoe, ok := ss.Add(sf.Brand, sf.Model, sf.Price, sf.Colors)
				if !ok {
					c.JSON(409, gin.H{"warning": "product already exists with that id"})
					return
				}
				c.IndentedJSON(201, gin.H{
					shoe.SID: shoe,
				})
			} else {
				c.JSON(200, gin.H{"status": "Required fields are empty"})
			}
		} else {
			c.JSON(500, gin.H{"status": "Error occured while adding shoe"})
			fmt.Fprintln(os.Stdout, "Error occured while adding shoe")
		}
	}
}

func listHandler(ss shop.ShoeServer) func(*gin.Context) {
	return func(c *gin.Context) {
		p := gin.H{}
		ps := ss.List()
		for _, v := range ps {
			p[v.SID] = v
		}
		c.IndentedJSON(200, p)
	}
}

func shoePageHandler(ss shop.ShoeServer) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		s, found := ss.Find(id)
		if !found {
			c.IndentedJSON(404, gin.H{"status": "Shoe not found with that id"})
			return
		}
		c.IndentedJSON(200, gin.H{id: s})
	}
}

func deleteHandler(ss shop.ShoeServer) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if ss.DeleteById(id) {
			c.JSON(200, gin.H{"status": fmt.Sprintf("Shoe %s has been deleted", id)})
		} else {
			c.JSON(404, gin.H{"status": fmt.Sprintf("Shoe %s not found", id)})
		}
	}
}

func defaultHandler(c *gin.Context) {
	c.Request.Header.Set("Access-Control-Allow-Origin", "*")
	c.Request.Header.Set("Access-Control-Allow-Methods", "*")
	c.Request.Header.Set("Access-Control-Allow-Header", "*")
	c.IndentedJSON(200, gin.H{
		"status": "ok",
		"actions": gin.H{
			"/version": "GET",
			"/":        "GET",
			"/shoes": []string{
				"GET",
				"POST",
			},
			"/shoes/{id}": []string{
				"GET",
				"DELETE",
			},
		},
	})
}

func versionHandler(c *gin.Context) {
	c.JSON(200, gin.H{"version": fmt.Sprintf("%s-alpha+%s", VERSION, GITCOMMIT)})
}

func NewHandler(ss shop.ShoeServer, r *gin.Engine) {
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	r.OPTIONS("/*path", defaultHandler)
	r.POST("/shoes", addHandler(ss))
	r.GET("/shoes", listHandler(ss))
	r.GET("/shoes/:id", shoePageHandler(ss))
	authorized.DELETE("/shoes/:id", deleteHandler(ss))
	r.GET("/version", versionHandler)
	authorized.GET("/", defaultHandler)
}

func main() {
	r := gin.Default()
	ss := shop.New("Deichmann", "Budapest", 1188)
	NewHandler(ss, r)
	// Defaults to 8080
	r.Run(port())
}
