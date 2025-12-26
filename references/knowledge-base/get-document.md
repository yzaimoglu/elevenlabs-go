# Get knowledge base document

GET https://api.elevenlabs.io/v1/convai/knowledge-base/{documentation_id}

Get details about a specific documentation making up the agent's knowledge base

Reference: https://elevenlabs.io/docs/api-reference/knowledge-base/get-document

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Get Documentation From Knowledge Base
  version: endpoint_conversationalAi/knowledgeBase/documents.get
paths:
  /v1/convai/knowledge-base/{documentation_id}:
    get:
      operationId: get
      summary: Get Documentation From Knowledge Base
      description: >-
        Get details about a specific documentation making up the agent's
        knowledge base
      tags:
        - - subpackage_conversationalAi
          - subpackage_conversationalAi/knowledgeBase
          - subpackage_conversationalAi/knowledgeBase/documents
      parameters:
        - name: documentation_id
          in: path
          description: >-
            The id of a document from the knowledge base. This is returned on
            document addition.
          required: true
          schema:
            type: string
        - name: agent_id
          in: query
          required: false
          schema:
            type: string
            default: ''
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
                $ref: >-
                  #/components/schemas/conversational_ai_knowledge_base_documents_get_Response_200
        '422':
          description: Validation Error
          content: {}
components:
  schemas:
    KnowledgeBaseDocumentMetadataResponseModel:
      type: object
      properties:
        created_at_unix_secs:
          type: integer
        last_updated_at_unix_secs:
          type: integer
        size_bytes:
          type: integer
      required:
        - created_at_unix_secs
        - last_updated_at_unix_secs
        - size_bytes
    DocumentUsageModeEnum:
      type: string
      enum:
        - value: prompt
        - value: auto
    ResourceAccessInfoRole:
      type: string
      enum:
        - value: admin
        - value: editor
        - value: commenter
        - value: viewer
    ResourceAccessInfo:
      type: object
      properties:
        is_creator:
          type: boolean
          description: Whether the user making the request is the creator of the agent
        creator_name:
          type: string
          description: Name of the agent's creator
        creator_email:
          type: string
          description: Email of the agent's creator
        role:
          $ref: '#/components/schemas/ResourceAccessInfoRole'
          description: The role of the user making the request
      required:
        - is_creator
        - creator_name
        - creator_email
        - role
    KnowledgeBaseFolderPathSegmentResponseModel:
      type: object
      properties:
        id:
          type: string
        name:
          type:
            - string
            - 'null'
      required:
        - id
        - name
    GetKnowledgeBaseURLResponseModel:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        metadata:
          $ref: '#/components/schemas/KnowledgeBaseDocumentMetadataResponseModel'
        supported_usages:
          type: array
          items:
            $ref: '#/components/schemas/DocumentUsageModeEnum'
        access_info:
          $ref: '#/components/schemas/ResourceAccessInfo'
        folder_parent_id:
          type:
            - string
            - 'null'
          description: >-
            The ID of the parent folder, or null if the document is at the root
            level.
        folder_path:
          type: array
          items:
            $ref: '#/components/schemas/KnowledgeBaseFolderPathSegmentResponseModel'
          description: >-
            The folder path segments leading to this entity, from root to parent
            folder.
        type:
          type: string
          enum:
            - type: stringLiteral
              value: url
        url:
          type: string
        extracted_inner_html:
          type: string
      required:
        - id
        - name
        - metadata
        - supported_usages
        - access_info
        - type
        - url
        - extracted_inner_html
    GetKnowledgeBaseFileResponseModel:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        metadata:
          $ref: '#/components/schemas/KnowledgeBaseDocumentMetadataResponseModel'
        supported_usages:
          type: array
          items:
            $ref: '#/components/schemas/DocumentUsageModeEnum'
        access_info:
          $ref: '#/components/schemas/ResourceAccessInfo'
        folder_parent_id:
          type:
            - string
            - 'null'
          description: >-
            The ID of the parent folder, or null if the document is at the root
            level.
        folder_path:
          type: array
          items:
            $ref: '#/components/schemas/KnowledgeBaseFolderPathSegmentResponseModel'
          description: >-
            The folder path segments leading to this entity, from root to parent
            folder.
        type:
          type: string
          enum:
            - type: stringLiteral
              value: file
        extracted_inner_html:
          type: string
      required:
        - id
        - name
        - metadata
        - supported_usages
        - access_info
        - type
        - extracted_inner_html
    GetKnowledgeBaseTextResponseModel:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        metadata:
          $ref: '#/components/schemas/KnowledgeBaseDocumentMetadataResponseModel'
        supported_usages:
          type: array
          items:
            $ref: '#/components/schemas/DocumentUsageModeEnum'
        access_info:
          $ref: '#/components/schemas/ResourceAccessInfo'
        folder_parent_id:
          type:
            - string
            - 'null'
          description: >-
            The ID of the parent folder, or null if the document is at the root
            level.
        folder_path:
          type: array
          items:
            $ref: '#/components/schemas/KnowledgeBaseFolderPathSegmentResponseModel'
          description: >-
            The folder path segments leading to this entity, from root to parent
            folder.
        type:
          type: string
          enum:
            - type: stringLiteral
              value: text
        extracted_inner_html:
          type: string
      required:
        - id
        - name
        - metadata
        - supported_usages
        - access_info
        - type
        - extracted_inner_html
    GetKnowledgeBaseFolderResponseModel:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        metadata:
          $ref: '#/components/schemas/KnowledgeBaseDocumentMetadataResponseModel'
        supported_usages:
          type: array
          items:
            $ref: '#/components/schemas/DocumentUsageModeEnum'
        access_info:
          $ref: '#/components/schemas/ResourceAccessInfo'
        folder_parent_id:
          type:
            - string
            - 'null'
          description: >-
            The ID of the parent folder, or null if the document is at the root
            level.
        folder_path:
          type: array
          items:
            $ref: '#/components/schemas/KnowledgeBaseFolderPathSegmentResponseModel'
          description: >-
            The folder path segments leading to this entity, from root to parent
            folder.
        type:
          type: string
          enum:
            - type: stringLiteral
              value: folder
        children_count:
          type: integer
      required:
        - id
        - name
        - metadata
        - supported_usages
        - access_info
        - type
        - children_count
    conversational_ai_knowledge_base_documents_get_Response_200:
      oneOf:
        - $ref: '#/components/schemas/GetKnowledgeBaseURLResponseModel'
        - $ref: '#/components/schemas/GetKnowledgeBaseFileResponseModel'
        - $ref: '#/components/schemas/GetKnowledgeBaseTextResponseModel'
        - $ref: '#/components/schemas/GetKnowledgeBaseFolderResponseModel'

```

## SDK Code Examples

```typescript
import { ElevenLabsClient } from "@elevenlabs/elevenlabs-js";

async function main() {
    const client = new ElevenLabsClient({
        environment: "https://api.elevenlabs.io",
    });
    await client.conversationalAi.knowledgeBase.documents.get("documentation_id", {});
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.knowledge_base.documents.get(
    documentation_id="documentation_id"
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

	url := "https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id"

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

url = URI("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Get.new(url)
request["xi-api-key"] = 'xi-api-key'

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.get("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id")
  .header("xi-api-key", "xi-api-key")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('GET', 'https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id', [
  'headers' => [
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id");
var request = new RestRequest(Method.GET);
request.AddHeader("xi-api-key", "xi-api-key");
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = ["xi-api-key": "xi-api-key"]

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id")! as URL,
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