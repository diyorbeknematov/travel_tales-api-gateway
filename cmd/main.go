package main

import (
	"api-gateway/api"
	"api-gateway/api/handler"
	"api-gateway/config"
	"api-gateway/logs"
	"api-gateway/service"
	"log"
	"log/slog"
)

func main() {
	cfg := config.Load()
	logs.InitLogger()

	serviceManager, err := service.NewServiceManager(&cfg)
	if err != nil {
		logs.Logger.Error("gRPC dial error", slog.String("error", err.Error()))
	}

	handler := handler.NewHandler(serviceManager, logs.Logger)
	router := api.NewRouter(handler)
	logs.Logger.Info("Server is running ... ", slog.String("PORT", cfg.HTTP_PORT))

	err = router.Run(cfg.HTTP_PORT)
	if err != nil {
		logs.Logger.Error("Routerni run qilishda xatolik beryapti", slog.String("error", err.Error()))
		log.Fatal(err)
	}
}
