package main

import (
	"dbsqlite/pkg/records"

	"github.com/gin-gonic/gin"
)

func main() {

	records.DBCSqlite("db.db")

	router := gin.Default()
	group := router.Group("/api/v1/records")

	group.Use(func(c *gin.Context) {
		c.Keys = make(map[string]interface{})

		flagz := records.Flags{
			Action:   c.Request.Method,
			Category: c.Query("category"),
			Domain:   c.Query("domain"),
			Args:     []string{"-"},
			Rdr:      c.Copy().Request.Body,
		}
		if flagz.Category == "" || flagz.Domain == "" {
			c.JSON(200, gin.H{
				"message": "category and domain are required",
				"error":   "category and domain are required",
			})
			c.Abort()
			return
		}
		c.Keys["flagz"] = &flagz
		c.Next()
	})

	group.POST("", func(c *gin.Context) {
		flagz := c.Keys["flagz"].(*records.Flags)
		data, err := records.Create(flagz)

		c.Keys["message"] = "list"
		c.Keys["error"] = err
		c.Keys["data"] = data
		c.JSON(200, c.Keys)
	})

	group.GET("", func(c *gin.Context) {
		flagz := c.Keys["flagz"].(*records.Flags)
		data, err := records.List(flagz)

		c.Keys["message"] = "list"
		c.Keys["error"] = err
		c.Keys["data"] = data
		c.JSON(200, c.Keys)
	})

	group.GET("/:uuid", func(c *gin.Context) {
		flagz := c.Keys["flagz"].(*records.Flags)
		flagz.Args = []string{c.Param("uuid")}
		data, err := records.Read(flagz)

		c.Keys["message"] = "read"
		c.Keys["error"] = err
		c.Keys["data"] = data
		c.JSON(200, c.Keys)
	})

	group.DELETE("/:uuid", func(c *gin.Context) {
		flagz := c.Keys["flagz"].(*records.Flags)
		flagz.Args = []string{c.Param("uuid")}
		data, err := records.Delete(flagz)

		c.Keys["message"] = "delete"
		c.Keys["error"] = err
		c.Keys["data"] = data
		c.JSON(200, c.Keys)
	})
	group.PUT("/:uuid", func(c *gin.Context) {
		flagz := c.Keys["flagz"].(*records.Flags)
		flagz.Args = []string{c.Param("uuid"), "-"}
		data, err := records.Update(flagz)

		c.Keys["message"] = "update"
		c.Keys["error"] = err
		c.Keys["data"] = data
		c.JSON(200, c.Keys)
	})

	router.Run(":9999")
}
