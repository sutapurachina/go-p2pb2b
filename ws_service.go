package p2pb2b

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strings"
	"time"
)

type LastPriceResponse struct {
	Method string      `json:"method"`
	Params []string    `json:"params"`
	Id     interface{} `json:"id"`
}

type LastPriceInfo struct {
	Method string      `json:"method"`
	Params []string    `json:"params"`
	Id     interface{} `json:"id"`
}

func LastPriceStream(symbols ...string) (chan *LastPriceInfo, chan struct{}, error) {
	const method = "price.subscribe"
	c, _, err := websocket.DefaultDialer.Dial(websocketApi, nil)
	if err != nil {
		return nil, nil, err
	}
	c.SetReadLimit(655535000000)

	jsonData, err := json.Marshal(newWsRequest(method, symbols...))
	err = c.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		return nil, nil, err
	}
	_, _, err = c.ReadMessage()
	if err != nil {
		return nil, nil, err
	}
	ksStop := keepAlive(c, 20*time.Second)
	stopC := make(chan struct{}, 1)
	lastPriceC := make(chan *LastPriceInfo)
	go func() {
		for {
			select {
			case <-stopC:
				fmt.Println("here")
				jsonData, err := json.Marshal(newUnsubscribeRequest("price"))
				err = c.WriteMessage(websocket.TextMessage, jsonData)
				if err != nil {
					return
				}
				_, _, err = c.ReadMessage()
				if err != nil {
					return
				}
				ksStop <- struct{}{}
				return
			default:
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Printf("LastPriceStream: can't read message`: %v\n", err)
					if strings.Contains(err.Error(), "clos") {
						break
					}
					continue
				}
				if strings.Contains(string(message), "pong") {
					break
				}
				res := &LastPriceInfo{}
				err = json.Unmarshal(message, res)
				if err != nil {
					log.Printf("LastPriceStream: can't unmarshal`: %v\n", err)
					break
				}
				go func() {
					lastPriceC <- res
				}()
			}
		}
	}()

	return lastPriceC, stopC, nil
}

func keepAlive(c *websocket.Conn, timeout time.Duration) chan struct{} {
	ticker := time.NewTicker(timeout)
	/*c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})*/
	stopC := make(chan struct{})
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				jsonData, err := json.Marshal(newPingRequest())
				if err != nil {
					log.Printf("keepalive: %v/n", err)
				}
				err = c.WriteMessage(websocket.TextMessage, jsonData)
				if err != nil {
					log.Printf("keepalive: %v/n", err)
					return
				}
			case <-stopC:
				return
			}
		}
	}()
	return stopC
}
