package health

import (
	. "github.com/paulohsl/api-health"
	"fmt"
)

type Caller interface {
	Get(url string, ch chan<- HttpResponse)
}

func Check(caller Caller, urls []string) {

	ch := make(chan HttpResponse)

	for _, url := range urls {
		go caller.Get(url, ch)
	}

	for _, url := range urls {
		fmt.Printf("Url: %s -> %s \n", url, (<-ch).HttpStatus)
	}

	fmt.Println((<-ch).HttpStatus)
}
