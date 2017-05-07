package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/VividCortex/mysqlerr"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gopkg.in/gin-gonic/gin.v1"

	"./database"
	"./models"
	"./mysqlhelper"
	"./utils"
)

func main() {

	username := os.Getenv("PICTURE_THIS_DB_USER")
	password := os.Getenv("PICTURE_THIS_DB_PASS")
	dbname := "test"

	db, connectError := gorm.Open("mysql", username+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()

	if connectError != nil {
		fmt.Println(connectError.Error())
		return
	}

	database.CreateTables(db)

	r := gin.Default()

	r.GET("/user/:user_id", func(c *gin.Context) {
		googleID := c.Param("user_id")

		user, err := database.GetUserFromGoogleID(db, googleID)

		if err != nil {
			utils.SendErrorResponse(c, err)
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
					utils.SendErrorResponse(c, err)
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
			utils.SendErrorResponse(c, err)
			return
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
			utils.SendErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, result)
	})

	r.POST("/challenge", func(c *gin.Context) {
		var createChallenge models.CreateChallenge
		if c.BindJSON(&createChallenge) != nil {
			challenge, err := database.CreateChallenge(db, createChallenge)
			if err != nil  {
				utils.SendErrorResponse(c, err)
			} else {
				c.JSON(http.StatusOK, challenge)
			}
		}
	})

	r.POST("/challenge/:challenge_id/response", func(c *gin.Context) {

		var response models.Response
		if c.BindJSON(&response) == nil {
			challengeId, err := strconv.ParseInt(c.Param("challenge_id"), 10, 0)

			if err != nil {
				utils.SendErrorResponse(c, err)
				return
			}

			response.ChallengeID = int(challengeId)
			updatedResponse, err := database.CreateResponse(db, response)
			if err != nil {
				utils.SendErrorResponse(c, err)
				return
			} else {
				c.JSON(http.StatusOK, updatedResponse)
			}
		}
	})

	r.POST("/response/:response_id/accept", func(c *gin.Context) {

		updatedResponse, err := database.AcceptResponse(db, c.Param("response_id"))
		if err != nil {
			utils.SendErrorResponse(c, err)
			return
		} else {
			c.JSON(http.StatusOK, updatedResponse)
		}
	})

	r.POST("/response/:response_id/decline", func(c *gin.Context) {

		updatedResponse, err := database.DeclineResponse(db, c.Param("response_id"))
		if err != nil {
			utils.SendErrorResponse(c, err)
			return
		} else {
			c.JSON(http.StatusOK, updatedResponse)
		}
	})

	r.POST("/image", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		fmt.Println(header.Filename)

		// Create images dir if it doesn't exist
		if _, err := os.Stat("images"); os.IsNotExist(err) {
			os.Mkdir("images", os.ModePerm)
		}

		output, err := exec.Command("uuidgen").Output()
		// Trim out the new line at the end
		filename := string(output[:len(output)-1])
		if err != nil {
			utils.SendErrorResponse(c, err)
			return
		}
		out, err := os.Create("./images/" + filename)

		if err != nil {
			utils.SendErrorResponse(c, err)
			return
		}

		defer out.Close()

		if err != nil {
			utils.SendErrorResponse(c, err)
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			utils.SendErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"filename": filename})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
