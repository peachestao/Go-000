package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func hello(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("hello golang"))
}

func main(){

	g,_:=errgroup.WithContext(context.Background())

	s:=http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/hello",hello)

	g.Go(func() error {
		return s.ListenAndServe()
	})

	g.Go(func() error {
		c := make(chan os.Signal)
		signal.Notify(c)
		for {
			_= <-c
			os.Exit(1)
		}
		return nil
	})

	err:=g.Wait()
	fmt.Println(err)
	if err!=nil{
	}
}
