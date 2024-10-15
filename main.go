package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type Text struct {
	Text string `json:"text"`
}
type Parts struct {
	Parts []Text `json:"parts"`
}

type Body struct {
	Contents []Parts `json:"contents"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error: %v. Error cargando las variables de entorno\n", err.Error())
	}
	apikey := os.Getenv("GEMINI_API_KEY")
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=%v", apikey)
	body := []byte(`{"contents":[{"parts":[{"text":"Que dia es hoy"}]}]}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Content-Type", "aplication/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()

	var result map[string]interface{}
	decErr := json.NewDecoder(res.Body).Decode(&result)
	if decErr != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(result)
}
