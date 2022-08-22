package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const fileDir = "/Users/ferhat.ozcullu/go/src/notion-english-csv-parser/"

func main() {
	var typeFlag = flag.Int("t", 1, "0 vocabulary, 1 quizlett")
	var quantityFlag = flag.Int("q", 10, "word quantity for export")
	flag.Parse()
	f, err := os.Open(fileDir + "file.csv")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	_records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return
	}
	var records [][]string

	//remove empty records
	for _, record := range _records {
		if record[1] != "" {
			records = append(records, record)
		}
	}

	//take last 10 elements
	for _, record := range records[len(records)-*quantityFlag:] {
		if record[1] == "" {
			continue
		}
		if *typeFlag == 0 {
			fmt.Println(record[1] + ",")
		} else {
			var meaning string
			response, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + record[1])
			if err != nil || response.StatusCode == 404 {
				meaning = record[2]
			} else {
				responseByte, _ := ioutil.ReadAll(response.Body)

				var res []T

				json.Unmarshal(responseByte, &res)

				wordDefination := res[0].Meanings[0].Definitions[0].Definition
				wordExample := ""
				wordType := ""

				definitions := res[0].Meanings[0].Definitions
				wordType = res[0].Meanings[0].PartOfSpeech

				for _, definition := range definitions {
					if definition.Example != "" {
						wordDefination = definition.Definition
						wordExample = definition.Example
						break
					}
				}
				meaning = fmt.Sprintf("(%s) %s\nEx: %s", wordType, wordDefination, wordExample)
			}

			fmt.Println(record[1] + "->" + meaning + "\n")
		}

	}
}

type T struct {
	Word      string `json:"word"`
	Phonetic  string `json:"phonetic"`
	Phonetics []struct {
		Text      string `json:"text"`
		Audio     string `json:"audio"`
		SourceUrl string `json:"sourceUrl"`
		License   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"license"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string        `json:"definition"`
			Synonyms   []interface{} `json:"synonyms"`
			Antonyms   []interface{} `json:"antonyms"`
			Example    string        `json:"example,omitempty"`
		} `json:"definitions"`
		Synonyms []string      `json:"synonyms"`
		Antonyms []interface{} `json:"antonyms"`
	} `json:"meanings"`
	License struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"license"`
	SourceUrls []string `json:"sourceUrls"`
}
