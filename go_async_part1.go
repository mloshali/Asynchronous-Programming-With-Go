// 1. The Synchronous Way
// The most easiest way to handle this would be to simple iterate over your slice of urls and try to make an HTTP
// GET request to the url. Check out the following code snippet:

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// A slice of sample websites
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	for _, url := range urls {
		checkUrl(url)
	}
}

//checks and prints a message if a website is up or down
func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, " is down.")
		return
	}
	fmt.Println(url, " is up and running.")
}

// The code snippet is fairly basic and when I run it, I get the following output:

// https://www.easyjet.com/ is up and running.
// https://www.skyscanner.de/ is up and running.
// https://www.ryanair.com is up and running.
// https://wizzair.com/ is up and running.
// https://www.swiss.com/ is up and running.

// This is good, but in this case we check if each website is up one by one.
// An ideal solution would be to handle this non-linearly and check if they are up concurrently.
// (See Example go_async_part2.go)
