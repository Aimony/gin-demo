package adapter

import (
	"gin-demo/mysql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Init() error {
	dataSouce := mysql.NewDataSource()

	bookController := NewBookController(dataSouce)

	engine := gin.Default()
	routerGroupBook := engine.Group("/book")
	routerGroupBook.POST("/add", bookController.AddBook)
	routerGroupBook.POST("/update", bookController.UpdateBook)
	routerGroupBook.POST("/delete", bookController.DeleteBook)
	routerGroupBook.GET("/list", bookController.ListByBookId)

	return engine.Run()
}
