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
	client, err := elevenlabs.NewClient(os.Getenv("ELEVENLABS_API_KEY"), elevenlabs.EnvironmentProduction)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create a new agent by specifying some parameters
	resp, err := client.CreateAgent(context.Background(), elevenlabs.NewCreateAgentReq(
		&elevenlabs.ConversationConfig{
			Agent: &elevenlabs.ConversationConfigAgent{
				FirstMessage: elevenlabs.Ptr("Hello, this is an example agent created with the Go API Client."),
			},
			Conversation: &elevenlabs.ConversationConfigConversation{
				MaxDurationSeconds: elevenlabs.Ptr(60),
			},
		}, &elevenlabs.PlatformSettings{
			Privacy: &elevenlabs.PlatformSettingsPrivacy{
				ZeroRetentionMode: elevenlabs.Ptr(true),
				RecordVoice:       elevenlabs.Ptr(false),
			},
		},
		elevenlabs.Ptr("Go API Client Example Agent"), nil))
	if err != nil {
		panic(err)
	}

	pp.Println(resp)
}
