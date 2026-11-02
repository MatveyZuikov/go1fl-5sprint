package actioninfo

import (
	"fmt"
	"log"
) 

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {

	if len(dataset) == 0 {
		fmt.Println("нет данных")
		return
	}

	for _, data := range dataset {

		err := dp.Parse(data)

		if err != nil {
            log.Printf("ошибка парсинга строки %q: %v", data, err)
            continue 
        }

		info, err := dp.ActionInfo()

		if err != nil {
            log.Printf("ошибка формирования отчёта: %v", err)
            continue
        }

		fmt.Println(info)
	} 
}
