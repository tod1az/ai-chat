package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
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
	resp, err := model.GenerateContent(ctx, genai.Text("Cuentame un chiste"))
	if err != nil {
		fmt.Printf("Error getting a response: %v\n", err.Error())
	}
	fmt.Println(resp.Candidates[0].Content.Parts[0])
}

func Initialize(ctx context.Context) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, err
	}
	return client, nil
}
