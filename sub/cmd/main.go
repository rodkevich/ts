package main

import (
	"net/http"
	"sync"
	"time"
)

var (
	RegisteredServiceStorage map[string]time.Time
	ServiceStorageMutex      sync.RWMutex
)

func main() {

	http.HandleFunc("/registerAndKeepAlive", RegisterAndKeepAlive)
	http.HandleFunc("/unregister", RemoveRegistered)
	http.HandleFunc("/sendMessage", HandleMessage)
	http.HandleFunc("/listSubscribers", HandleSubscriberListing)

	go KillZombies()
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		panic(err)
	}
}

func RegisterAndKeepAlive(writer http.ResponseWriter, request *http.Request) {

}

func RemoveRegistered(writer http.ResponseWriter, request *http.Request) {

}

func HandleMessage(writer http.ResponseWriter, request *http.Request) {

}

func HandleSubscriberListing(writer http.ResponseWriter, request *http.Request) {

}

func KillZombies() {
	// set interval
	t := time.Tick(1 * time.Minute)
	// check if a service should no longer get any messages
	for range t {
		now := time.Now()
		ServiceStorageMutex.Lock()
		for receiver, timeKeepAlive := range RegisteredServiceStorage {
			if now.Sub(timeKeepAlive).Minutes() > 2 {
				delete(RegisteredServiceStorage, receiver)
			}
		}
		ServiceStorageMutex.Unlock()
	}
}
