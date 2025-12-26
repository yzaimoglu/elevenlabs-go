# Get secrets

GET https://api.elevenlabs.io/v1/convai/secrets

Get all workspace secrets for the user

Reference: https://elevenlabs.io/docs/api-reference/workspace/secrets/list

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Get Convai Workspace Secrets
  version: endpoint_conversationalAi/secrets.list
paths:
  /v1/convai/secrets:
    get:
      operationId: list
      summary: Get Convai Workspace Secrets
      description: Get all workspace secrets for the user
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
                $ref: '#/components/schemas/GetWorkspaceSecretsResponseModel'
        '422':
          description: Validation Error
          content: {}
components:
  schemas:
    DependentAvailableToolIdentifierAccessLevel:
      type: string
      enum:
        - value: admin
        - value: editor
        - value: commenter
        - value: viewer
    DependentAvailableToolIdentifier:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        type:
          type: string
          enum:
            - type: stringLiteral
              value: available
        created_at_unix_secs:
          type: integer
        access_level:
          $ref: '#/components/schemas/DependentAvailableToolIdentifierAccessLevel'
      required:
        - id
        - name
        - created_at_unix_secs
        - access_level
    DependentUnknownToolIdentifier:
      type: object
      properties:
        type:
          type: string
          enum:
            - type: stringLiteral
              value: unknown
    ConvAiStoredSecretDependenciesToolsItems:
      oneOf:
        - $ref: '#/components/schemas/DependentAvailableToolIdentifier'
        - $ref: '#/components/schemas/DependentUnknownToolIdentifier'
    DependentAvailableAgentIdentifierAccessLevel:
      type: string
      enum:
        - value: admin
        - value: editor
        - value: commenter
        - value: viewer
    DependentAvailableAgentIdentifier:
      type: object
      properties:
        referenced_resource_ids:
          type: array
          items:
            type: string
          description: >-
            If the agent is a transitive dependent, contains IDs of the
            resources that the agent depends on directly.
        id:
          type: string
        name:
          type: string
        type:
          type: string
          enum:
            - type: stringLiteral
              value: available
        created_at_unix_secs:
          type: integer
        access_level:
          $ref: '#/components/schemas/DependentAvailableAgentIdentifierAccessLevel'
      required:
        - id
        - name
        - created_at_unix_secs
        - access_level
    DependentUnknownAgentIdentifier:
      type: object
      properties:
        referenced_resource_ids:
          type: array
          items:
            type: string
          description: >-
            If the agent is a transitive dependent, contains IDs of the
            resources that the agent depends on directly.
        type:
          type: string
          enum:
            - type: stringLiteral
              value: unknown
    ConvAiStoredSecretDependenciesAgentsItems:
      oneOf:
        - $ref: '#/components/schemas/DependentAvailableAgentIdentifier'
        - $ref: '#/components/schemas/DependentUnknownAgentIdentifier'
    SecretDependencyType:
      type: string
      enum:
        - value: conversation_initiation_webhook
    TelephonyProvider:
      type: string
      enum:
        - value: twilio
        - value: sip_trunk
    DependentPhoneNumberIdentifier:
      type: object
      properties:
        phone_number_id:
          type: string
        phone_number:
          type: string
        label:
          type: string
        provider:
          $ref: '#/components/schemas/TelephonyProvider'
      required:
        - phone_number_id
        - phone_number
        - label
        - provider
    ConvAIStoredSecretDependencies:
      type: object
      properties:
        tools:
          type: array
          items:
            $ref: '#/components/schemas/ConvAiStoredSecretDependenciesToolsItems'
        agents:
          type: array
          items:
            $ref: '#/components/schemas/ConvAiStoredSecretDependenciesAgentsItems'
        others:
          type: array
          items:
            $ref: '#/components/schemas/SecretDependencyType'
        phone_numbers:
          type: array
          items:
            $ref: '#/components/schemas/DependentPhoneNumberIdentifier'
      required:
        - tools
        - agents
        - others
    ConvAIWorkspaceStoredSecretConfig:
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
        used_by:
          $ref: '#/components/schemas/ConvAIStoredSecretDependencies'
      required:
        - type
        - secret_id
        - name
        - used_by
    GetWorkspaceSecretsResponseModel:
      type: object
      properties:
        secrets:
          type: array
          items:
            $ref: '#/components/schemas/ConvAIWorkspaceStoredSecretConfig'
      required:
        - secrets

```

## SDK Code Examples

```typescript
import { ElevenLabsClient } from "@elevenlabs/elevenlabs-js";

async function main() {
    const client = new ElevenLabsClient({
        environment: "https://api.elevenlabs.io",
    });
    await client.conversationalAi.secrets.list();
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.secrets.list()

```

```go
package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://api.elevenlabs.io/v1/convai/secrets"

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

url = URI("https://api.elevenlabs.io/v1/convai/secrets")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Get.new(url)
request["xi-api-key"] = 'xi-api-key'

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.get("https://api.elevenlabs.io/v1/convai/secrets")
  .header("xi-api-key", "xi-api-key")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('GET', 'https://api.elevenlabs.io/v1/convai/secrets', [
  'headers' => [
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/secrets");
var request = new RestRequest(Method.GET);
request.AddHeader("xi-api-key", "xi-api-key");
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = ["xi-api-key": "xi-api-key"]

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/secrets")! as URL,
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