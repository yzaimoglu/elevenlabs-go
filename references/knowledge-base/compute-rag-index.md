# Compute RAG index

POST https://api.elevenlabs.io/v1/convai/knowledge-base/{documentation_id}/rag-index
Content-Type: application/json

In case the document is not RAG indexed, it triggers rag indexing task, otherwise it just returns the current status.

Reference: https://elevenlabs.io/docs/api-reference/knowledge-base/compute-rag-index

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Compute Rag Index.
  version: endpoint_conversationalAi/knowledgeBase/document.compute_rag_index
paths:
  /v1/convai/knowledge-base/{documentation_id}/rag-index:
    post:
      operationId: compute-rag-index
      summary: Compute Rag Index.
      description: >-
        In case the document is not RAG indexed, it triggers rag indexing task,
        otherwise it just returns the current status.
      tags:
        - - subpackage_conversationalAi
          - subpackage_conversationalAi/knowledgeBase
          - subpackage_conversationalAi/knowledgeBase/document
      parameters:
        - name: documentation_id
          in: path
          description: >-
            The id of a document from the knowledge base. This is returned on
            document addition.
          required: true
          schema:
            type: string
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
                $ref: '#/components/schemas/RAGDocumentIndexResponseModel'
        '422':
          description: Validation Error
          content: {}
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/RAGIndexRequestModel'
components:
  schemas:
    EmbeddingModelEnum:
      type: string
      enum:
        - value: e5_mistral_7b_instruct
        - value: multilingual_e5_large_instruct
    RAGIndexRequestModel:
      type: object
      properties:
        model:
          $ref: '#/components/schemas/EmbeddingModelEnum'
      required:
        - model
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

```

## SDK Code Examples

```typescript
import { ElevenLabsClient } from "@elevenlabs/elevenlabs-js";

async function main() {
    const client = new ElevenLabsClient({
        environment: "https://api.elevenlabs.io",
    });
    await client.conversationalAi.knowledgeBase.document.computeRagIndex("documentation_id", {
        model: "e5_mistral_7b_instruct",
    });
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.knowledge_base.document.compute_rag_index(
    documentation_id="documentation_id",
    model="e5_mistral_7b_instruct"
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

	url := "https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/rag-index"

	payload := strings.NewReader("{\n  \"model\": \"e5_mistral_7b_instruct\"\n}")

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

url = URI("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/rag-index")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Post.new(url)
request["xi-api-key"] = 'xi-api-key'
request["Content-Type"] = 'application/json'
request.body = "{\n  \"model\": \"e5_mistral_7b_instruct\"\n}"

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.post("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/rag-index")
  .header("xi-api-key", "xi-api-key")
  .header("Content-Type", "application/json")
  .body("{\n  \"model\": \"e5_mistral_7b_instruct\"\n}")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('POST', 'https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/rag-index', [
  'body' => '{
  "model": "e5_mistral_7b_instruct"
}',
  'headers' => [
    'Content-Type' => 'application/json',
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/rag-index");
var request = new RestRequest(Method.POST);
request.AddHeader("xi-api-key", "xi-api-key");
request.AddHeader("Content-Type", "application/json");
request.AddParameter("application/json", "{\n  \"model\": \"e5_mistral_7b_instruct\"\n}", ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = [
  "xi-api-key": "xi-api-key",
  "Content-Type": "application/json"
]
let parameters = ["model": "e5_mistral_7b_instruct"] as [String : Any]

let postData = JSONSerialization.data(withJSONObject: parameters, options: [])

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/rag-index")! as URL,
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