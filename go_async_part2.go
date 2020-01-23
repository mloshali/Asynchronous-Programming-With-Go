// Go Routines
// Go ships with it go routines, which executes a function asynchronously. It is a lightweight thread of execution.
// The keyword for it is go. Lets say, you have a function f() and you would like to execute that in a go routine,
// then simply use, go f(). So in the above example, all we need to do is simply add the keyword go before the
// function call for checkUrl, and the function would be executed asynchronously in a go routine.

package main

import (
	"fmt"
	"net/http"
	"time"
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
		// Call the function check
		go checkUrl(url)
	}
	time.Sleep(5 * time.Second)
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

// Now if we execute the code, we get the following output:

// https://www.ryanair.com is up and running.
// https://www.swiss.com/ is up and running.
// https://www.skyscanner.de/ is up and running.
// https://wizzair.com/ is up and running.
// https://www.easyjet.com/ is up and running.

// Notice, that the order of the output is different than that of our urls slice. This is because they have been
// executed asynchronously. Now, having a sleep function is not a good solution. Please note that when you execute
// this, you might have a different order in your output.

// In the next section, we will talk about a primitive way to handle this with Wait Groups.
// (See Example  go_async_part3.go)
