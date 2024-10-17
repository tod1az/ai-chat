package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/iterator"
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
	var prompt string
	for prompt != "salir" {
		fmt.Println("Pregunta algo ( escriba salir para terminar )")
		prompt := GetPromt()
		if prompt == "salir" {
			fmt.Println("chau chau")
			os.Exit(1)
		}
		fmt.Printf("Usted: %v\n", prompt)
		iter := model.GenerateContentStream(ctx, genai.Text(prompt))
		fmt.Print("Ai: ")
		for {
			resp, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Printf("Error getting a response: %v\n", err.Error())
			}
			fmt.Print(resp.Candidates[0].Content.Parts[0])
		}
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
