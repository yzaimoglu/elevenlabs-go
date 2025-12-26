# Update knowledge base document

PATCH https://api.elevenlabs.io/v1/convai/knowledge-base/{documentation_id}
Content-Type: application/json

Update the name of a document

Reference: https://elevenlabs.io/docs/api-reference/knowledge-base/update

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Update Document
  version: endpoint_conversationalAi/knowledgeBase/documents.update
paths:
  /v1/convai/knowledge-base/{documentation_id}:
    patch:
      operationId: update
      summary: Update Document
      description: Update the name of a document
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
                  #/components/schemas/conversational_ai_knowledge_base_documents_update_Response_200
        '422':
          description: Validation Error
          content: {}
      requestBody:
        content:
          application/json:
            schema:
              $ref: >-
                #/components/schemas/Body_Update_document_v1_convai_knowledge_base__documentation_id__patch
components:
  schemas:
    Body_Update_document_v1_convai_knowledge_base__documentation_id__patch:
      type: object
      properties:
        name:
          type: string
          description: A custom, human-readable name for the document.
      required:
        - name
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
    conversational_ai_knowledge_base_documents_update_Response_200:
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
    await client.conversationalAi.knowledgeBase.documents.update("documentation_id", {
        name: "Updated Product Documentation",
    });
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.knowledge_base.documents.update(
    documentation_id="documentation_id",
    name="Updated Product Documentation"
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

	url := "https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id"

	payload := strings.NewReader("{\n  \"name\": \"Updated Product Documentation\"\n}")

	req, _ := http.NewRequest("PATCH", url, payload)

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

url = URI("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Patch.new(url)
request["xi-api-key"] = 'xi-api-key'
request["Content-Type"] = 'application/json'
request.body = "{\n  \"name\": \"Updated Product Documentation\"\n}"

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.patch("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id")
  .header("xi-api-key", "xi-api-key")
  .header("Content-Type", "application/json")
  .body("{\n  \"name\": \"Updated Product Documentation\"\n}")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('PATCH', 'https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id', [
  'body' => '{
  "name": "Updated Product Documentation"
}',
  'headers' => [
    'Content-Type' => 'application/json',
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id");
var request = new RestRequest(Method.PATCH);
request.AddHeader("xi-api-key", "xi-api-key");
request.AddHeader("Content-Type", "application/json");
request.AddParameter("application/json", "{\n  \"name\": \"Updated Product Documentation\"\n}", ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = [
  "xi-api-key": "xi-api-key",
  "Content-Type": "application/json"
]
let parameters = ["name": "Updated Product Documentation"] as [String : Any]

let postData = JSONSerialization.data(withJSONObject: parameters, options: [])

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/knowledge-base/documentation_id")! as URL,
                                        cachePolicy: .useProtocolCachePolicy,
                                    timeoutInterval: 10.0)
request.httpMethod = "PATCH"
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