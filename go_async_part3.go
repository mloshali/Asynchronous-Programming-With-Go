// Wait Groups
// The sync package of Go provides some basic synchronization primitives such as mutexes, WaitGroups, etc.
// In this example, we would be using the WaitGroup.

// A WaitGroup waits for a collection of go routines to finish. Just simply add the number of go routines
// to wait for in the main routine. Let us use it in our example:

package main

import (
	"fmt"
	"net/http"
	"sync"
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
	var wg sync.WaitGroup

	for _, u := range urls {
		// Increment the wait group counter
		wg.Add(1)
		go func(url string) {
			// Decrement the counter when the go routine completes
			defer wg.Done()
			// Call the function check
			checkUrl(url)
		}(u)
	}
	// Wait for all the checkWebsite calls to finish
	wg.Wait()
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

// In line#25, we have defined a WaitGroup instance. Now check line#29, where we increment with 1 just before we
// execute the go routine. In line#32, we have a deferred call to the Done() method of WaitGroup. This will be a
// signal to the WaitGroup to let it know that the execution of the go routine is complete. Finally in line#38,
// we added the Wait() method call. This will let the main routine to simply wait till all the Done() calls of
// the WaitGroup have been executed. If we now execute this code snippet, we will see that we get an asynchronous
// output similar to the previous section.

// WaitGroup generally should be used for much low level designs. For a better communication, we should use channels.
// Let us jump into that in the next section.
// (See Example go_async_part4.go)
