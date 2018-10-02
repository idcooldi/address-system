package main

import (
	//"bufio"
	"encoding/xml"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Address struct {
	Short string
	Name  string
}

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:secret@tcp(db:3306)/xml")
	if err != nil {
		log.Printf("Ошибка подключения к DB %v\n", err)
		panic(err)
	}
	defer func() {
		if err = db.Close(); err != nil {
			panic(err)
		}
	}()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&Address{})
	parse()
}
func parse() {
	file, err := os.Open("example.xml")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	parser := xml.NewDecoder(file)
	var address Address
	for {
		token, err := parser.Token()
		if err != nil {			
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			elmt := xml.StartElement(t)
			for _, a := range elmt.Attr {
				switch a.Name.Local {
				case "SHORTNAME":
					var short string
					switch a.Value {
					case "б-р":
						short = "бульвар"
					case "дор":
						short = "дорога"
					case "пл":
						short = "площадь"
					case "пр-д":
						short = "проезд"
					case "ул":
						short = "улица"
					case "пер":
						short = "переулок"
					case "г":
						short = "город"
					case "д":
						short = "деревня"
					case "п":
						short = "посёлок"
					case "с":
						short = "село"
					case "ш":
						short = "шоссе"
					case "тер":
						short = "территория"
					case "мкр":
						short = "микрорайон"
					case "кв-л":
						short = "квартал"
					default:
						short = a.Value
					}
					address.Short = short
				case "FORMALNAME":
					address.Name = a.Value
				default:
					log.Printf("Name:%v  Value:%v\n", a.Name.Local, a.Value)
				}
			}
			log.Printf("PRODUCT %v %v\n", address.Short, address.Name)
			db.Create(&address)
		}
	}
}
