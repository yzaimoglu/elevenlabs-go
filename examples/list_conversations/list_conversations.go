package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp/v3"
	"github.com/yzaimoglu/elevenlabs-go/elevenlabs"
)

func main() {
	godotenv.Load()

	// Create a new ElevenLabs API client
	// The Environment can either be EnvironmentProduction, EnvironmentProductionUS, EnvironmentProductionEU or EnvironmentProductionIndia
	client, err := elevenlabs.NewClient(os.Getenv("ELEVENLABS_API_KEY"), elevenlabs.EnvironmentProductionEU)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// List conversations with a minimum call duration of 300 seconds
	resp, err := client.ListConversations(context.Background(), &elevenlabs.ListConversationsReq{
		CallDurationMinSecs: elevenlabs.Ptr(300),
	})
	if err != nil {
		panic(err)
	}

	pp.Println(resp)
}
