package presenter

import (
	"os"

	"github.com/hatobus/TDK_gakushoku/cmd/models"
)

func GetTop10() ([]models.Student, error) {
	engine, err := SetUpEngine(os.Getenv("SOURCE"))
	if err != nil {
		return nil, err
	}

	var top10 = make([]models.Student, 0, 10)
	err = engine.Desc("sumofcoin").Limit(10, 0).Find(&top10)
	if err != nil {
		return nil, err
	}

	return top10, nil
}
