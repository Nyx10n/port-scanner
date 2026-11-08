package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func worker(ip string, portsToScan, results chan int){
	for{
		port:=<-portsToScan
		i_P:=ip+":"+strconv.Itoa(port)
			dial,err:=net.DialTimeout("tcp", i_P, time.Millisecond*200)
			if err != nil {
				results<-0
			}else{
				results<-port
				dial.Close()
				continue
			}
	}
}

func main(){
	if len(os.Args)!=2{
		fmt.Println("Usage:\n ./portscanner <IP>")
		os.Exit(1)
	}else{
		portsToScan:= make(chan int)
		results:= make(chan int)
		IP:= os.Args[1]
		fmt.Println("[+] Scanning ports on IP", IP)

		for i := 1; i < 1025; i++ {
			go worker(IP, portsToScan, results)
		}

		go func() {
			for i := 1; i < 65536; i++ {
				portsToScan<-i 
			}
		}()

		for i := 1; i < 65536; i++ {
			stat:=<-results
			if stat!=0 {
				fmt.Println("Port", stat, "is open")
			}else{
				continue
			}
		}
	}
}