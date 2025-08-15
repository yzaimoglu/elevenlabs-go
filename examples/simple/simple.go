package examples

import (
	"context"
	"fmt"
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

	// Use a method that calls the API (Listing all Conversational AI agents here)
	agents, err := client.ListAgents(context.Background(),
		elevenlabs.NewListAgentsReq(nil, elevenlabs.Ptr(2), nil))
	if err != nil {
		log.Fatalf("Error fetching agents: %v", err)
	}

	for _, a := range agents.Agents {
		fmt.Printf("Agent Name: %s (ID: %s)\n", a.Name, a.AgentId)
	}
}
