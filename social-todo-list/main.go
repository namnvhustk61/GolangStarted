package main

import (
	"fmt"
	"log"
	"net/http"

	ginitem "social-todo-list/modules/item/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// `id` int NOT NULL AUTO_INCREMENT,
//   `title` varchar(255) DEFAULT NULL,
//   `image` json DEFAULT NULL,
//   `description` text,
//   `status` enum('Doing','Done','Deleted') DEFAULT 'Doing',
//   `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
//   `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3307)/todo_list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)

	// CRUD: Create Read Update Delete
	// POST /v1/items create new item
	// GET /v1/items list items  /v1/items?page=1
	// GET /v1/items/:id  get item by id
	// PUTorPATCH /v1/items/:id  update item by id
	// DELETE /v1/items/:id  delete item by id

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
