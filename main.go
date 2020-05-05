package main

import (
	"net/http"
	"strconv"

	"github.com/al_hadyd/http-service/model"
	"github.com/labstack/echo"
)

//QueryParam
/* func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		user := c.QueryParam("user")
		return c.String(http.StatusOK, user)
	})
	e.Logger.Fatal(e.Start(":8080"))
} */

//FormValue
/* func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		user := c.FormValue("user")
		return c.String(http.StatusOK, user)
	})
	e.Logger.Fatal(e.Start(":8080"))
} */

/* func main() {
	e := echo.New()
	e.GET("articles", func(c echo.Context) error {
		return c.String(http.StatusOK, "Untuk Mendapatkan Data List")
	})
	e.GET("articles/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Untuk Mendapatkan Data List ID")
	})
	e.POST("articles", func(c echo.Context) error {
		return c.String(http.StatusOK, "Untuk Create Article")
	})
	e.PUT("articles/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Untuk Update Article")
	})
	e.DELETE("articles/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Untuk Delete Article")
	})

	e.Logger.Fatal(e.Start(":8080"))
} */

func main() {

	store := model.NewArticleStoreInmemory()

	e := echo.New()
	e.GET("/articles", func(c echo.Context) error {
		articles := store.ArticleMap
		return c.JSON(http.StatusOK, articles)
	})

	e.POST("/articles", func(c echo.Context) error {
		title := c.FormValue("Title")
		body := c.FormValue("Body")
		article, _ := model.CreateArticle(title, body)
		store.Save(article)
		return c.JSON(http.StatusOK, article)
	})

	e.GET("/articles/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		article := store.ArticleMap
		return c.JSON(http.StatusOK, article[id-1])
	})
	e.DELETE("/articles/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		article := store.Remove(id)
		return c.JSON(http.StatusOK, article)
	})
	e.PUT("/articles/:id", func(c echo.Context) error {
		title := c.FormValue("Title")
		body := c.FormValue("Body")
		id, _ := strconv.Atoi(c.Param("id"))
		article := store.EditArticle(title, body, id)
		return c.JSON(http.StatusOK, article)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
