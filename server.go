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

        active := c.DefaultQuery("active", "true")

        is_active, parse_err := strconv.ParseBool(active)

        result, err := database.GetChallengesForUser(db, user_id, is_active)

        if (err != nil) {
            fmt.Print(err.Error())
            c.Status(http.StatusInternalServerError)
        }

        c.JSON(http.StatusOK, result)
    })

    r.GET("/user/:user_id/challenges/created", func(c *gin.Context) {
        user_id := c.Param("user_id")

        active := c.DefaultQuery("active", "true")

        is_active, parse_err := strconv.ParseBool(active)

        result, err := database.GetChallengesCreatedByUser(db, user_id, is_active)

        if (err != nil) {
            fmt.Print(err.Error())
            c.Status(http.StatusInternalServerError)
        }

        c.JSON(http.StatusOK, result)
    })

    r.POST("/response/:response_id/accept", func (c *gin.Context) {
        response_id := c.Param("response_id")

        err := database.AcceptResponse(db, response_id)

        if (err != nil) {
            fmt.Print(err.Error())
            c.Status(http.StatusInternalServerError)
        } else {
            c.Done()
        }
    })

    r.POST("/response/:response_id/decline", func (c *gin.Context) {
        response_id := c.Param("response_id")

        err := database.DeclineResponse(db, response_id)

        if (err != nil) {
            fmt.Print(err.Error())
            c.Status(http.StatusInternalServerError)
        } else {
            c.Done()
        }
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

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
