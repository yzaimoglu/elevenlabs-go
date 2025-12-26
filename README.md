<br/>
<p align="center">
  <a href="https://github.com/yzaimoglu/elevenlabs-go">
    <img src=".github/images/gopher_elevenlabs.png" alt="Logo" width="160" height="160">
  </a>

  <h3 align="center">ElevenLabs Go Client</h3>

  <p align="center">
    Unofficial Go Client for the ElevenLabs API
    <br/>
    (implements cli version <strong>v2.15.0</strong>)
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
This project has all of the endpoints for the Agent Platform from v2.15.0 of the Elevenlabs API documentation implemented.

### Conversational AI - Agents
- [X] Create Agent (POST /v1/convai/agents/create)
- [X] List Agents (GET /v1/convai/agents)
- [X] Get Agent (GET /v1/convai/agents/{agent_id})
- [X] Update Agent (PATCH /v1/convai/agents/{agent_id})
- [X] Delete Agent (DELETE /v1/convai/agents/{agent_id})
- [X] Duplicate Agent (POST /v1/convai/agents/{agent_id}/duplicate)
- [X] Get Agent Link (GET /v1/convai/agents/{agent_id}/link)
- [X] Simulate Conversation (POST /v1/convai/agents/{agent_id}/simulate-conversation)
- [X] Stream Simulate Conversation (POST /v1/convai/agents/{agent_id}/simulate-conversation/stream)
- [X] Run Tests on Agent (POST /v1/convai/agents/{agent_id}/run-tests)
- [X] Calculate Agent LLM Usage (POST /v1/convai/agent/{agent_id}/llm-usage/calculate)

### Conversational AI - Conversations
- [X] List Conversations (GET /v1/convai/conversations)
- [X] Get Conversation (GET /v1/convai/conversations/{conversation_id})
- [X] Get Conversation Audio (GET /v1/convai/conversations/{conversation_id}/audio)
- [X] Delete Conversation (DELETE /v1/convai/conversations/{conversation_id})
- [X] Get Conversation Token (GET /v1/convai/conversation/token)
- [X] Get Signed URL (GET /v1/convai/conversation/get-signed-url)
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

### Conversational AI - Phone Numbers
- [X] List Phone Numbers (GET /v1/convai/phone-numbers)
- [X] Get Phone Number (GET /v1/convai/phone-numbers/{phone_number_id})
- [X] Create Phone Number (POST /v1/convai/phone-numbers)
- [X] Update Phone Number (PATCH /v1/convai/phone-numbers/{phone_number_id})
- [X] Delete Phone Number (DELETE /v1/convai/phone-numbers/{phone_number_id})

### Conversational AI - Batch Calling
- [X] Submit Batch Call (POST /v1/convai/batch-calling/submit)
- [X] Get Batch Call (GET /v1/convai/batch-calling/{batch_id})
- [X] List Batch Calls (GET /v1/convai/batch-calling/workspace)
- [X] Retry Batch Call (POST /v1/convai/batch-calling/{batch_id}/retry)
- [X] Cancel Batch Call (POST /v1/convai/batch-calling/{batch_id}/cancel)

### Conversational AI - Twilio
- [X] Register Twilio Call (POST /v1/convai/twilio/register-call)
- [X] Twilio Outbound Call (POST /v1/convai/twilio/outbound-call)

### Conversational AI - SIP Trunk
- [X] SIP Trunk Outbound Call (POST /v1/convai/sip-trunk/outbound-call)

### Conversational AI - WhatsApp
- [X] WhatsApp Outbound Call (POST /v1/convai/whatsapp/outbound-call)

### Conversational AI - WhatsApp Accounts
- [X] List WhatsApp Accounts (GET /v1/convai/whatsapp-accounts)
- [X] Get WhatsApp Account (GET /v1/convai/whatsapp-accounts/{phone_number_id})
- [X] Import WhatsApp Account (POST /v1/convai/whatsapp-accounts)
- [X] Update WhatsApp Account (PATCH /v1/convai/whatsapp-accounts/{phone_number_id})
- [X] Delete WhatsApp Account (DELETE /v1/convai/whatsapp-accounts/{phone_number_id})

### Conversational AI - Analytics
- [X] Get Live Count (GET /v1/convai/analytics/live-count)

### Conversational AI - Widget
- [X] Get Agent Widget (GET /v1/convai/agents/{agent_id}/widget)
- [X] Post Agent Avatar (POST /v1/convai/agents/{agent_id}/avatar)

### Conversational AI - Workspace Settings
- [X] Get Settings (GET /v1/convai/settings)
- [X] Update Settings (PATCH /v1/convai/settings)

### Conversational AI - Tests
- [X] List Tests (GET /v1/convai/agent-testing)
- [X] Get Test (GET /v1/convai/agent-testing/{test_id})
- [X] Create Test (POST /v1/convai/agent-testing/create)
- [X] Update Test (PUT /v1/convai/agent-testing/{test_id})
- [X] Delete Test (DELETE /v1/convai/agent-testing/{test_id})
- [X] Get Test Summaries (POST /v1/convai/agent-testing/summaries)

### Conversational AI - MCP Servers
- [X] List MCP Servers (GET /v1/convai/mcp-servers)
- [X] Get MCP Server (GET /v1/convai/mcp-servers/{mcp_server_id})
- [X] Create MCP Server (POST /v1/convai/mcp-servers)
- [X] Update MCP Server (PATCH /v1/convai/mcp-servers/{mcp_server_id})
- [X] Delete MCP Server (DELETE /v1/convai/mcp-servers/{mcp_server_id})
- [X] List MCP Server Tools (GET /v1/convai/mcp-servers/{mcp_server_id}/tools)

### Conversational AI - MCP Tool Approvals
- [X] Create Tool Approval (POST /v1/convai/mcp-servers/{mcp_server_id}/tool-approvals)
- [X] Update Approval Policy (PATCH /v1/convai/mcp-servers/{mcp_server_id}/approval-policy)
- [X] Delete Tool Approval (DELETE /v1/convai/mcp-servers/{mcp_server_id}/tool-approvals/{tool_name})

### Conversational AI - MCP Tool Configuration
- [X] Create Tool Config (POST /v1/convai/mcp-servers/{mcp_server_id}/tool-configs)
- [X] Get Tool Config (GET /v1/convai/mcp-servers/{mcp_server_id}/tool-configs/{tool_name})
- [X] Update Tool Config (PATCH /v1/convai/mcp-servers/{mcp_server_id}/tool-configs/{tool_name})
- [X] Delete Tool Config (DELETE /v1/convai/mcp-servers/{mcp_server_id}/tool-configs/{tool_name})

### Conversational AI - Workspace Secrets
- [X] List Secrets (GET /v1/convai/secrets)
- [X] Create Secret (POST /v1/convai/secrets)
- [X] Update Secret (PATCH /v1/convai/secrets/{secret_id})
- [X] Delete Secret (DELETE /v1/convai/secrets/{secret_id})

### Conversational AI - Workspace Dashboard
- [X] Get Dashboard (GET /v1/convai/settings/dashboard)
- [X] Update Dashboard (PATCH /v1/convai/settings/dashboard)

### Conversational AI - LLM Usage
- [X] Calculate LLM Usage (POST /v1/convai/llm-usage/calculate)

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

This client aims to be a 1:1 wrapper for the official ElevenLabs API Agent Platform endpoints.
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
- [Claude Code](https://github.com/anthropics/claude-code) for easing the process of porting over the client methods from the official Elevenlabs API documentation