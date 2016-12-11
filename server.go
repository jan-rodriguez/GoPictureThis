package main

import (
	"fmt"
	"net/http"
    "gopkg.in/gin-gonic/gin.v1"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "./models"
	"./database"
)

func main() {

	// TODO: Change these to new database
    db, connectError := sql.Open("mysql", "root:neveragain@tcp(127.0.0.1:3306)/test")

	if connectError != nil {
		fmt.Print(connectError.Error())
	}

	createError := database.CreateTables(db)

	if createError != nil {
		fmt.Print(createError.Error())
	}

	defer db.Close()

	// make sure connection is available
	err := db.Ping()

	if err != nil {
		fmt.Print(err.Error())
	}

    r := gin.Default()
    r.GET("/user/:user_id/challenges", func(c *gin.Context) {
		user_id := c.Param("user_id")

		result, err := database.GetChallengesForUser(db, user_id, false)

		if (err != nil) {
			fmt.Print(err.Error())
		}

        c.JSON(http.StatusOK, result)
    })

	r.POST("/challenge", func(c *gin.Context) {
		var json models.Create_Challenge
		if c.BindJSON(&json) == nil {
			err := database.CreateChallenge(db, json)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				c.Done()
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Unrecognized json"})
		}
	})

    r.PUT("/test/:name", func(c *gin.Context) {

        name := c.Param("name")

        stmt, err := db.Prepare("INSERT INTO test (name) values(?);")
		if err != nil {
			fmt.Print(err.Error())
		}
        defer stmt.Close()

        _, err = stmt.Exec(name)

		if err != nil {
			fmt.Print(err.Error())
		}

        c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", name),
		})
    })

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
