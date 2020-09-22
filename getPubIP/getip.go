package getPubIP

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
)

func checkErr(e error) {
	if e != nil {
		fmt.Fprintln(os.Stdout, e)
	}
}

// GetIP 获取公网ip模块
func GetIP() (ip net.IP, err error) {
	url := "https://202020.ip138.com/"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36")
	req.Header.Set("Accept-Charset", "gb2312")
	resp, err := client.Do(req)
	checkErr(err)
	if resp.StatusCode != 200 {
		log.Fatalln("error request")
	}

	scanner := bufio.NewScanner(resp.Body)
	ipf := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	var adr string
	for scanner.Scan() {
		ipArr := ipf.FindAllString(scanner.Text(), -1)
		if len(ipArr) == 1 {
			adr = ipArr[0]
		} else {
			fmt.Fprintln(os.Stdout, "not found IP in ip138.com")
			return ip, nil
		}
	}

	return net.ParseIP(adr), nil

}
