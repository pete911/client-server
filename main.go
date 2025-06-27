package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	flags, err := ParseFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("flags set to: %+v\n", flags)

	// start server
	go server(flags.Port)
	time.Sleep(1 * time.Second)

	// start sending requests
	sendRequests(flags)
}

func sendRequests(flags Flags) {
	client := NewClient(flags)
	requests := atomic.Int64{}
	forever := make(chan struct{})

	netstatAllCmd := fmt.Sprintf("netstat -n -p tcp | grep %d | wc -l", flags.Port)
	netstatTimeWaitCmd := fmt.Sprintf("netstat -n -p tcp | grep %d | grep TIME_WAIT | wc -l", flags.Port)
	netstatCloseWaitCmd := fmt.Sprintf("netstat -n -p tcp | grep %d | grep CLOSE_WAIT | wc -l", flags.Port)
	netstatEstablishedCmd := fmt.Sprintf("netstat -n -p tcp | grep %d | grep ESTABLISHED | wc -l", flags.Port)

	// every 2 seconds print how many requests we have made and additional info
	fmt.Println("requests\tconnections\tTIME_WAIT\tCLOSE_WAIT\tESTABLISHED")
	go func() {
		for {
			time.Sleep(2 * time.Second)
			numOfRequests := requests.Load()
			allConnections := execCmd("bash", "-c", netstatAllCmd)
			timeWaitConnections := execCmd("bash", "-c", netstatTimeWaitCmd)
			closeWaitConnections := execCmd("bash", "-c", netstatCloseWaitCmd)
			establishedConnections := execCmd("bash", "-c", netstatEstablishedCmd)
			fmt.Printf("%8d\t%11.11s\t%9.9s\t%10.10s\t%11.11s\n", numOfRequests, allConnections, timeWaitConnections, closeWaitConnections, establishedConnections)
		}
	}()

	url := fmt.Sprintf("http://localhost:%d/message", flags.Port)
	for i := 0; i < flags.Concurrency; i++ {
		go func() {
			for {
				time.Sleep(time.Millisecond)

				req, err := http.NewRequest(http.MethodGet, url, nil)
				res, err := client.Do(req)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if flags.ReadBody {
					io.Copy(io.Discard, res.Body)
				}
				if flags.CloseBody {
					res.Body.Close()
				}

				requests.Add(1)
			}
		}()
	}

	<-forever
}
