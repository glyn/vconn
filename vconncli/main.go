package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	for i := 0; i < 20; i++ {
		go func(index int) {

			res, err := http.Get("http://localhost:8888")
			if err != nil {
				log.Fatal(err)
			}
			result, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%d: %s\n", index, result)
		}(i)
	}

	select {}
}
