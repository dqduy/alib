package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

//var url = "https://hacker-news.firebaseio.com/v0/maxitem.json"
var url = "https://hacker-news.firebaseio.com/v0/item/"
var start = 1000
var neededItem = 100000
var threads = 100

func fetch(count int, wg *sync.WaitGroup) {
	for i := 10; i < count; i++ {
		response, err := http.Get(url + strconv.Itoa(start+i) + ".json")

		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", data)

	}
	wg.Done()
}

func main() {
	var startTime = time.Now()
	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go fetch(neededItem/threads, &wg)
	}
	wg.Wait()

	fmt.Println("Elapsed time: ", time.Since(startTime).Seconds())

}
