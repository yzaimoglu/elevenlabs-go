# Compute RAG index in batch

POST https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index
Content-Type: application/json

Retrieves and/or creates RAG indexes for multiple knowledge base documents in a single request.

Reference: https://elevenlabs.io/docs/api-reference/knowledge-base/compute-rag-index-batch

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Compute Rag Indexes In Batch
  version: endpoint_conversationalAi/knowledgeBase.get_or_create_rag_indexes
paths:
  /v1/convai/knowledge-base/rag-index:
    post:
      operationId: get-or-create-rag-indexes
      summary: Compute Rag Indexes In Batch
      description: >-
        Retrieves and/or creates RAG indexes for multiple knowledge base
        documents in a single request.
      tags:
        - - subpackage_conversationalAi
          - subpackage_conversationalAi/knowledgeBase
      parameters:
        - name: xi-api-key
          in: header
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Successful Response
          content:
            application/json:
              schema:
                type: object
                additionalProperties:
                  $ref: >-
                    #/components/schemas/V1ConvaiKnowledgeBaseRagIndexPostResponsesContentApplicationJsonSchema
        '422':
          description: Validation Error
          content: {}
      requestBody:
        content:
          application/json:
            schema:
              $ref: >-
                #/components/schemas/Body_Compute_RAG_indexes_in_batch_v1_convai_knowledge_base_rag_index_post
components:
  schemas:
    EmbeddingModelEnum:
      type: string
      enum:
        - value: e5_mistral_7b_instruct
        - value: multilingual_e5_large_instruct
    GetOrCreateRAGIndexRequestModel:
      type: object
      properties:
        document_id:
          type: string
          description: ID of the knowledgebase document for which to retrieve the index
        create_if_missing:
          type: boolean
          description: Whether to create the RAG index if it does not exist
        model:
          $ref: '#/components/schemas/EmbeddingModelEnum'
          description: Embedding model to use for the RAG index
      required:
        - document_id
        - create_if_missing
        - model
    Body_Compute_RAG_indexes_in_batch_v1_convai_knowledge_base_rag_index_post:
      type: object
      properties:
        items:
          type: array
          items:
            $ref: '#/components/schemas/GetOrCreateRAGIndexRequestModel'
          description: List of requested RAG indexes.
      required:
        - items
    RAGIndexStatus:
      type: string
      enum:
        - value: created
        - value: processing
        - value: failed
        - value: succeeded
        - value: rag_limit_exceeded
        - value: document_too_small
        - value: cannot_index_folder
    RAGDocumentIndexUsage:
      type: object
      properties:
        used_bytes:
          type: integer
      required:
        - used_bytes
    RAGDocumentIndexResponseModel:
      type: object
      properties:
        id:
          type: string
        model:
          $ref: '#/components/schemas/EmbeddingModelEnum'
        status:
          $ref: '#/components/schemas/RAGIndexStatus'
        progress_percentage:
          type: number
          format: double
        document_model_index_usage:
          $ref: '#/components/schemas/RAGDocumentIndexUsage'
      required:
        - id
        - model
        - status
        - progress_percentage
        - document_model_index_usage
    RAGIndexBatchSuccessfulResponseModel:
      type: object
      properties:
        status:
          type: string
          enum:
            - type: stringLiteral
              value: success
        data:
          $ref: '#/components/schemas/RAGDocumentIndexResponseModel'
      required:
        - status
        - data
    BatchFailureResponseModel:
      type: object
      properties:
        status:
          type: string
          enum:
            - type: stringLiteral
              value: failure
        error_code:
          type: integer
        error_status:
          type: string
        error_message:
          type: string
      required:
        - status
        - error_code
        - error_status
        - error_message
    V1ConvaiKnowledgeBaseRagIndexPostResponsesContentApplicationJsonSchema:
      oneOf:
        - $ref: '#/components/schemas/RAGIndexBatchSuccessfulResponseModel'
        - $ref: '#/components/schemas/BatchFailureResponseModel'

```

## SDK Code Examples

```typescript
import { ElevenLabsClient } from "@elevenlabs/elevenlabs-js";

async function main() {
    const client = new ElevenLabsClient({
        environment: "https://api.elevenlabs.io",
    });
    await client.conversationalAi.knowledgeBase.getOrCreateRagIndexes({
        items: [
            {
                documentId: "kb-doc-987654321",
                createIfMissing: true,
                model: "e5_mistral_7b_instruct",
            },
            {
                documentId: "kb-doc-123456789",
                createIfMissing: false,
                model: "multilingual_e5_large_instruct",
            },
        ],
    });
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.knowledge_base.get_or_create_rag_indexes(
    items=[
        {
            "document_id": "kb-doc-987654321",
            "create_if_missing": True,
            "model": "e5_mistral_7b_instruct"
        },
        {
            "document_id": "kb-doc-123456789",
            "create_if_missing": False,
            "model": "multilingual_e5_large_instruct"
        }
    ]
)

```

```go
package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
)

func main() {

	url := "https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index"

	payload := strings.NewReader("{\n  \"items\": [\n    {\n      \"document_id\": \"kb-doc-987654321\",\n      \"create_if_missing\": true,\n      \"model\": \"e5_mistral_7b_instruct\"\n    },\n    {\n      \"document_id\": \"kb-doc-123456789\",\n      \"create_if_missing\": false,\n      \"model\": \"multilingual_e5_large_instruct\"\n    }\n  ]\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("xi-api-key", "xi-api-key")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
```

```ruby
require 'uri'
require 'net/http'

url = URI("https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Post.new(url)
request["xi-api-key"] = 'xi-api-key'
request["Content-Type"] = 'application/json'
request.body = "{\n  \"items\": [\n    {\n      \"document_id\": \"kb-doc-987654321\",\n      \"create_if_missing\": true,\n      \"model\": \"e5_mistral_7b_instruct\"\n    },\n    {\n      \"document_id\": \"kb-doc-123456789\",\n      \"create_if_missing\": false,\n      \"model\": \"multilingual_e5_large_instruct\"\n    }\n  ]\n}"

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.post("https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index")
  .header("xi-api-key", "xi-api-key")
  .header("Content-Type", "application/json")
  .body("{\n  \"items\": [\n    {\n      \"document_id\": \"kb-doc-987654321\",\n      \"create_if_missing\": true,\n      \"model\": \"e5_mistral_7b_instruct\"\n    },\n    {\n      \"document_id\": \"kb-doc-123456789\",\n      \"create_if_missing\": false,\n      \"model\": \"multilingual_e5_large_instruct\"\n    }\n  ]\n}")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('POST', 'https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index', [
  'body' => '{
  "items": [
    {
      "document_id": "kb-doc-987654321",
      "create_if_missing": true,
      "model": "e5_mistral_7b_instruct"
    },
    {
      "document_id": "kb-doc-123456789",
      "create_if_missing": false,
      "model": "multilingual_e5_large_instruct"
    }
  ]
}',
  'headers' => [
    'Content-Type' => 'application/json',
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index");
var request = new RestRequest(Method.POST);
request.AddHeader("xi-api-key", "xi-api-key");
request.AddHeader("Content-Type", "application/json");
request.AddParameter("application/json", "{\n  \"items\": [\n    {\n      \"document_id\": \"kb-doc-987654321\",\n      \"create_if_missing\": true,\n      \"model\": \"e5_mistral_7b_instruct\"\n    },\n    {\n      \"document_id\": \"kb-doc-123456789\",\n      \"create_if_missing\": false,\n      \"model\": \"multilingual_e5_large_instruct\"\n    }\n  ]\n}", ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = [
  "xi-api-key": "xi-api-key",
  "Content-Type": "application/json"
]
let parameters = ["items": [
    [
      "document_id": "kb-doc-987654321",
      "create_if_missing": true,
      "model": "e5_mistral_7b_instruct"
    ],
    [
      "document_id": "kb-doc-123456789",
      "create_if_missing": false,
      "model": "multilingual_e5_large_instruct"
    ]
  ]] as [String : Any]

let postData = JSONSerialization.data(withJSONObject: parameters, options: [])

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index")! as URL,
                                        cachePolicy: .useProtocolCachePolicy,
                                    timeoutInterval: 10.0)
request.httpMethod = "POST"
request.allHTTPHeaderFields = headers
request.httpBody = postData as Data

let session = URLSession.shared
let dataTask = session.dataTask(with: request as URLRequest, completionHandler: { (data, response, error) -> Void in
  if (error != nil) {
    print(error as Any)
  } else {
    let httpResponse = response as? HTTPURLResponse
    print(httpResponse)
  }
})

dataTask.resume()
```