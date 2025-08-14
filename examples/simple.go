package examples

import (
	"fmt"
	"log"

	"github.com/yzaimoglu/elevenlabs-go/elevenlabs"
)

func main() {
	// Create a new ElevenLabs API client
	// The Environment can either be EnvironmentProduction, EnvironmentProductionUS, EnvironmentProductionEU or EnvironmentProductionIndia
	client, err := elevenlabs.NewClient("<YOUR_API_KEY>", elevenlabs.EnvironmentProductionEU)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Use a method that calls the API (Listing all Conversational AI agents here)
	agents, err := client.ConversationalAI.Agents.ListAgents()
	if err != nil {
		log.Fatalf("Error fetching agents: %v", err)
	}

	for _, a := range agents {
		fmt.Printf("Agent Name: %s (ID: %s)\n", a.Name, a.AgentId)
	}
}
