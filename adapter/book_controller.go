package adapter

import (
	"gin-demo/model"
	"gin-demo/mysql"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"log"
	"net/http"
)

type BookController struct {
	dataSource *mysql.DataSouce
}

func NewBookController(dataSource *mysql.DataSouce) BookController {
	return BookController{
		dataSource: dataSource,
	}
}

func (b *BookController) AddBook(ctx *gin.Context) {
	var book model.Book
	if err := ctx.ShouldBind(&book); err != nil {
		logrus.Error("读取Book信息失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
		return
	}
	bookDao := b.dataSource.BookDao()
	err := bookDao.AddBook(&book)
	if err != nil {
		logrus.Error("添加Book失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (b *BookController) UpdateBook(ctx *gin.Context) {
	log.Printf("Entering UpdateBook handler")
	var book model.Book
	if err := ctx.ShouldBind(&book); err != nil {
		logrus.Error("读取Book信息失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
		return
	}
	bookDao := b.dataSource.BookDao()
	err := bookDao.UpdateBook(&book)
	if err != nil {
		logrus.Error("更新Book失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (b *BookController) DeleteBook(ctx *gin.Context) {
	var book model.Book
	if err := ctx.ShouldBind(&book); err != nil {
		logrus.Error("读取Book信息失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
		return
	}
	bookDao := b.dataSource.BookDao()
	err := bookDao.DeleteBook(&book)
	if err != nil {
		logrus.Error("删除Book失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (b *BookController) ListByBookId(ctx *gin.Context) {
	id := cast.ToUint(ctx.Query("id"))
	bookDao := b.dataSource.BookDao()
	book, err := bookDao.ListByBookId(id)
	if err != nil {
		logrus.Error("查询Book失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, book)
}
