package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yzaimoglu/elevenlabs-go/elevenlabs"
)

func main() {
	godotenv.Load()

	// Create a new ElevenLabs API client
	// The Environment can either be EnvironmentProduction, EnvironmentProductionUS, EnvironmentProductionEU or EnvironmentProductionIndia
	client, err := elevenlabs.NewClient(os.Getenv("ELEVENLABS_API_KEY"), elevenlabs.EnvironmentProduction)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ids := []string{"agent_0401k2wv75h5envs8g9yf9g6b9vg", "agent_7201k2wv32eaejzsmgtgkhh8sezz", "agent_8301k2wtznwmentrdxx8m5hcbrh1"}

	for _, id := range ids {
		if err := client.DeleteAgent(context.Background(), elevenlabs.NewDeleteAgentReq(id)); err != nil {
			log.Printf("Failed to delete agent %s: %v", id, err)
		}
		log.Printf("Successfully deleted agent %s", id)
	}
}
