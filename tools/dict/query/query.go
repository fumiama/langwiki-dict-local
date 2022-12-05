package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fumiama/langwiki-dict-local/db"
	q "github.com/fumiama/langwiki-dict-local/query"
)

var isexit = false

// setupMainSignalHandler is for main to do cleanup
func setupMainSignalHandler() {
	mc := make(chan os.Signal, 2)
	signal.Notify(mc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-mc
		isexit = true
	}()
}

// curl 'https://langwiki.org/tools/dict/php/query.php' \
//   -H 'Accept: application/json, text/javascript, */*; q=0.01' \
//   -H 'Accept-Language: zh,zh-CN;q=0.9,zh-HK;q=0.8,zh-TW;q=0.7,ja;q=0.6,en;q=0.5,en-GB;q=0.4,en-US;q=0.3' \
//   -H 'Connection: keep-alive' \
//   -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' \
//   -H 'Cookie: _ga=GA1.2.382918067.1670138716; _gat=1' \
//   -H 'Origin: https://langwiki.org' \
//   -H 'Referer: https://langwiki.org/tools/dict/' \
//   -H 'Sec-Fetch-Dest: empty' \
//   -H 'Sec-Fetch-Mode: cors' \
//   -H 'Sec-Fetch-Site: same-origin' \
//   -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.62' \
//   -H 'X-Requested-With: XMLHttpRequest' \
//   -H 'sec-ch-ua: "Microsoft Edge";v="107", "Chromium";v="107", "Not=A?Brand";v="24"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'sec-ch-ua-platform: "macOS"' \
//   --data-raw 'string=%E5%95%8A&mode=0&flag%5B%5D=0&flag%5B%5D=1&flag%5B%5D=1&setting%5B%5D=0&setting%5B%5D=4&setting%5B%5D=0&setting%5B%5D=0&setting%5B%5D=0&setting%5B%5D=0&setting%5B%5D=1&setting%5B%5D=0&setting%5B%5D=0&setting%5B%5D=0&bot=0' \
//   --compressed

func main() {
	s := flag.String("s", "一", "start char")
	flag.Parse()
	b, err := db.NewDB(flag.Args()[0], time.Second)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer b.Close()
	setupMainSignalHandler()
	for c := []rune(*s)[0]; c <= 0x9fa5 && !isexit; c++ {
		fmt.Println("query: ", string(c))
		data, err := query(string(c), [4]uint8{0, 0, 1, 1}, [10]uint8{0, 4, 0, 0, 0, 0, 1, 0, 0, 0})
		if err != nil {
			fmt.Println("ERROR:", err)
			c--
			time.Sleep(time.Second + time.Duration(rand.Intn(1000))*time.Millisecond)
			continue
		}
		content, err := q.NewContent([]byte(data))
		if err == q.ErrZeroLang {
			fmt.Println("未入库的字:", &content)
			continue
		}
		if err != nil {
			fmt.Println("ERROR:", err)
			c--
			time.Sleep(time.Second + time.Duration(rand.Intn(1000))*time.Millisecond)
			continue
		}
		simis := make([]string, len(content.Cards))
		for i, card := range content.Cards {
			simis[i] = string(card.Char)
			for _, lang := range card.Langs {
				err = b.AddChar(lang.Type, &db.Lang{Char: card.Char, Info: lang.Info})
				if err != nil {
					fmt.Println("ERROR:", err)
					return
				}
			}
		}
		if len(simis) > 1 {
			err = b.AddSimilar(content.Cards[0].Char, simis[1:]...)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
		}
		fmt.Println("add:", &content)
	}
}

func query(chr string, modeflags [4]uint8, settings [10]uint8) (string, error) {
	body := bytes.NewBuffer([]byte("string="))
	body.WriteString(url.QueryEscape(chr))
	body.WriteString("&mode=")
	body.WriteByte('0' + modeflags[0])
	body.WriteString("&flag%5B%5D=")
	body.WriteByte('0' + modeflags[1])
	body.WriteString("&flag%5B%5D=")
	body.WriteByte('0' + modeflags[2])
	body.WriteString("&flag%5B%5D=")
	body.WriteByte('0' + modeflags[3])
	for _, s := range settings {
		body.WriteString("&setting%5B%5D=")
		body.WriteByte('0' + s)
	}
	body.WriteString("&bot=0")
	req, err := http.NewRequest("POST", "https://langwiki.org/tools/dict/php/query.php", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://langwiki.org")
	req.Header.Set("Referer", "https://langwiki.org/tools/dict/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.62")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var rsp struct {
		Status string `json:"status"`
		HTML   string
	}
	err = json.NewDecoder(resp.Body).Decode(&rsp)
	if err != nil {
		return "", err
	}
	if rsp.Status != "success" {
		return "", errors.New(rsp.Status)
	}
	return rsp.HTML, nil
}
