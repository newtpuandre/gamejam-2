package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Highscore struct {
	Highscore int64
}

var highscoreStruct Highscore

func ConfigInit() {

	if _, err := os.Stat("./highscore.json"); os.IsNotExist(err) {
		log.Println("highscore.json did not exist and have been created. Please fill in the fields")
		file, _ := json.MarshalIndent(highscoreStruct, "", " ")

		_ = ioutil.WriteFile("./highscore.json", file, 0644)

	}

	file, err := os.Open("./highscore.json")
	if err != nil {
		log.Println(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&highscoreStruct)
	if err != nil {
		log.Println(err)
	}

	highscore = highscoreStruct.Highscore

}

func writeHighscore() {
	highscoreStruct.Highscore = highscore
	file, _ := json.MarshalIndent(highscoreStruct, "", " ")

	_ = ioutil.WriteFile("highscore.json", file, 0644)
}
