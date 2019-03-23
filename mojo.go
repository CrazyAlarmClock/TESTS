package main


import (
	"fmt"      // пакет для форматированного ввода вывода
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола
	"strings"  // пакет для работы с  UTF-8 строками
	"time"
)


var p = fmt.Println
var s2 string

func main() {
	timer()
}

func timer()  {
	ticker := time.NewTicker(time.Millisecond * 9900)
	parse()

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1000)
	timer()
}

func parse()  {url := "http://player.ru/forumdisplay.php?f=6"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP error:", err)
		return
	}

	defer resp.Body.Close()
	utf8, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println("Encoding error:", err)
		return
	}
	body, err := ioutil.ReadAll(utf8)
	if err != nil {
		fmt.Println("IO error:", err)
		return
	}

	s := string(body[:])
	changeValue(&s)

	if s2 == s {
		fmt.Println("noting...")
	} else {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println(s)
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!")
		s2 = s
		bot()
	}
}

func changeValue(s *string){
	i:= 0
	var stroka string
	k := strings.Split(*s, "\n")

	for  i<1000{

		//arr1 = append(arr1, k[i])
		if  strings.Contains(k[i],"Chord mojo") == true{
			stroka = k[i]
		} else {}
		i =i+1
	}
	*s = stroka
}

func bot()  {
	bot, err := tgbotapi.NewBotAPI("756272327:AAHrBbK0TOs39to9DQ30BKjfkRMMyALLVtE")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(1)
	u.Timeout = 0
	msg := tgbotapi.NewMessage(196191173,"1")
	msg.Text = s2
	bot.Send(msg)
}
