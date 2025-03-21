package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/LOOK-MOM-I-CAN-FLY/StandartWebServer/internal/app/api"
)

var (
	configPath string = "configs/api.toml"
)

func init() {
	//Скажем что наше приложение будет на этапе запуска получать путь то config файла из внешнего мира
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

// TODO добавить всё в Makefile
// TODO сделать чтобы могли ещё запускать с помощью .env файла а если нет и того и другого то запускаем с дефолтными значениями (в нашем проекта это api.toml), если и его нет то запускаем со значениями из структуры Config
// Добавить в код необходимые блоки для того чтобы можно было запускать приложение следующими командами
// api - format .env -path configs/api.env
// api - format .toml -path configs/api.toml
func main() {
	//В этот момент происходит инициализация перемнной configPath значением
	flag.Parse()
	log.Println("It's API")
	//Server instance initialization
	config := api.NewConfig()
	//Теперь тут надо попробовать прочитать из .toml/env так как там может быть новая информация
	//короче создаём сначала конфиг с дефолтными значениями, а потом считываем из .toml/env авось что-то и новое докинем
	_, err := toml.DecodeFile(configPath, config) //десеарилизуем содержимое .toml в config
	if err != nil {
		log.Println("Can't fing config file. Using default valuses", err)
	}
	server := api.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

//Самый важный пункт про конфигурацию
//Правило: В го принято:
// 1 Конфигурационные файлы лежат в сторонних файла (.toml, .env, .dockerenv)
// 2 В го проектах ВСЕГДА присутствуют дефолтные настройки (исключение: БД не дефолтится !!!)
