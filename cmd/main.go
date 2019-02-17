package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hatobus/TDK_gakushoku/cmd/models"
	"github.com/hatobus/TDK_gakushoku/cmd/presenter"
	"github.com/hatobus/TDK_gakushoku/cmd/slackbot"
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
	os.Setenv("SlackTOKEN", "xoxp-552584571425-552408532688-554403447383-d55835881f49461fa9895046640229d5")
	os.Setenv("slackChannelID", "CG951BD6X")

	// err = presenter.InsertDummyUser()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders: []string{"content-type", "boundary"},
	}))

	r.GET(conf.BaseURL+"/rank", GetRanking)
	r.POST(conf.BaseURL+"/new", CreateWork)
	r.POST(conf.BaseURL+"/user/accept", GetAcceptUser)

	r.Run(conf.MySQL.Local)
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
	creq := &models.PostReq{}
	err := c.BindJSON(creq)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	var work string
	switch creq.Category {
	case "1":
		log.Println("雑用")
		work = "雑用"
	case "2":
		log.Println("Geek Dojo")
		work = "Geek Dojo"
	case "3":
		log.Println("フロントエンド")
		work = "フロントエンド"
	case "4":
		log.Println("サーバーサイド")
		work = "サーバーサイド"
	case "5":
		log.Println("インフラ")
		work = "インフラ"
	case "6":
		log.Println("セキュリティ")
		work = "セキュリティ"
	}

	err = slackbot.PostNewTalk(creq.UserID, work)
	if err != nil {
		log.Println("slackbot post error")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "posted",
	})

}

func GetAcceptUser(c *gin.Context) {
	// res := &slack.slackevents.MessageAction{}
	// err := c.BindJSON(res)

}
