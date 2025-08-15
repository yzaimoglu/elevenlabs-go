package examples

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
	pp.Println(agents)

	// Check whether there are more agents to fetch, if so fetch them using the cursor
	if agents.HasMore {
		nextAgents, err := client.ListAgents(context.Background(),
			elevenlabs.NewListAgentsReq(agents.NextCursor, elevenlabs.Ptr(2), nil))
		if err != nil {
			log.Fatalf("Error fetching next agents: %v", err)

		}
		pp.Println(nextAgents)
	}
}
