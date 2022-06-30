// Logging package
package logging

import (	
	"os"
	"path/filepath"
	"time"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init(){
	// Set logging level
	zerolog.SetGlobalLevel(zerolog.DebugLevel) 	
}

//Console logger init
func InitConsoleLogger(service string) (*zerolog.Logger, error){
	//Создание логера для вывода в консоль.	
	logger := log.With().Str("service", service).Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123 })
	return &logger, nil	
}

//File logger init
func InitFileLogger(service string, log_path string) (*zerolog.Logger, error ){	
	//Check log filepath exist
	dir_path := filepath.Dir(log_path)
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {		
		err := os.Mkdir(dir_path, 0777)
		//Create file path if not exist
		if err!=nil {			
			return nil, err
		}
	}
	file, err := os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {		
		return nil, err
    }
	logger := log.With().Str("service", service).Logger().Output(file)
	return &logger, nil
}

