package utils

import "os"


// work better on log idea...
// just some drafts..


type LogIO struct {
	IO *os.File
}

func InitLogIO() (*LogIO, error) {
	log, err := os.Create("server1.log")
	if err != nil {
		return &LogIO{}, err
	}
	return &LogIO{IO: log}, nil
}

func (log *LogIO) WriteRequestLog(path, requestBody string) error {
	
	_, err := log.IO.Write([]byte(path))
	if err != nil {
		return err
	} else {
		_, err = log.IO.Write([]byte(requestBody))
		if err != nil {
			return err
		}
	}

	return nil
}

func (log *LogIO) WriteResponseLog(path, response, statusCode string) error {
	_, err := log.IO.Write([]byte(path))
	if err != nil {
		return err
	} else {
		_, err = log.IO.Write([]byte(response))
		if err != nil {
			return err
		}
	}

	return nil
}
