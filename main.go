package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type PlayerPPPlusData struct {
	username  string
	pp        string
	aim_jump  string
	aim_flow  string
	precision string
	speed     string
	stamina   string
	accuracy  string
	total     string
}

func getPPPlusDoucment(username string) *goquery.Document {
	res, err := http.Get("https://syrin.me/pp+/u/" + username)
	if err != nil {
		fmt.Print(err)
	}

	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	return doc
}

func removeCommaAndUnit(value string) string {
	valueTrimSuffix := strings.TrimSuffix(value, "pp")

	res := strings.ReplaceAll(valueTrimSuffix, ",", "")
	return res
}

func getValueFromSelector(document *goquery.Document, selector string) string {
	return removeCommaAndUnit(document.Find(selector).Text())
}

func readUserNameList() []string {
	fmt.Println("Reading Player List...")
	fmt.Println()

	var userNameList []string

	file, err := os.Open("./players.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
		userNameList = append(userNameList, string(a))
	}

	return userNameList
}

func _getPlayerPPPlusDataWithChannel(userName string, channel chan *PlayerPPPlusData) {
	doc := getPPPlusDoucment(userName)
	var aim_jump = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(2) > td:nth-child(2)")
	var aim_flow = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(3) > td:nth-child(2)")
	var precision = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(4) > td:nth-child(2)")
	var speed = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(5) > td:nth-child(2)")
	var stamina = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(6) > td:nth-child(2)")
	var accuracy = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(7) > td:nth-child(2)")
	var total = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(8) > td:nth-child(2)")
	var pp = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > thead > tr > th:nth-child(2)")

	channel <- &PlayerPPPlusData{userName, pp, aim_jump, aim_flow, precision, speed, stamina, accuracy, total}
}

func getPlayerPPPlusData(userName string) *PlayerPPPlusData {
	doc := getPPPlusDoucment(userName)
	var aim_jump = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(2) > td:nth-child(2)")
	var aim_flow = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(3) > td:nth-child(2)")
	var precision = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(4) > td:nth-child(2)")
	var speed = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(5) > td:nth-child(2)")
	var stamina = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(6) > td:nth-child(2)")
	var accuracy = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(7) > td:nth-child(2)")
	var total = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > tbody > tr:nth-child(8) > td:nth-child(2)")
	var pp = getValueFromSelector(doc, "body > div > div:nth-child(8) > div:nth-child(2) > div.col-sm-7 > div > div > div.col-sm-7 > div > table > thead > tr > th:nth-child(2)")

	return &PlayerPPPlusData{userName, pp, aim_jump, aim_flow, precision, speed, stamina, accuracy, total}
}

func getPlayerPPPlusDataList(userNameList []string) []*PlayerPPPlusData {

	fmt.Println()
	fmt.Println("Getting Player PP Plus Data...")

	var res []*PlayerPPPlusData

	// use goroutine
	// channel := make(chan *PlayerPPPlusData)

	// for i := 0; i < len(userNameList); i++ {
	// 	go getPlayerPPPlusData(userNameList[i], channel)
	// }

	// for i := 0; i < len(userNameList); i++ {
	// 	player := <-channel
	// 	fmt.Println(player)
	// 	res = append(res, player)
	// }

	// sync code
	for i := 0; i < len(userNameList); i++ {
		player := getPlayerPPPlusData(userNameList[i])
		fmt.Println(player)
		res = append(res, player)
	}

	return res
}

func writeResultToFile(playDataList []*PlayerPPPlusData) {

	now := time.Now().Format("2006-01-02-15-03-04")

	filename := "pp-plus-" + now

	fmt.Println()
	fmt.Println("Writing Data To File " + filename + "...")

	file, err := os.Create("./" + filename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < len(playDataList); i++ {

		writer.WriteString("Player: " + playDataList[i].username + "\n")
		writer.WriteString("pp: " + playDataList[i].pp + "\n")
		writer.WriteString("Total: " + playDataList[i].total + "\n")
		writer.WriteString("Aim Jump: " + playDataList[i].aim_jump + "\n")
		writer.WriteString("Aim Flow: " + playDataList[i].aim_flow + "\n")
		writer.WriteString("Precision: " + playDataList[i].precision + "\n")
		writer.WriteString("Speed: " + playDataList[i].speed + "\n")
		writer.WriteString("Stanima: " + playDataList[i].stamina + "\n")
		writer.WriteString("Accuracy: " + playDataList[i].accuracy + "\n\n")
	}

	writer.Flush()

	fmt.Println("Write File " + filename + " Done.")
}

func main() {

	var startTime = time.Now().UnixMilli()

	userNameList := readUserNameList()
	playerPPPlusDataList := getPlayerPPPlusDataList(userNameList)

	// for i := 0; i < len(playerPPPlusDataList); i++ {
	// 	fmt.Println(playerPPPlusDataList[i])
	// }

	writeResultToFile(playerPPPlusDataList)

	var endTime = time.Now().UnixMilli()

	fmt.Println()
	fmt.Println(fmt.Sprint(endTime-startTime) + "ms cost")
}
