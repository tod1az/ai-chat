package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading enviroment variables!")
	}
	ctx := context.Background()
	client, err := Initialize(ctx)
	if err != nil {
		fmt.Printf("Error initializing: %v\n", err.Error())
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-flash")
	userPromt := ""
	for userPromt != "salir" {
		fmt.Println("Pregunta algo ( escriba salir para terminar )")
		prompt := GetPromt()
		if prompt == "salir" {
			fmt.Println("chau chau")
			os.Exit(1)
		}
		fmt.Printf("Usted: %v\n", prompt)
		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			fmt.Printf("Error getting a response: %v\n", err.Error())
		}
		fmt.Printf("Ai: %v\n", resp.Candidates[0].Content.Parts[0])
	}

}

func GetPromt() string {
	scanner := bufio.NewScanner(os.Stdin)
	var prompt string
	if scanner.Scan() {
		prompt = scanner.Text()
	}
	return prompt
}

func Initialize(ctx context.Context) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, err
	}
	return client, nil
}
