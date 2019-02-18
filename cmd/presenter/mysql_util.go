package presenter

import (
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/hatobus/TDK_gakushoku/cmd/models"
)

func InitEngine(source string) (bool, error) {
	RetryCnt := 10
	var i int
	var err error

	for i = 0; i <= RetryCnt; i++ {
		_, err := xorm.NewEngine("mysql", source)
		if err != nil {
			log.Printf("Retry: Cannot connect database %v times: %v \n", i, err.Error())
			time.Sleep(time.Second * 1)
			continue
		}
		break
	}

	if i == RetryCnt {
		return false, err
	}

	return true, nil
}

func SetUpEngine(source string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", source)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return engine, nil
}

func InsertDummyUser() error {
	e, _ := SetUpEngine(os.Getenv("SOURCE"))

	var users = make([]models.Student, 0, 0)
	name := []string{"ando", "gpioblink", "shiho", "bus", "HKato", "haga", "hato", "nozo", "noah", "eta"}
	coin := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for i, n := range name {
		users = append(users, models.Student{Name: n, Sumofcoin: coin[i], Lastworked: time.Now()})
	}

	_, err := e.Insert(&users)
	if err != nil {
		return err
	}

	return nil
}
