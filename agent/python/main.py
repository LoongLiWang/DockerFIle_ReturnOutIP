#!/usr/bin/env python3

import requests

def main():
    url = "http://ip.wang-li.top:93/4u6385IP"
    MyIP = requests.get(url).text
    print(MyIP)

if __name__ == '__main__':
    main()