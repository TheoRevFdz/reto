package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/TheoRevFdz/reto/model"
)

func GetByName(url string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nombre de heroe: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")
	name = strings.Replace(name, " ", "%20%", -1)

	getData(fmt.Sprintf("%s&name=%s", url, name))
}

func GetAll(url string) {
	getData(url)
}

func getData(url string) model.Marvel {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	res.Body.Close()

	marvel := model.Marvel{}
	json.Unmarshal(data, &marvel)

	fmt.Println(" ==== RESULT ====")
	fmt.Println(marvel.Copyright)
	fmt.Println(marvel.AttributionText)
	fmt.Println(fmt.Sprintf("Total: %v", marvel.Data.Total))
	// fmt.Println(marvel.Data)
	for _, d := range marvel.Data.Results {
		fmt.Println("=================================================")
		fmt.Println(fmt.Sprintf("ID: %v", d.ID))
		fmt.Println(fmt.Sprintf("Name: %s", d.Name))
		fmt.Println(fmt.Sprintf("Descripción: %s", d.Description))
		fmt.Println(fmt.Sprintf("Fecha de creacion: %s", d.Modified))
	}
	return marvel
}
