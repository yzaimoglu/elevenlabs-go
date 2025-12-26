# Create secret

POST https://api.elevenlabs.io/v1/convai/secrets
Content-Type: application/json

Create a new secret for the workspace

Reference: https://elevenlabs.io/docs/api-reference/workspace/secrets/create

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Create Convai Workspace Secret
  version: endpoint_conversationalAi/secrets.create
paths:
  /v1/convai/secrets:
    post:
      operationId: create
      summary: Create Convai Workspace Secret
      description: Create a new secret for the workspace
      tags:
        - - subpackage_conversationalAi
          - subpackage_conversationalAi/secrets
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
                $ref: '#/components/schemas/PostWorkspaceSecretResponseModel'
        '422':
          description: Validation Error
          content: {}
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/PostWorkspaceSecretRequest'
components:
  schemas:
    PostWorkspaceSecretRequest:
      type: object
      properties:
        type:
          type: string
          enum:
            - type: stringLiteral
              value: new
        name:
          type: string
        value:
          type: string
      required:
        - type
        - name
        - value
    PostWorkspaceSecretResponseModel:
      type: object
      properties:
        type:
          type: string
          enum:
            - type: stringLiteral
              value: stored
        secret_id:
          type: string
        name:
          type: string
      required:
        - type
        - secret_id
        - name

```

## SDK Code Examples

```typescript
import { ElevenLabsClient } from "@elevenlabs/elevenlabs-js";

async function main() {
    const client = new ElevenLabsClient({
        environment: "https://api.elevenlabs.io",
    });
    await client.conversationalAi.secrets.create({
        type: "new",
        name: "DATABASE_PASSWORD",
        value: "s3cureP@ssw0rd2024!",
    });
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.secrets.create(
    type="new",
    name="DATABASE_PASSWORD",
    value="s3cureP@ssw0rd2024!"
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

	url := "https://api.elevenlabs.io/v1/convai/secrets"

	payload := strings.NewReader("{\n  \"type\": \"new\",\n  \"name\": \"DATABASE_PASSWORD\",\n  \"value\": \"s3cureP@ssw0rd2024!\"\n}")

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

url = URI("https://api.elevenlabs.io/v1/convai/secrets")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Post.new(url)
request["xi-api-key"] = 'xi-api-key'
request["Content-Type"] = 'application/json'
request.body = "{\n  \"type\": \"new\",\n  \"name\": \"DATABASE_PASSWORD\",\n  \"value\": \"s3cureP@ssw0rd2024!\"\n}"

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.post("https://api.elevenlabs.io/v1/convai/secrets")
  .header("xi-api-key", "xi-api-key")
  .header("Content-Type", "application/json")
  .body("{\n  \"type\": \"new\",\n  \"name\": \"DATABASE_PASSWORD\",\n  \"value\": \"s3cureP@ssw0rd2024!\"\n}")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('POST', 'https://api.elevenlabs.io/v1/convai/secrets', [
  'body' => '{
  "type": "new",
  "name": "DATABASE_PASSWORD",
  "value": "s3cureP@ssw0rd2024!"
}',
  'headers' => [
    'Content-Type' => 'application/json',
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/secrets");
var request = new RestRequest(Method.POST);
request.AddHeader("xi-api-key", "xi-api-key");
request.AddHeader("Content-Type", "application/json");
request.AddParameter("application/json", "{\n  \"type\": \"new\",\n  \"name\": \"DATABASE_PASSWORD\",\n  \"value\": \"s3cureP@ssw0rd2024!\"\n}", ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = [
  "xi-api-key": "xi-api-key",
  "Content-Type": "application/json"
]
let parameters = [
  "type": "new",
  "name": "DATABASE_PASSWORD",
  "value": "s3cureP@ssw0rd2024!"
] as [String : Any]

let postData = JSONSerialization.data(withJSONObject: parameters, options: [])

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/secrets")! as URL,
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