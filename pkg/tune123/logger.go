package tune123

import (
	"log"
	"os"
	"time"
)

/*
   logger.go - logfile writer
*/

// LogFile - logfile struct
type LogFile struct {
	Filename string // path to file
}

// Add - writes line to logfile
func (l *LogFile) Add(line string) (err error) {
	prefix := time.Now().Format("2006-01-02 15:04:05")
	file, err := os.OpenFile(l.Filename, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Printf("LogFile Add: %v\n", err)
		return err
	}

	_, err = file.WriteString(prefix + "\t" + line + "\n")
	if err != nil {
		log.Printf("LogFile WriteString: %v\n", err)
		return err
	}

	return err
}

// Init - new logfile, new logdatefile with date
// Create files if files not exists, sets GlobalLogDatFile.date
func (l *LogFile) Init() (err error) {
	//str := fmt.Sprintf("%02d-%02d-%04dT00:00:01+00:00", date.Day(), date.Month(), date.Year())
	// rewrite log-file
	err = l.MakeFile()
	if err != nil {
		return err
	}

	l.Add("LogFile initialization completed.")

	return err
}

// MakeNewFile - creates or rewrites file
func (l *LogFile) MakeNewFile() (err error) {
	file, err := os.Create(l.Filename)
	defer file.Close()
	return err
}

// MakeFile - creates file if it not exists
func (l *LogFile) MakeFile() (err error) {
	if _, err = os.Stat(l.Filename); err == nil {
		//File exist and will not be rewrited
		//fmt.Println(os.IsExist(err),err)
		return err
	}
	file, err := os.Create(l.Filename)
	defer file.Close()
	return err
}

// IsExist - check if file exist
func (l *LogFile) IsExist() (exist bool) {
	if _, err := os.Stat(l.Filename); os.IsNotExist(err) {
		return false //file not exist
	}
	return true
}
