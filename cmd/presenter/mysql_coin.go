package presenter

import (
	"log"
	"os"

	"github.com/hatobus/TDK_gakushoku/cmd/models"
)

func UpdateUserCoin(username string) error {
	engine, err := SetUpEngine(os.Getenv("SOURCE"))
	if err != nil {
		log.Println(err)
		return err
	}

	usr := &models.Student{Name: username}
	has, err := engine.Get(&usr)
	if err != nil {
		log.Println("Get from username error : ", err)
		return err
	} else if !has {
		usr.Sumofcoin = 1
		_, err = engine.Insert(usr)
		if err != nil {
			log.Println("New user insert error : ", err)
			return err
		}
		return nil
	}

	usr.Sumofcoin += 1
	_, err = engine.ID(usr.No).Update(usr)
	if err != nil {
		log.Println("Coin update error : ", err)
		return err
	}

	return nil
}
