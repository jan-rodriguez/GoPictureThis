package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/VividCortex/mysqlerr"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gin-gonic/gin.v1"

	"./database"
	"./models"
	"./mysqlhelper"
)

func main() {

	db, connectError := sql.Open("mysql", "root:241neveragain@tcp(127.0.0.1:3306)/test")

	if connectError != nil {
		fmt.Println(connectError.Error())
	}

	createError := database.CreateTables(db)

	if createError != nil {
		fmt.Println(createError.Error())
	}

	defer db.Close()

	// make sure connection is available
	err := db.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := gin.Default()

	r.GET("/user/:user_id", func(c *gin.Context) {
		googleID := c.Param("user_id")

		user, err := database.GetUserFromGoogleID(db, googleID)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, user)
	})

	r.POST("/user", func(c *gin.Context) {
		var json models.User
		if c.BindJSON(&json) == nil {
			createdUser, err := database.CreateUser(db, json)
			if err != nil {
				if mysqlhelper.GetMysqlCodeForError(err) ==
					mysqlerr.ER_DUP_ENTRY {
					c.Status(http.StatusConflict)
				} else {
					c.Status(http.StatusInternalServerError)
				}
			} else {
				// TODO: Handle duplicate errors better
				c.JSON(http.StatusOK, createdUser)
			}
		}
	})

	r.GET("/user/:user_id/challenges", func(c *gin.Context) {
		userID := c.Param("user_id")

		active := c.DefaultQuery("active", "true")

		isActive, _ := strconv.ParseBool(active)

		result, err := database.GetChallengesForUser(db, userID, isActive)

		if err != nil {
			fmt.Print(err.Error())
			c.Status(http.StatusInternalServerError)
		}

		c.JSON(http.StatusOK, result)
	})

	r.GET("/user/:user_id/challenges/created", func(c *gin.Context) {
		userID := c.Param("user_id")

		active := c.DefaultQuery("active", "true")

		isActive, _ := strconv.ParseBool(active)

		result, err := database.GetChallengesCreatedByUser(db, userID, isActive)

		if err != nil {
			fmt.Print(err.Error())
			c.Status(http.StatusInternalServerError)
		}

		c.JSON(http.StatusOK, result)
	})

	r.POST("/response/:response_id/accept", func(c *gin.Context) {
		responseID := c.Param("response_id")

		err := database.AcceptResponse(db, responseID)

		if err != nil {
			fmt.Println(err.Error())
			c.Status(http.StatusInternalServerError)
		} else {
			c.Done()
		}
	})

	r.POST("/response/:response_id/decline", func(c *gin.Context) {
		responseID := c.Param("response_id")

		err := database.DeclineResponse(db, responseID)

		if err != nil {
			fmt.Print(err.Error())
			c.Status(http.StatusInternalServerError)
		} else {
			c.Done()
		}
	})

	r.POST("/challenge", func(c *gin.Context) {
		var json models.CreateChallenge
		if c.BindJSON(&json) == nil {
			challenge, err := database.CreateChallenge(db, json)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, challenge)
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Unrecognized json"})
		}
	})

	r.POST("/image", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		filename := header.Filename
		fmt.Println(header.Filename)
		out, err := os.Create("./tmp/" + filename)
		defer out.Close()
		if err != nil {
			fmt.Println(err.Error())
			c.String(http.StatusInternalServerError, "Failed at creating file")
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Println(err.Error())
			c.String(http.StatusInternalServerError, "Failed copying file")
			return
		}

		response := models.ImageCreatedResponse {
			Location: 
		}

		c.JSON(http.StatusOK, obj)
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
