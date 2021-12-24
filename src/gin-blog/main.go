package main

import (
	"fmt"
	"log"
	"net/http"

	"xing.learning.gin/src/gin-blog/pkg/setting"
	routers "xing.learning.gin/src/gin-blog/routers"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
