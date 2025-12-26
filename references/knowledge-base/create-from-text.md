# Create knowledge base document from text

POST https://api.elevenlabs.io/v1/convai/knowledge-base/text
Content-Type: application/json

Create a knowledge base document containing the provided text.

Reference: https://elevenlabs.io/docs/api-reference/knowledge-base/create-from-text

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Create Text Document
  version: endpoint_conversationalAi/knowledgeBase/documents.create_from_text
paths:
  /v1/convai/knowledge-base/text:
    post:
      operationId: create-from-text
      summary: Create Text Document
      description: Create a knowledge base document containing the provided text.
      tags:
        - - subpackage_conversationalAi
          - subpackage_conversationalAi/knowledgeBase
          - subpackage_conversationalAi/knowledgeBase/documents
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
                $ref: '#/components/schemas/AddKnowledgeBaseResponseModel'
        '422':
          description: Validation Error
          content: {}
      requestBody:
        content:
          application/json:
            schema:
              $ref: >-
                #/components/schemas/Body_Create_text_document_v1_convai_knowledge_base_text_post
components:
  schemas:
    Body_Create_text_document_v1_convai_knowledge_base_text_post:
      type: object
      properties:
        text:
          type: string
          description: Text content to be added to the knowledge base.
        name:
          type:
            - string
            - 'null'
          description: A custom, human-readable name for the document.
        parent_folder_id:
          type:
            - string
            - 'null'
          description: >-
            If set, the created document or folder will be placed inside the
            given folder.
      required:
        - text
    AddKnowledgeBaseResponseModel:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
      required:
        - id
        - name

```

## SDK Code Examples

```typescript
import { ElevenLabsClient } from "@elevenlabs/elevenlabs-js";

async function main() {
    const client = new ElevenLabsClient({
        environment: "https://api.elevenlabs.io",
    });
    await client.conversationalAi.knowledgeBase.documents.createFromText({
        text: "ElevenLabs provides advanced AI voice synthesis technology that enables developers to create realistic and expressive speech applications.",
    });
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.knowledge_base.documents.create_from_text(
    text="ElevenLabs provides advanced AI voice synthesis technology that enables developers to create realistic and expressive speech applications."
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

	url := "https://api.elevenlabs.io/v1/convai/knowledge-base/text"

	payload := strings.NewReader("{\n  \"text\": \"ElevenLabs provides advanced AI voice synthesis technology that enables developers to create realistic and expressive speech applications.\"\n}")

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

url = URI("https://api.elevenlabs.io/v1/convai/knowledge-base/text")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Post.new(url)
request["xi-api-key"] = 'xi-api-key'
request["Content-Type"] = 'application/json'
request.body = "{\n  \"text\": \"ElevenLabs provides advanced AI voice synthesis technology that enables developers to create realistic and expressive speech applications.\"\n}"

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.post("https://api.elevenlabs.io/v1/convai/knowledge-base/text")
  .header("xi-api-key", "xi-api-key")
  .header("Content-Type", "application/json")
  .body("{\n  \"text\": \"ElevenLabs provides advanced AI voice synthesis technology that enables developers to create realistic and expressive speech applications.\"\n}")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('POST', 'https://api.elevenlabs.io/v1/convai/knowledge-base/text', [
  'body' => '{
  "text": "ElevenLabs provides advanced AI voice synthesis technology that enables developers to create realistic and expressive speech applications."
}',
  'headers' => [
    'Content-Type' => 'application/json',
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/knowledge-base/text");
var request = new RestRequest(Method.POST);
request.AddHeader("xi-api-key", "xi-api-key");
request.AddHeader("Content-Type", "application/json");
request.AddParameter("application/json", "{\n  \"text\": \"ElevenLabs provides advanced AI voice synthesis technology that enables developers to create realistic and expressive speech applications.\"\n}", ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = [
  "xi-api-key": "xi-api-key",
  "Content-Type": "application/json"
]
let parameters = ["text": "ElevenLabs provides advanced AI voice synthesis technology that enables developers to create realistic and expressive speech applications."] as [String : Any]

let postData = JSONSerialization.data(withJSONObject: parameters, options: [])

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/knowledge-base/text")! as URL,
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