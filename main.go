package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TheoRevFdz/reto/services"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	const pubKey = "bed2e3a8f2d65285bff7810dc68b7b08"
	const privKey = "299782ed5f6ee70e13951e761cfd15c1ee7f759e"
	var ts = strconv.FormatInt(time.Now().Unix(), 10)

	hasher := md5.New()
	hasher.Write([]byte(ts + privKey + pubKey))
	hash := hex.EncodeToString(hasher.Sum(nil))
	// hash = "0c88e2f215377943c7c8a48b3ff8fa3d"

	url := fmt.Sprintf("http://gateway.marvel.com/v1/public/characters?ts=%s&apikey=%s&hash=%s", ts, pubKey, hash)

	fmt.Println("1. Buscar por nombre")
	fmt.Println("2. Listar")

	fmt.Print("Ingrese una opcion: ")
	opt, _ := reader.ReadString('\n')
	opt = strings.TrimRight(opt, "\r\n")

	// fmt.Println("\n--------------------------------------")
	switch opt {
	case "1":
		services.GetByName(url)
	case "2":
		services.GetAll(url)
	default:
		fmt.Println(fmt.Sprintf("Error: |%s|", opt))
	}

}
