package main

import (
	"fmt"
	"log"
	"net/http"

	"xing.learning.gin/src/gin-blog/models"
	"xing.learning.gin/src/gin-blog/pkg/logging"
	"xing.learning.gin/src/gin-blog/pkg/setting"
	routers "xing.learning.gin/src/gin-blog/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
