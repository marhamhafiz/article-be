package posts

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id          int       `json:"id"`
	Title       string    `json:"title" validate:"required,min=20"`
	Content     string    `json:"content" validate:"required,min=200"`
	Category    string    `json:"category" validate:"required,min=3"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Status      string    `json:"status" validate:"required,oneof='publish' 'draft' 'trash'"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func connectToDB() *sql.DB {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/articles?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to Database")
	return db
}

func getArticles(c *gin.Context) {
	article_limit := c.Query("limit")
	article_offset := c.Query("offset")

	var getQuery string

	db := connectToDB()
	if article_limit == "" && article_offset == "" {
		getQuery = "SELECT * FROM posts ORDER BY id DESC"
	} else {
		getQuery = "SELECT * FROM posts ORDER BY id DESC LIMIT " + article_limit + " OFFSET " + article_offset
	}
	selDB, err := db.Query(getQuery)

	if err != nil {
		panic(err.Error())
	}
	articles := Article{}
	res := []Article{}
	for selDB.Next() {
		var id int
		var title, content, category, status string
		var created_date, updated_date time.Time
		err = selDB.Scan(&id, &title, &content, &category, &created_date, &updated_date, &status)
		if err != nil {
			panic(err.Error())
		}

		articles.Id = id
		articles.Title = title
		articles.Content = content
		articles.Category = category
		articles.CreatedDate = created_date
		articles.UpdatedDate = updated_date
		articles.Status = status
		res = append(res, articles)
	}
	defer db.Close()

	c.IndentedJSON(http.StatusOK, res)
}

func getArticleById(c *gin.Context) {
	article_id := c.Param("id")

	db := connectToDB()
	selDB, err := db.Query("SELECT * FROM posts WHERE id = " + article_id)
	if err != nil {
		panic(err.Error())
	}
	articles := Article{}
	res := []Article{}
	for selDB.Next() {
		var id int
		var title, content, category, status string
		var created_date, updated_date time.Time
		err = selDB.Scan(&id, &title, &content, &category, &created_date, &updated_date, &status)
		if err != nil {
			panic(err.Error())
		}

		articles.Id = id
		articles.Title = title
		articles.Content = content
		articles.Category = category
		articles.CreatedDate = created_date
		articles.UpdatedDate = updated_date
		articles.Status = status
		res = append(res, articles)
	}
	defer db.Close()

	c.IndentedJSON(http.StatusOK, res)
}

func validateArticle(article Article) (bool, map[string]string) {
	validate = validator.New()
	err := validate.Struct(article)

	if err != nil {
		aValidate := make(map[string]string)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return true, nil
		}

		for _, err := range err.(validator.ValidationErrors) {
			aValidate["error"] = err.StructField() + " " + err.ActualTag() + " " + err.Param()
			return false, aValidate
		}
		return false, nil
	}

	return true, nil
}

func postArticle(c *gin.Context) {
	var article_post Article

	db := connectToDB()
	if err := c.BindJSON(&article_post); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		valid, err_valid := validateArticle(article_post)

		if valid {
			insert, err := db.Query("INSERT INTO posts (title,content,category,status,created_date,updated_date) VALUES (?,?,?,?, now(), now())", article_post.Title, article_post.Content, article_post.Category, article_post.Status)
			// if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}

			fmt.Println("Insert to Database success")
			// be careful deferring Queries if you are using transactions
			defer insert.Close()
			c.IndentedJSON(http.StatusCreated, `{}`)
		} else {
			fmt.Println("Insert Article failed, data not valid")
			c.IndentedJSON(http.StatusCreated, err_valid)
		}
	}
}

func updateArticle(c *gin.Context) {
	article_id := c.Param("id")
	var article_post Article

	db := connectToDB()
	if err := c.BindJSON(&article_post); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		valid, err_valid := validateArticle(article_post)

		if valid {
			update, err := db.Query("UPDATE posts SET title='" + article_post.Title + "',content='" + article_post.Content + "',category='" + article_post.Category + "',status='" + article_post.Status + "',updated_date=now() WHERE id=" + article_id)
			// if there is an error updateing, handle it
			if err != nil {
				panic(err.Error())
			}

			fmt.Println("Update Article to Database success")
			// be careful deferring Queries if you are using transactions
			defer update.Close()
			c.IndentedJSON(http.StatusCreated, `{}`)
		} else {
			fmt.Println("Update Article failed, data not valid")
			c.IndentedJSON(http.StatusCreated, err_valid)
		}
	}
}

func deleteArticle(c *gin.Context) {
	article_id := c.Param("id")

	db := connectToDB()
	delete, err := db.Query("DELETE FROM posts WHERE id=" + article_id)
	// if there is an error deleteing, handle it
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Delete Article from Database success")
	// be careful deferring Queries if you are using transactions
	defer delete.Close()
	c.IndentedJSON(http.StatusCreated, `{}`)
}

func HandleRequests() {
	fmt.Println("server run...")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	dataRoutes := router.Group("api")
	{
		dataRoutes.GET("/article", getArticles)          //get all article
		dataRoutes.GET("/article/:id", getArticleById)   //get article by id
		dataRoutes.POST("/article", postArticle)         //create new article
		dataRoutes.PATCH("/article/:id", updateArticle)  //update an article
		dataRoutes.DELETE("/article/:id", deleteArticle) //delete an article
	}

	// router.Use(cors.Default())

	// router.GET("/article", getArticles)          //get all article
	// router.GET("/article/:id", getArticleById)   //get article by id
	// router.POST("/article", postArticle)         //create new article
	// router.PATCH("/article/:id", updateArticle)  //update an article
	// router.DELETE("/article/:id", deleteArticle) //delete an article

	router.Run("127.0.0.1:10000")
}
