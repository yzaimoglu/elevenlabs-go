# Get document chunk

GET https://api.elevenlabs.io/v1/convai/knowledge-base/{documentation_id}/chunk/{chunk_id}

Get details about a specific documentation part used by RAG.

Reference: https://elevenlabs.io/docs/api-reference/knowledge-base/get-chunk

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Get Documentation Chunk From Knowledge Base
  version: endpoint_conversationalAi/knowledgeBase/documents/chunk.get
paths:
  /v1/convai/knowledge-base/{documentation_id}/chunk/{chunk_id}:
    get:
      operationId: get
      summary: Get Documentation Chunk From Knowledge Base
      description: Get details about a specific documentation part used by RAG.
      tags:
        - - subpackage_conversationalAi
          - subpackage_conversationalAi/knowledgeBase
          - subpackage_conversationalAi/knowledgeBase/documents
          - subpackage_conversationalAi/knowledgeBase/documents/chunk
      parameters:
        - name: documentation_id
          in: path
          description: >-
            The id of a document from the knowledge base. This is returned on
            document addition.
          required: true
          schema:
            type: string
        - name: chunk_id
          in: path
          description: The id of a document RAG chunk from the knowledge base.
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
                $ref: '#/components/schemas/KnowledgeBaseDocumentChunkResponseModel'
        '422':
          description: Validation Error
          content: {}
components:
  schemas:
    KnowledgeBaseDocumentChunkResponseModel:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        content:
          type: string
      required:
        - id
        - name
        - content

```

## SDK Code Examples

```typescript
import { ElevenLabsClient } from "@elevenlabs/elevenlabs-js";

async function main() {
    const client = new ElevenLabsClient({
        environment: "https://api.elevenlabs.io",
    });
    await client.conversationalAi.knowledgeBase.documents.chunk.get("documentation_id", "chunk_id");
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.knowledge_base.documents.chunk.get(
    documentation_id="documentation_id",
    chunk_id="chunk_id"
)

```

```go
package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/chunk/chunk_id"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("xi-api-key", "xi-api-key")

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

url = URI("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/chunk/chunk_id")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Get.new(url)
request["xi-api-key"] = 'xi-api-key'

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.get("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/chunk/chunk_id")
  .header("xi-api-key", "xi-api-key")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('GET', 'https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/chunk/chunk_id', [
  'headers' => [
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/chunk/chunk_id");
var request = new RestRequest(Method.GET);
request.AddHeader("xi-api-key", "xi-api-key");
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = ["xi-api-key": "xi-api-key"]

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id/chunk/chunk_id")! as URL,
                                        cachePolicy: .useProtocolCachePolicy,
                                    timeoutInterval: 10.0)
request.httpMethod = "GET"
request.allHTTPHeaderFields = headers

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