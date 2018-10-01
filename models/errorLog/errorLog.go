package errorLog

import (
	"log"
	"os"
)

//Write writes an error log message to the log at /cashFlowApp/serverLogs/errorLog.txt and prints it out to the terminal
func Write(m string, e error) {
	f, err := os.OpenFile("./serverLogs/errorLog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	var message string
	message = m + e.Error()

	log.Println(message)
	log.SetOutput(f)
	log.Println(message)
}
