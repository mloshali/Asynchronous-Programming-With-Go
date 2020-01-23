// GoChannels

package main

import (
	"fmt"
	"net/http"
)

type urlStatus struct {
	url    string
	status bool
}

func main() {
	// A slice of sample websites
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	c := make(chan urlStatus)
	for _, url := range urls {
		go checkUrl(url, c)

	}
	result := make([]urlStatus, len(urls))
	for i, _ := range result {
		result[i] = <-c
		if result[i].status {
			fmt.Println(result[i].url, " is up and running.")
		} else {
			fmt.Println(result[i].url, " is down.")
		}
	}

}

//checks and prints a message if a website is up or down
func checkUrl(url string, c chan urlStatus) {
	_, err := http.Get(url)
	if err != nil {
		// The website is down
		c <- urlStatus{url, false}
	} else {
		// The website is up
		c <- urlStatus{url, true}
	}
}

// In line#10, we have declared a struct called urlStatus. This simply contains two fields, the url and the status if it
// is up or not. We will use this as the type for the communication in the channel. Line#43, the checkUrl function has
// been modified to also accept a channel as an input. Instead of returning a result, we will send the result as an
// instance of urlStatus to the channel. Lets jump to the main routine now. We instantiate this channel with type
// urlStatus in line#25. We will pass this channel as an input parameter in the function calls for checkUrl. We will now
// instantiate a slice result which simply will collect all the results which we will receive from the channel (Line#31).
// In each element of the result slice, we simply add the channel response (Line#33). All we have to do now is to check
// these values and print out the appropriate message.
