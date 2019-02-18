package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hatobus/TDK_gakushoku/cmd/models"
	"github.com/hatobus/TDK_gakushoku/cmd/presenter"
	"github.com/hatobus/TDK_gakushoku/cmd/slackbot"
	"github.com/hatobus/TDK_gakushoku/cmd/util"
	"github.com/k0kubun/pp"
	"github.com/nlopes/slack"
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

type Payload struct {
	Payload slack.InteractionCallback `json:"payload"`
}

func GetAcceptUser(c *gin.Context) {
	res := &slack.InteractionCallback{}
	// res := &Payload{}
	// pp.Println(c.Request)

	buf, err := ioutil.ReadAll(c.Request.Body)
	log.Println(err)
	jsonstr, err := url.QueryUnescape(string(buf)[8:])
	log.Println(err)
	pp.Println(jsonstr)

	err = json.Unmarshal([]byte(jsonstr), res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	pp.Println(res)

	err = presenter.UpdateUserCoin(res.User.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
