# Technology Radar
This is a training project aiming to create a Technology radar were each team could manage technologies and their maturity

## Version: 1.0.0

### /technologies

#### POST
##### Summary

Create technology

##### Description

Create a new technology to specified team

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 400 | the request sent is not valid |
| 409 | The specified resource was already created |
| 500 | our application did not behave well |

### /team/{team}/technologies

#### GET
##### Summary

Get all technologies by team

##### Description

<strong>Given </strong> a team <br> <strong>I want to</strong> get all technologies <br> <strong>So i could</strong> create a technology radar from the given team

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| team | path | Team name | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 404 | the resource required was not found |
| 500 | our application did not behave well |

### /team/{team}/quadrants/{quadrant}/technologies

#### GET
##### Summary

Get all technologies by team and quadrant

##### Description

<strong>Given </strong> a team and quadrant <br> <strong>I want to</strong> get all technologies <br> <strong>So i could</strong> create zoom in into specific quadrant

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| team | path | Team name | Yes | string |
| quadrant | path | Friendly title generated during the technology creation | Yes | integer |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 404 | the resource required was not found |
| 500 | our application did not behave well |

### /team/{team}/technologies/{technology_friendly_title}

#### GET
##### Summary

Get a technology by team and title

##### Description

<strong>Given </strong> a team and title<br> <strong>I want to</strong> get the related technologies <br> <strong>So i could</strong> display the details and all transitions

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| team | path | Team name | Yes | string |
| technology_friendly_title | path | Friendly title generated during the technology creation | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 404 | the resource required was not found |
| 500 | our application did not behave well |

#### PATCH
##### Summary

Update technology

##### Description

<strong>Given </strong> a team and title<br> <strong>I want to</strong> update some fields <br> <strong>So i could</strong> correct or change some characteristics

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| team | path | Team name | Yes | string |
| technology_friendly_title | path | Friendly title generated during the technology creation | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 404 | the resource required was not found |
| 500 | our application did not behave well |

#### POST
##### Summary

Move technology

##### Description

<strong>Given </strong> a team and title<br> <strong>I want to</strong> move the technology between ring <br> <strong>So i could</strong> evolve my technology radar

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| team | path | Team name | Yes | string |
| technology_friendly_title | path | Friendly title generated during the technology creation | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 204 | NO_CONTENT |
| 404 | the resource required was not found |
| 500 | our application did not behave well |

### /infos/team

#### GET
##### Summary

Get all quadrants

##### Description

<strong>Given </strong> i call this endpoint<br> <strong>I want to</strong> all quadrants and name <br> <strong>So i could</strong> create the quadrants chart on frontend

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | our application did not behave well |

### /infos/quadrants

#### GET
##### Summary

Get all quadrants

##### Description

<strong>Given </strong> i call this endpoint<br> <strong>I want to</strong> all quadrants and name <br> <strong>So i could</strong> create the quadrants chart on frontend

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |
| 500 | our application did not behave well |

#### POST
##### Summary

Create quadrant

##### Description

<strong>Given </strong> i call this endpoint<br> <strong>I want to</strong> create or update the quadrant <br> <strong>So i could</strong> manage my quadrants

##### Responses

| Code | Description |
| ---- | ----------- |
| 204 | NO_CONTENT |
| 500 | our application did not behave well |

### Models

#### TechnologyResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| team | string | Team name<br>_Example:_ `"github"` | No |
| title | string | Technology's title<br>_Example:_ `"Four-key metrics"` | No |
| friendlyTitle | string | Technology's friendly title, this field is internally created replacing spaces with slashes and converting to lower case<br>_Example:_ `"four-key-metrics"` | No |
| description | string | Technology's description<br>_Example:_ `"To measure software delivery performance, more and more organizations are adopting four key metrics as interpreted by the DORA research program: lead time, deployment frequency, mean time to restore (MTTR) and improvement failure enhancement. Its statistical analysis shows a clear delivery and search among high performers that provides a great indicator of performance for the same time, or even an entire organization."` | No |
| quadrant | integer | Quadrant number id where tech belongs<br>_Example:_ `1` | No |
| score | integer | Technology score<br>_Example:_ `3` | No |
| active | boolean | Identify if the technology is active | No |
| moved | boolean | Identify if the last movement changed the score | No |
| history | [ object ] | Map all transitions and their values | No |
| created_at | string | Date when the resource was created<br>_Example:_ `{}` | No |
| updated_at | string | Date of last update<br>_Example:_ `{}` | No |

#### QuadrantResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | number | Quadrant's id<br>_Example:_ `3` | No |
| title | string | Quadrant title<br>_Example:_ `"Framework and Languages"` | No |

#### TechnologyRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| team | string | Team name<br>_Example:_ `"github"` | Yes |
| title | string | Technology's title<br>_Example:_ `"Four-key metrics"` | Yes |
| description | string | Technology's description<br>_Example:_ `"To measure software delivery performance, more and more organizations are adopting four key metrics as interpreted by the DORA research program: lead time, deployment frequency, mean time to restore (MTTR) and improvement failure enhancement. Its statistical analysis shows a clear delivery and search among high performers that provides a great indicator of performance for the same time, or even an entire organization."` | Yes |
| quadrant | integer | Quadrant number id where tech belongs<br>_Example:_ `1` | Yes |
| score | integer | Technology score<br>_Example:_ `3` | Yes |

#### MoveTechnologyRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| comments | string | Quadrant title<br>_Example:_ `"Framework and Languages"` | No |
| score | integer | Technology score<br>_Example:_ `3` | No |

#### UpdateTechnologyRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| description | string | Technology's description<br>_Example:_ `"To measure software delivery performance, more and more organizations are adopting four key metrics as interpreted by the DORA research program: lead time, deployment frequency, mean time to restore (MTTR) and improvement failure enhancement. Its statistical analysis shows a clear delivery and search among high performers that provides a great indicator of performance for the same time, or even an entire organization."` | No |
| quadrant | integer | Quadrant number id where tech belongs<br>_Example:_ `1` | No |
| active | boolean | Identify if the technology is active | No |

#### CreateQuadrantRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | number | Quadrant's id<br>_Example:_ `3` | Yes |
| title | string | Quadrant title<br>_Example:_ `"Framework and Languages"` | Yes |
