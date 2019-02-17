package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hatobus/TDK_gakushoku/cmd/presenter"
	"github.com/hatobus/TDK_gakushoku/cmd/util"
)

func main() {

	_, err := util.Init()
	if err != nil {
		log.Fatalln(err)
	}
	conf := util.GetConfig()

	source := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8",
		conf.MySQL.User,
		conf.MySQL.Password,
		conf.MySQL.Host,
		conf.MySQL.Port,
		conf.MySQL.Database)

	log.Println(source)

	_, err = presenter.InitEngine(source)
	if err != nil {
		log.Fatalln(err)
	}

	os.Setenv("SOURCE", source)

	// err = presenter.InsertDummyUser()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	r := gin.Default()

	r.GET("/rank", GetRanking)
	r.POST("/new", CreateWork)

	r.Run(":8080")
}

func GetRanking(c *gin.Context) {
	top10, err := presenter.GetTop10()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, top10)
}

func CreateWork(c *gin.Context) {

}
