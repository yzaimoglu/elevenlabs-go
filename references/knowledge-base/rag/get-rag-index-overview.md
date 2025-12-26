# Get RAG index overview

GET https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index

Provides total size and other information of RAG indexes used by knowledgebase documents

Reference: https://elevenlabs.io/docs/api-reference/knowledge-base/rag-index-overview

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Get Rag Index Overview.
  version: endpoint_conversationalAi.rag_index_overview
paths:
  /v1/convai/knowledge-base/rag-index:
    get:
      operationId: rag-index-overview
      summary: Get Rag Index Overview.
      description: >-
        Provides total size and other information of RAG indexes used by
        knowledgebase documents
      tags:
        - - subpackage_conversationalAi
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
                $ref: '#/components/schemas/RAGIndexOverviewResponseModel'
        '422':
          description: Validation Error
          content: {}
components:
  schemas:
    EmbeddingModelEnum:
      type: string
      enum:
        - value: e5_mistral_7b_instruct
        - value: multilingual_e5_large_instruct
    RAGIndexOverviewEmbeddingModelResponseModel:
      type: object
      properties:
        model:
          $ref: '#/components/schemas/EmbeddingModelEnum'
        used_bytes:
          type: integer
      required:
        - model
        - used_bytes
    RAGIndexOverviewResponseModel:
      type: object
      properties:
        total_used_bytes:
          type: integer
        total_max_bytes:
          type: integer
        models:
          type: array
          items:
            $ref: '#/components/schemas/RAGIndexOverviewEmbeddingModelResponseModel'
      required:
        - total_used_bytes
        - total_max_bytes
        - models

```

## SDK Code Examples

```typescript
import { ElevenLabsClient } from "@elevenlabs/elevenlabs-js";

async function main() {
    const client = new ElevenLabsClient({
        environment: "https://api.elevenlabs.io",
    });
    await client.conversationalAi.ragIndexOverview();
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.rag_index_overview()

```

```go
package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index"

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

url = URI("https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Get.new(url)
request["xi-api-key"] = 'xi-api-key'

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.get("https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index")
  .header("xi-api-key", "xi-api-key")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('GET', 'https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index', [
  'headers' => [
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index");
var request = new RestRequest(Method.GET);
request.AddHeader("xi-api-key", "xi-api-key");
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = ["xi-api-key": "xi-api-key"]

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/knowledge-base/rag-index")! as URL,
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