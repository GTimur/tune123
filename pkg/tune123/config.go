package tune123

import (
	"log"
)

// Config - структура файла конфигурации
type Config struct {
	AudioPath    []string //пути к папкам, где находятся файлы аудио
	PlaylistPath string   //путь к папке где будут храниться плейлисты
	InitVolume   int      //громкость при первом старте программы (в процентах)
	InitPlay     bool     //Включать воспроизведение после запуска
	FileName     string   //Имя файла для выгрузки в формате JSON
	Log          LogFile  //Вложенная структура для реализации лог файла

}

func (c *Config) WriteJSON() error {
	err := WriteJSON(c, c.FileName)
	if err != nil {
		log.Println("WriteJSON [", c.FileName, "]:", err)
	}
	return err
}

func (c *Config) ReadJSON() error {
	err := ReadJSON(c, c.FileName)
	if err != nil {
		log.Println("ReadJSON [", c.FileName, "]:", err)
	}
	return err
}

func (c *Config) CreateJSONFile() error {
	err := CreateJSONFile(c, c.FileName)
	if err != nil {
		log.Println("CreateJSONFile [", c.FileName, "]:", err)
	}
	return err
}
