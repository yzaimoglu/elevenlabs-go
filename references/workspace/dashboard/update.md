# Update Convai Dashboard Settings

PATCH https://api.elevenlabs.io/v1/convai/settings/dashboard
Content-Type: application/json

Update Convai dashboard settings for the workspace

Reference: https://elevenlabs.io/docs/api-reference/workspace/dashboard/update

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Update Convai Dashboard Settings
  version: endpoint_conversationalAi/dashboard/settings.update
paths:
  /v1/convai/settings/dashboard:
    patch:
      operationId: update
      summary: Update Convai Dashboard Settings
      description: Update Convai dashboard settings for the workspace
      tags:
        - - subpackage_conversationalAi
          - subpackage_conversationalAi/dashboard
          - subpackage_conversationalAi/dashboard/settings
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
                $ref: '#/components/schemas/GetConvAIDashboardSettingsResponseModel'
        '422':
          description: Validation Error
          content: {}
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/PatchConvAIDashboardSettingsRequest'
components:
  schemas:
    DashboardCallSuccessChartModel:
      type: object
      properties:
        name:
          type: string
        type:
          type: string
          enum:
            - type: stringLiteral
              value: call_success
      required:
        - name
    DashboardCriteriaChartModel:
      type: object
      properties:
        name:
          type: string
        type:
          type: string
          enum:
            - type: stringLiteral
              value: criteria
        criteria_id:
          type: string
      required:
        - name
        - criteria_id
    DashboardDataCollectionChartModel:
      type: object
      properties:
        name:
          type: string
        type:
          type: string
          enum:
            - type: stringLiteral
              value: data_collection
        data_collection_id:
          type: string
      required:
        - name
        - data_collection_id
    PatchConvAiDashboardSettingsRequestChartsItems:
      oneOf:
        - $ref: '#/components/schemas/DashboardCallSuccessChartModel'
        - $ref: '#/components/schemas/DashboardCriteriaChartModel'
        - $ref: '#/components/schemas/DashboardDataCollectionChartModel'
    PatchConvAIDashboardSettingsRequest:
      type: object
      properties:
        charts:
          type: array
          items:
            $ref: >-
              #/components/schemas/PatchConvAiDashboardSettingsRequestChartsItems
    GetConvAiDashboardSettingsResponseModelChartsItems:
      oneOf:
        - $ref: '#/components/schemas/DashboardCallSuccessChartModel'
        - $ref: '#/components/schemas/DashboardCriteriaChartModel'
        - $ref: '#/components/schemas/DashboardDataCollectionChartModel'
    GetConvAIDashboardSettingsResponseModel:
      type: object
      properties:
        charts:
          type: array
          items:
            $ref: >-
              #/components/schemas/GetConvAiDashboardSettingsResponseModelChartsItems

```

## SDK Code Examples

```typescript
import { ElevenLabsClient } from "@elevenlabs/elevenlabs-js";

async function main() {
    const client = new ElevenLabsClient({
        environment: "https://api.elevenlabs.io",
    });
    await client.conversationalAi.dashboard.settings.update({
        charts: [
            {
                type: "call_success",
                name: "Call Success Rate Overview",
            },
            {
                type: "criteria",
                name: "Customer Satisfaction Criteria",
                criteriaId: "crit-8f4a2b7d",
            },
            {
                type: "data_collection",
                name: "User Feedback Data Collection",
                dataCollectionId: "dc-3e9f1c2a",
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

client.conversational_ai.dashboard.settings.update(
    charts=[
        {
            "type": "call_success",
            "name": "Call Success Rate Overview"
        },
        {
            "type": "criteria",
            "name": "Customer Satisfaction Criteria",
            "criteria_id": "crit-8f4a2b7d"
        },
        {
            "type": "data_collection",
            "name": "User Feedback Data Collection",
            "data_collection_id": "dc-3e9f1c2a"
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

	url := "https://api.elevenlabs.io/v1/convai/settings/dashboard"

	payload := strings.NewReader("{\n  \"charts\": [\n    {\n      \"name\": \"Call Success Rate Overview\",\n      \"type\": \"call_success\"\n    },\n    {\n      \"name\": \"Customer Satisfaction Criteria\",\n      \"type\": \"criteria\",\n      \"criteria_id\": \"crit-8f4a2b7d\"\n    },\n    {\n      \"name\": \"User Feedback Data Collection\",\n      \"type\": \"data_collection\",\n      \"data_collection_id\": \"dc-3e9f1c2a\"\n    }\n  ]\n}")

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

url = URI("https://api.elevenlabs.io/v1/convai/settings/dashboard")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Patch.new(url)
request["xi-api-key"] = 'xi-api-key'
request["Content-Type"] = 'application/json'
request.body = "{\n  \"charts\": [\n    {\n      \"name\": \"Call Success Rate Overview\",\n      \"type\": \"call_success\"\n    },\n    {\n      \"name\": \"Customer Satisfaction Criteria\",\n      \"type\": \"criteria\",\n      \"criteria_id\": \"crit-8f4a2b7d\"\n    },\n    {\n      \"name\": \"User Feedback Data Collection\",\n      \"type\": \"data_collection\",\n      \"data_collection_id\": \"dc-3e9f1c2a\"\n    }\n  ]\n}"

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.patch("https://api.elevenlabs.io/v1/convai/settings/dashboard")
  .header("xi-api-key", "xi-api-key")
  .header("Content-Type", "application/json")
  .body("{\n  \"charts\": [\n    {\n      \"name\": \"Call Success Rate Overview\",\n      \"type\": \"call_success\"\n    },\n    {\n      \"name\": \"Customer Satisfaction Criteria\",\n      \"type\": \"criteria\",\n      \"criteria_id\": \"crit-8f4a2b7d\"\n    },\n    {\n      \"name\": \"User Feedback Data Collection\",\n      \"type\": \"data_collection\",\n      \"data_collection_id\": \"dc-3e9f1c2a\"\n    }\n  ]\n}")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('PATCH', 'https://api.elevenlabs.io/v1/convai/settings/dashboard', [
  'body' => '{
  "charts": [
    {
      "name": "Call Success Rate Overview",
      "type": "call_success"
    },
    {
      "name": "Customer Satisfaction Criteria",
      "type": "criteria",
      "criteria_id": "crit-8f4a2b7d"
    },
    {
      "name": "User Feedback Data Collection",
      "type": "data_collection",
      "data_collection_id": "dc-3e9f1c2a"
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
var client = new RestClient("https://api.elevenlabs.io/v1/convai/settings/dashboard");
var request = new RestRequest(Method.PATCH);
request.AddHeader("xi-api-key", "xi-api-key");
request.AddHeader("Content-Type", "application/json");
request.AddParameter("application/json", "{\n  \"charts\": [\n    {\n      \"name\": \"Call Success Rate Overview\",\n      \"type\": \"call_success\"\n    },\n    {\n      \"name\": \"Customer Satisfaction Criteria\",\n      \"type\": \"criteria\",\n      \"criteria_id\": \"crit-8f4a2b7d\"\n    },\n    {\n      \"name\": \"User Feedback Data Collection\",\n      \"type\": \"data_collection\",\n      \"data_collection_id\": \"dc-3e9f1c2a\"\n    }\n  ]\n}", ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = [
  "xi-api-key": "xi-api-key",
  "Content-Type": "application/json"
]
let parameters = ["charts": [
    [
      "name": "Call Success Rate Overview",
      "type": "call_success"
    ],
    [
      "name": "Customer Satisfaction Criteria",
      "type": "criteria",
      "criteria_id": "crit-8f4a2b7d"
    ],
    [
      "name": "User Feedback Data Collection",
      "type": "data_collection",
      "data_collection_id": "dc-3e9f1c2a"
    ]
  ]] as [String : Any]

let postData = JSONSerialization.data(withJSONObject: parameters, options: [])

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/settings/dashboard")! as URL,
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