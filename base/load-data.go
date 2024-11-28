package base

import (
	"bufio"
	"os"
	"strings"
	"ujiTest/models"
)

func (s *Server) LoadDataFromFile(filePath string) ([]models.Item, error) {
	var items []models.Item
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) == 5 {
			item := models.Item{
				Code:   fields[0],
				Name:   fields[1],
				Model:  fields[2],
				Tech:   strings.Split(fields[3], "|"),
				Status: fields[4],
			}
			items = append(items, item)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return items, nil
}
