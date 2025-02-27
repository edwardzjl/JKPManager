package controllers

import (
	"log"
	"net/http"

	"zjuici.com/tablegpt/jkpmanager/src/common"
	"zjuici.com/tablegpt/jkpmanager/src/models"
	"zjuici.com/tablegpt/jkpmanager/src/storage"
)

func PopKernelHandler(cfg *models.Config, httpClient *common.HTTPClient, redisClient *storage.RedisClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		poppedKernels, err := redisClient.BRPop(cfg.RedisKey)
		log.Println("poppedKernels:", poppedKernels[1])

		if err != nil {
			log.Printf("Cannot pop the kernel from redis. error %v", err)
			http.Error(w, "Cannot pop the kernel", http.StatusInternalServerError)
			return
		}

		go common.StartKernels(cfg, httpClient, redisClient, 1)

		w.Write([]byte(poppedKernels[1]))
	}
}
