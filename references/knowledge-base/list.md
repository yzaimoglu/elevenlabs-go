# List knowledge base documents

GET https://api.elevenlabs.io/v1/convai/knowledge-base

Get a list of available knowledge base documents

Reference: https://elevenlabs.io/docs/api-reference/knowledge-base/list

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Get Knowledge Base List
  version: endpoint_conversationalAi/knowledgeBase.list
paths:
  /v1/convai/knowledge-base:
    get:
      operationId: list
      summary: Get Knowledge Base List
      description: Get a list of available knowledge base documents
      tags:
        - - subpackage_conversationalAi
          - subpackage_conversationalAi/knowledgeBase
      parameters:
        - name: page_size
          in: query
          description: >-
            How many documents to return at maximum. Can not exceed 100,
            defaults to 30.
          required: false
          schema:
            type: integer
            default: 30
        - name: search
          in: query
          description: >-
            If specified, the endpoint returns only such knowledge base
            documents whose names start with this string.
          required: false
          schema:
            type:
              - string
              - 'null'
        - name: show_only_owned_documents
          in: query
          description: >-
            If set to true, the endpoint will return only documents owned by you
            (and not shared from somebody else).
          required: false
          schema:
            type: boolean
            default: false
        - name: types
          in: query
          description: >-
            If present, the endpoint will return only documents of the given
            types.
          required: false
          schema:
            type:
              - array
              - 'null'
            items:
              $ref: '#/components/schemas/KnowledgeBaseDocumentType'
        - name: parent_folder_id
          in: query
          description: >-
            If set, the endpoint will return only documents that are direct
            children of the given folder.
          required: false
          schema:
            type:
              - string
              - 'null'
        - name: ancestor_folder_id
          in: query
          description: >-
            If set, the endpoint will return only documents that are descendants
            of the given folder.
          required: false
          schema:
            type:
              - string
              - 'null'
        - name: folders_first
          in: query
          description: Whether folders should be returned first in the list of documents.
          required: false
          schema:
            type: boolean
            default: false
        - name: sort_direction
          in: query
          description: The direction to sort the results
          required: false
          schema:
            $ref: '#/components/schemas/SortDirection'
        - name: sort_by
          in: query
          description: The field to sort the results by
          required: false
          schema:
            oneOf:
              - $ref: '#/components/schemas/KnowledgeBaseSortBy'
              - type: 'null'
        - name: use_typesense
          in: query
          description: >-
            If set to true, the endpoint will use typesense DB to search for the
            documents).
          required: false
          schema:
            type: boolean
            default: false
        - name: cursor
          in: query
          description: Used for fetching next page. Cursor is returned in the response.
          required: false
          schema:
            type:
              - string
              - 'null'
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
                $ref: '#/components/schemas/GetKnowledgeBaseListResponseModel'
        '422':
          description: Validation Error
          content: {}
components:
  schemas:
    KnowledgeBaseDocumentType:
      type: string
      enum:
        - value: file
        - value: url
        - value: text
        - value: folder
    SortDirection:
      type: string
      enum:
        - value: asc
        - value: desc
    KnowledgeBaseSortBy:
      type: string
      enum:
        - value: name
        - value: created_at
        - value: updated_at
        - value: size
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
    GetKnowledgeBaseSummaryUrlResponseModelDependentAgentsItems:
      oneOf:
        - $ref: '#/components/schemas/DependentAvailableAgentIdentifier'
        - $ref: '#/components/schemas/DependentUnknownAgentIdentifier'
    GetKnowledgeBaseSummaryURLResponseModel:
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
        dependent_agents:
          type: array
          items:
            $ref: >-
              #/components/schemas/GetKnowledgeBaseSummaryUrlResponseModelDependentAgentsItems
          description: >-
            This field is deprecated and will be removed in the future, use the
            separate endpoint to get dependent agents instead.
        type:
          type: string
          enum:
            - type: stringLiteral
              value: url
        url:
          type: string
      required:
        - id
        - name
        - metadata
        - supported_usages
        - access_info
        - dependent_agents
        - type
        - url
    GetKnowledgeBaseSummaryFileResponseModelDependentAgentsItems:
      oneOf:
        - $ref: '#/components/schemas/DependentAvailableAgentIdentifier'
        - $ref: '#/components/schemas/DependentUnknownAgentIdentifier'
    GetKnowledgeBaseSummaryFileResponseModel:
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
        dependent_agents:
          type: array
          items:
            $ref: >-
              #/components/schemas/GetKnowledgeBaseSummaryFileResponseModelDependentAgentsItems
          description: >-
            This field is deprecated and will be removed in the future, use the
            separate endpoint to get dependent agents instead.
        type:
          type: string
          enum:
            - type: stringLiteral
              value: file
      required:
        - id
        - name
        - metadata
        - supported_usages
        - access_info
        - dependent_agents
        - type
    GetKnowledgeBaseSummaryTextResponseModelDependentAgentsItems:
      oneOf:
        - $ref: '#/components/schemas/DependentAvailableAgentIdentifier'
        - $ref: '#/components/schemas/DependentUnknownAgentIdentifier'
    GetKnowledgeBaseSummaryTextResponseModel:
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
        dependent_agents:
          type: array
          items:
            $ref: >-
              #/components/schemas/GetKnowledgeBaseSummaryTextResponseModelDependentAgentsItems
          description: >-
            This field is deprecated and will be removed in the future, use the
            separate endpoint to get dependent agents instead.
        type:
          type: string
          enum:
            - type: stringLiteral
              value: text
      required:
        - id
        - name
        - metadata
        - supported_usages
        - access_info
        - dependent_agents
        - type
    GetKnowledgeBaseSummaryFolderResponseModelDependentAgentsItems:
      oneOf:
        - $ref: '#/components/schemas/DependentAvailableAgentIdentifier'
        - $ref: '#/components/schemas/DependentUnknownAgentIdentifier'
    GetKnowledgeBaseSummaryFolderResponseModel:
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
        dependent_agents:
          type: array
          items:
            $ref: >-
              #/components/schemas/GetKnowledgeBaseSummaryFolderResponseModelDependentAgentsItems
          description: >-
            This field is deprecated and will be removed in the future, use the
            separate endpoint to get dependent agents instead.
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
        - dependent_agents
        - type
        - children_count
    GetKnowledgeBaseListResponseModelDocumentsItems:
      oneOf:
        - $ref: '#/components/schemas/GetKnowledgeBaseSummaryURLResponseModel'
        - $ref: '#/components/schemas/GetKnowledgeBaseSummaryFileResponseModel'
        - $ref: '#/components/schemas/GetKnowledgeBaseSummaryTextResponseModel'
        - $ref: '#/components/schemas/GetKnowledgeBaseSummaryFolderResponseModel'
    GetKnowledgeBaseListResponseModel:
      type: object
      properties:
        documents:
          type: array
          items:
            $ref: >-
              #/components/schemas/GetKnowledgeBaseListResponseModelDocumentsItems
        next_cursor:
          type:
            - string
            - 'null'
        has_more:
          type: boolean
      required:
        - documents
        - has_more

```

## SDK Code Examples

```typescript
import { ElevenLabsClient } from "@elevenlabs/elevenlabs-js";

async function main() {
    const client = new ElevenLabsClient({
        environment: "https://api.elevenlabs.io",
    });
    await client.conversationalAi.knowledgeBase.list({});
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.knowledge_base.list()

```

```go
package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://api.elevenlabs.io/v1/convai/knowledge-base"

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

url = URI("https://api.elevenlabs.io/v1/convai/knowledge-base")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Get.new(url)
request["xi-api-key"] = 'xi-api-key'

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.get("https://api.elevenlabs.io/v1/convai/knowledge-base")
  .header("xi-api-key", "xi-api-key")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('GET', 'https://api.elevenlabs.io/v1/convai/knowledge-base', [
  'headers' => [
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/knowledge-base");
var request = new RestRequest(Method.GET);
request.AddHeader("xi-api-key", "xi-api-key");
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = ["xi-api-key": "xi-api-key"]

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/knowledge-base")! as URL,
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