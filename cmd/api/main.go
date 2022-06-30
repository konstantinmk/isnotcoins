package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/konstantinmk/isnotcoins/logging"
)

func main() {
	current_path, err := os.Getwd()
	if err != nil {
		panic("Unsupported error getting current working directory")		
	}
	//Создание пути файла для лога результата
	log_path:= filepath.Join(current_path,"logs","log.txt")
	//Определение логеров	
	logger,err:=logging.InitFileLogger("main",log_path)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Close program.")		
		os.Exit(1)
	}
	logger.Info().Msg("Start API")

	logger.Info().Msg("Stop API")
}
