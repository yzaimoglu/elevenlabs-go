# Get dashboard settings

GET https://api.elevenlabs.io/v1/convai/settings/dashboard

Retrieve Convai dashboard settings for the workspace

Reference: https://elevenlabs.io/docs/api-reference/workspace/dashboard/get

## OpenAPI Specification

```yaml
openapi: 3.1.1
info:
  title: Get Convai Dashboard Settings
  version: endpoint_conversationalAi/dashboard/settings.get
paths:
  /v1/convai/settings/dashboard:
    get:
      operationId: get
      summary: Get Convai Dashboard Settings
      description: Retrieve Convai dashboard settings for the workspace
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
    await client.conversationalAi.dashboard.settings.get();
}
main();

```

```python
from elevenlabs import ElevenLabs

client = ElevenLabs(
    base_url="https://api.elevenlabs.io"
)

client.conversational_ai.dashboard.settings.get()

```

```go
package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://api.elevenlabs.io/v1/convai/settings/dashboard"

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

url = URI("https://api.elevenlabs.io/v1/convai/settings/dashboard")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Get.new(url)
request["xi-api-key"] = 'xi-api-key'

response = http.request(request)
puts response.read_body
```

```java
HttpResponse<String> response = Unirest.get("https://api.elevenlabs.io/v1/convai/settings/dashboard")
  .header("xi-api-key", "xi-api-key")
  .asString();
```

```php
<?php

$client = new \GuzzleHttp\Client();

$response = $client->request('GET', 'https://api.elevenlabs.io/v1/convai/settings/dashboard', [
  'headers' => [
    'xi-api-key' => 'xi-api-key',
  ],
]);

echo $response->getBody();
```

```csharp
var client = new RestClient("https://api.elevenlabs.io/v1/convai/settings/dashboard");
var request = new RestRequest(Method.GET);
request.AddHeader("xi-api-key", "xi-api-key");
IRestResponse response = client.Execute(request);
```

```swift
import Foundation

let headers = ["xi-api-key": "xi-api-key"]

let request = NSMutableURLRequest(url: NSURL(string: "https://api.elevenlabs.io/v1/convai/settings/dashboard")! as URL,
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