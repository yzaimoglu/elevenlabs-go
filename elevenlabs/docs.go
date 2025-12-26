// Package elevenlabs provides a comprehensive Go client for the ElevenLabs API.
//
// This package offers a type-safe, idiomatic Go interface for all ElevenLabs API endpoints,
// including Conversational AI, text-to-speech, speech-to-text, and voice management.
//
// # Getting Started
//
// To use this package, you'll need an ElevenLabs API key. You can obtain one from:
// https://elevenlabs.io/api
//
// Basic usage:
//
//	// Create a new client
//	client, err := elevenlabs.NewClient("your-api-key", elevenlabs.EnvironmentProduction)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// List all your conversational AI agents
//	agents, err := client.ListAgents(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, agent := range agents.Tools {
//	    fmt.Printf("Agent: %s (ID: %s)\n", agent.Name, agent.Id)
//	}
//
// # Environments
//
// The package supports multiple ElevenLabs environments:
//   - EnvironmentProduction: Main production API
//   - EnvironmentProductionUS: US-specific production endpoint
//   - EnvironmentProductionEU: EU-specific production endpoint
//   - EnvironmentProductionIndia: India-specific production endpoint
//
// Choose the environment closest to your location for optimal performance.
//
// # API Structure
//
// The client is organized into logical API groups, each implementing a specific interface:
//
//   - ConvaiAgentsAPI: Manage conversational AI agents (create, update, list, delete)
//   - ConvaiConversationsAPI: Handle conversation history and details
//   - ConvaiToolsAPI: Manage tools that agents can use
//   - ConvaiKnowledgeBaseAPI: Manage knowledge base documents
//   - ConvaiKnowledgeBaseDocumentsAPI: CRUD operations for knowledge base documents
//   - ConvaiKnowledgeBaseRAGAPI: RAG (Retrieval-Augmented Generation) indexing
//
// # Request and Response Patterns
//
// All API methods follow consistent patterns:
//
// Request structs use constructors for proper initialization:
//
//	// Create agent request with constructor
//	req := elevenlabs.NewCreateAgentReq(convConfig)
//	req.Name = elevenlabs.String("My Assistant")
//	req.Tags = []string{"support", "production"}
//
//	agent, err := client.CreateAgent(context.Background(), req)
//
// Optional fields use pointers to distinguish between zero values and unset fields:
//
//	// String helper creates string pointers
//	req.Name = elevenlabs.String("Updated Name")
//
//	// Int helper creates int pointers
//	req.ResponseTimeoutSecs = elevenlabs.Int(30)
//
// # Context Support
//
// All API methods accept context.Context for request cancellation and timeout control:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	agents, err := client.ListAgents(ctx)
//
// # Error Handling
//
// The package returns errors that provide detailed information about API failures:
//
//	_, err := client.GetAgent(context.Background(), elevenlabs.NewGetAgentReq("agent-id"))
//	if err != nil {
//	    // Use errors.As to check if it's an ElevenlabsError
//	    var apiErr elevenlabs.ElevenlabsError
//	    if errors.As(err, &apiErr) {
//	        // Check the status code to handle different error types
//	        switch apiErr.Status() {
//	        case 404:
//	            log.Println("Agent not found")
//	        case 429:
//	            log.Println("Rate limited - too many requests")
//	        case 422:
//	            log.Println("Validation error")
//	        default:
//	            log.Printf("API error %d: %v", apiErr.Status(), err)
//	        }
//	    } else {
//	        log.Printf("Error: %v", err)
//	    }
//	}
//
// Common HTTP status codes:
//   - 401: Invalid API key
//   - 404: Resource not found
//   - 422: Validation error
//   - 429: Rate limiting
//   - 5xx: Server errors
//
// # Conversational AI Example
//
// Creating a conversational AI agent with custom configuration:
//
//	convConfig := elevenlabs.ConversationConfig{
//	    Agent: elevenlabs.ConversationConfigAgent{
//	        Prompt: elevenlabs.AgentPrompt{
//	            Prompt: "You are a helpful assistant...",
//	            LLM:    elevenlabs.LLMGPT4_1,
//	        },
//	        Language: elevenlabs.String("en"),
//	    },
//	}
//
//	createReq := elevenlabs.NewCreateAgentReq(convConfig)
//	createReq.Name = elevenlabs.String("Support Bot")
//
//	agent, err := client.CreateAgent(context.Background(), createReq)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	fmt.Printf("Created agent: %s (ID: %s)\n", agent.Name, agent.Id)
//
// # Working with Conversations
//
// Retrieve conversation details with full typing:
//
//	conv, err := client.GetConversation(context.Background(),
//	    elevenlabs.NewGetConversationReq("conversation-id"))
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Access typed transcript messages
//	for _, msg := range conv.Transcript {
//	    if msg.Message != nil {
//	        fmt.Printf("%s: %s\n", msg.Role, *msg.Message)
//	    }
//
//	    // Access typed tool calls if present
//	    for _, toolCall := range msg.ToolCalls {
//	        fmt.Printf("  Tool: %s\n", toolCall.ToolName)
//	    }
//	}
//
// # Knowledge Base Management
//
// Adding documents to your knowledge base:
//
//	// From URL
//	urlReq := elevenlabs.NewCreateKnowledgeBaseDocumentFromURLReq("https://example.com/docs")
//	doc, err := client.CreateKnowledgeBaseDocumentFromURL(context.Background(), urlReq)
//
//	// From text
//	textReq := elevenlabs.NewCreateKnowledgeBaseDocumentFromTextReq("Your content here")
//	textReq.Name = elevenlabs.String("Important Notes")
//	doc, err = client.CreateKnowledgeBaseDocumentFromText(context.Background(), textReq)
//
//	// From file
//	file, _ := os.Open("document.pdf")
//	defer file.Close()
//
//	fileReq := &elevenlabs.CreateKnowledgeBaseDocumentFromFileReq{
//	    File: file,
//	    Name: elevenlabs.String("Document"),
//	}
//	doc, err = client.CreateKnowledgeBaseDocumentFromFile(context.Background(), fileReq)
//
// # Tools Management
//
// Create a webhook tool for your agents:
//
//	toolConfig := elevenlabs.ToolConfig{
//	    Type: elevenlabs.ConvaiToolTypeWebhook,
//	    Name: "Get Weather",
//	    Description: "Get current weather for a location",
//	    APISchema: elevenlabs.WebhookToolApiSchemaConfig{
//	        URL: "https://api.weather.com/current",
//	        Method: elevenlabs.WebhookMethodGET,
//	        QueryParamsSchema: map[string]elevenlabs.LiteralJsonSchemaProperty{
//	            "city": {
//	                Type: elevenlabs.LiteralJsonSchemaPropertyTypeString,
//	                Description: elevenlabs.String("City name"),
//	            },
//	        },
//	    },
//	}
//
//	tool, err := client.CreateTool(context.Background(),
//	    elevenlabs.NewCreateToolReq(toolConfig))
//
// # Pagination and List Operations
//
// Many list operations support pagination:
//
//	req := elevenlabs.NewListConversationsReq()
//	req.PageSize = 50 // Default page size
//
//	conversations, err := client.ListConversations(context.Background(), req)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, conv := range conversations.Conversations {
//	    fmt.Printf("Conversation: %s\n", conv.ConversationId)
//	}
//
//	// Check for more results
//	if conversations.HasMore {
//	    fmt.Printf("More results available, use cursor: %s\n", *conversations.NextCursor)
//	}
//
// # Query Parameters
//
// Many endpoints support filtering and sorting:
//
//	req := elevenlabs.NewListAgentsReq()
//	req.Archived = elevenlabs.Bool(false)           // Filter by archived status
//	req.ShowOnlyOwnedAgents = elevenlabs.Bool(true) // Show only owned agents
//	req.SortBy = elevenlabs.ListAgentsSortByName    // Sort by name
//	req.SortDirection = elevenlabs.SortDirectionAsc // Ascending order
//
//	agents, err := client.ListAgents(context.Background(), req)
//
// # Configuration Options
//
// The client supports additional configuration:
//
//	client, err := elevenlabs.NewClient("api-key", elevenlabs.EnvironmentProduction,
//	    elevenlabs.WithHTTPClient(myCustomClient),
//	    elevenlabs.WithDebug(true))
//
// Debug mode logs all requests and responses for troubleshooting.
//
// # Thread Safety
//
// The client is thread-safe and can be shared across multiple goroutines:
//
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//	    wg.Add(1)
//	    go func() {
//	        defer wg.Done()
//	        agents, _ := client.ListAgents(context.Background())
//	        fmt.Printf("Found %d agents\n", len(agents.Tools))
//	    }()
//	}
//	wg.Wait()
//
// # Best Practices
//
// 1. Always use context with timeouts for production code
// 2. Reuse the client instance - it's thread-safe
// 3. Check for nil pointers on optional fields
// 4. Use the provided constructor functions for requests
// 5. Handle errors appropriately and check for specific error types
// 6. Use pointer helper (elevenlabs.Ptr) for optional fields
//
// # Rate Limiting
//
// The ElevenLabs API has rate limits. When exceeded, the client returns an error with status 429.
// Use errors.As to check if an error is of type elevenlabs.ElevenlabsError and inspect the status code:
//
//	for attempt := 1; attempt <= 3; attempt++ {
//	    _, err := client.CreateAgent(context.Background(), req)
//	    if err != nil {
//	        // Use errors.As to check if the error is of type elevenlabs.ElevenlabsError
//	        var apiErr elevenlabs.ElevenlabsError
//	        if errors.As(err, &apiErr) {
//	            // Check the status code - if it's 429, we're rate limited
//	            if apiErr.Status() == 429 {
//	                // Implement exponential backoff or other rate limiting strategy
//	                backoff := time.Duration(attempt*attempt) * time.Second
//	                time.Sleep(backoff)
//	                continue
//	            }
//	        }
//	        return err // Other errors should be handled differently
//	    }
//	    break // Success
//	}
//
// # Additional Resources
//
// For more information, see:
//   - ElevenLabs API Reference: https://elevenlabs.io/docs/api-reference/introduction
//   - GoDoc for this package: https://pkg.go.dev/github.com/yzaimoglu/elevenlabs-go
//   - GitHub Repository: https://github.com/yzaimoglu/elevenlabs-go
//
// # Contributing
//
// This is an unofficial SDK. For issues, feature requests, or contributions,
// please visit the GitHub repository.
package elevenlabs
