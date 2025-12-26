<br/>
<p align="center">
  <a href="https://github.com/yzaimoglu/elevenlabs-go">
    <img src=".github/images/gopher_elevenlabs.png" alt="Logo" width="160" height="160">
  </a>

  <h3 align="center">ElevenLabs Go Client</h3>

  <p align="center">
    Unofficial Go Client for the ElevenLabs API
    <br/>
    <br/>
    <a href="https://github.com/yzaimoglu/elevenlabs-go/issues">Report Bug</a>
    .
    <a href="https://github.com/yzaimoglu/elevenlabs-go/issues">Request Feature</a>
  </p>
</p>

![Contributors](https://img.shields.io/github/contributors/yzaimoglu/elevenlabs-go?color=dark-green) ![Issues](https://img.shields.io/github/issues/yzaimoglu/elevenlabs-go) ![License](https://img.shields.io/github/license/yzaimoglu/elevenlabs-go) 

A Go client for interacting with the [ElevenLabs API](https://api.elevenlabs.io/), providing easy-to-use methods for the endpoints for Administration, Conversational AI and advanced features like TTS, STT and other ElevenLabs features.

## Features
This project is currently work in progress and not all of the endpoints are supported.

### Conversational AI - Agents
- [X] Create Agent (POST /v1/convai/agents/create)
- [X] List Agents (GET /v1/convai/agents)
- [X] Get Agent (GET /v1/convai/agents/{agent_id})
- [X] Update Agent (PATCH /v1/convai/agents/{agent_id})
- [X] Delete Agent (DELETE /v1/convai/agents/{agent_id})
- [X] Duplicate Agent (POST /v1/convai/agents/{agent_id}/duplicate)
- [X] Simulate Conversation (POST /v1/convai/agents/{agent_id}/simulate)
- [X] Stream Simulate Conversation (POST /v1/convai/agents/{agent_id}/simulate-conversation/stream)
- [X] Get Agent Dependent Agents (GET /v1/convai/agents/{agent_id}/dependent-agents)
- [X] Get Agent Link (GET /v1/convai/agents/{agent_id}/link)

### Conversational AI - Conversations
- [X] List Conversations (GET /v1/convai/conversations)
- [X] Get Conversation (GET /v1/convai/conversations/{conversation_id})
- [X] Get Conversation Audio (GET /v1/convai/conversations/{conversation_id}/audio)
- [X] Delete Conversation (DELETE /v1/convai/conversations/{conversation_id})
- [X] Get Conversation Token (GET /v1/convai/conversations/{conversation_id}/conversations-token)
- [X] Get Signed URL (GET /v1/convai/conversations/{conversation_id}/get-signed-url)
- [X] Send Conversation Feedback (POST /v1/convai/conversations/{conversation_id}/feedback)

### Conversational AI - Tools
- [X] List Tools (GET /v1/convai/tools)
- [X] Get Tool (GET /v1/convai/tools/{tool_id})
- [X] Create Tool (POST /v1/convai/tools)
- [X] Update Tool (PATCH /v1/convai/tools/{tool_id})
- [X] Delete Tool (DELETE /v1/convai/tools/{tool_id})
- [X] Get Tool Dependent Agents (GET /v1/convai/tools/{tool_id}/dependent-agents)

### Conversational AI - Knowledge Base
- [X] Get Knowledge Base Dependent Agents (GET /v1/convai/knowledge-base/{documentation_id}/dependent-agents)
- [X] Get Knowledge Base Size (GET /v1/convai/agent/{agent_id}/knowledge-base/size)
- [X] Get Knowledge Base Summaries (GET /v1/convai/knowledge-base/summaries)

### Conversational AI - Knowledge Base Documents
- [X] List Knowledge Base Documents (GET /v1/convai/knowledge-base)
- [X] Delete Knowledge Base Document (DELETE /v1/convai/knowledge-base/{documentation_id})
- [X] Get Knowledge Base Document (GET /v1/convai/knowledge-base/{documentation_id})
- [X] Update Knowledge Base Document (PATCH /v1/convai/knowledge-base/{documentation_id})
- [X] Create Knowledge Base Document from URL (POST /v1/convai/knowledge-base/url)
- [X] Create Knowledge Base Document from Text (POST /v1/convai/knowledge-base/text)
- [X] Create Knowledge Base Document from File (POST /v1/convai/knowledge-base/file)
- [X] Get Document Content (GET /v1/convai/knowledge-base/{documentation_id}/content)
- [X] Get Document Chunk (GET /v1/convai/knowledge-base/{documentation_id}/chunk/{chunk_id})

### Conversational AI - Knowledge Base RAG Indexing
- [X] Compute RAG Index (POST /v1/convai/knowledge-base/{documentation_id}/rag-index)
- [X] Get RAG Index (GET /v1/convai/knowledge-base/{documentation_id}/rag-index)
- [X] Compute RAG Index Batch (POST /v1/convai/knowledge-base/rag-index)
- [X] Get RAG Index Overview (GET /v1/convai/knowledge-base/rag-index)
- [X] Delete RAG Index (DELETE /v1/convai/knowledge-base/{documentation_id}/rag-index/{rag_index_id})

## Installation
```bash
go get github.com/yzaimoglu/elevenlabs-go
```

## Usage
```go
package main

import (
    "fmt"
    "log"

    elevenlabs "github.com/yzaimoglu/elevenlabs-go"
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
```

## Documentation
Full API documentation:
- [Elevenlabs API Reference](https://elevenlabs.io/docs/api-reference/introduction)

This client aims to be a 1:1 wrapper for the official ElevenLabs API endpoints.
Refer to the code and GoDoc for detailed usage examples.

## Contributing
1. Fork the repository
2. Create a feature branch: git checkout -b feat/my-feature
3. Commit changes: git commit -m 'Add my feature'
4. Push to branch: git push origin feat/my-feature
5. Create a Pull Request

## License
[BSD-3 License](https://github.com/yzaimoglu/elevenlabs-go/blob/main/LICENSE)

## Acknowledgments
- This project is heavily inspired by the structure of the [unofficial Zendesk Go Client](https://github.com/nukosuke/go-zendesk) by nukosuke.
- Gopher created with [Gopherize](https://gopherize.me/)