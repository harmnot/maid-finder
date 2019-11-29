package main

import (
	"bantoo/Http"
	"bantoo/connection"
	"bantoo/middleware/Auth"
	"bantoo/models"
	"github.com/gin-contrib/cors"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	db := connection.ConnectDB()

	defer db.Close()
	router := gin.Default()
	router.Use(cors.Default())

	r := router.Group("/api")
	{
		r.GET("/ping", func(g *gin.Context) {
			var people models.Person

			db.Preload("Data").Where(&models.Person{Email: "ona@email.com", Password: "12345678"}).Find(&people)
			g.JSON(http.StatusOK, people)
		})

		r.GET("/test", Auth.CheckHeader(), Auth.ReadTokenHeader(db), func(g *gin.Context) {
			var people []models.Person
			db.Exec("select * from people").Find(&people)
			g.JSON(200, people)
		})
		r.POST("/login", Http.Login(db))
		r.GET("/filtermaid", Http.FilterMaid(db))
		r.POST("/createUser/:token", Http.CreatePeopleFromToken(db))
		r.PUT("/updateUserToken", Http.UpdateDataUser(true, false,  db))
		r.PUT("/updatePassword", Http.UpdateDataUser(false, true,  db))
		r.PUT("/updateUser", Http.UpdateDataUser(false, false,  db))
		r.POST("/forgotPassword", Http.ForgotPassword(db))
		r.DELETE("/deletetype", Http.DeleteTypeAndSkillMaid("delete_type", db))
		r.DELETE("/deleteskill", Http.DeleteTypeAndSkillMaid("delete_skill", db))
		r.POST("/upload/:as/:id", Http.UploadFile(db))

		r.PUT("/get", Http.Find(db))
		r.POST("/register", Http.RegistrationBySendEmail(db))
		r.POST("/add", Http.CreatePeople(db))
		r.POST("/maid", Http.CreateProvider(db))
		r.PUT("/addmaid", Http.ProviderAddMaids(db))
	}

	_ = router.Run(":2000")
}
