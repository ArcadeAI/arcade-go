# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/shared#AuthorizationContext">AuthorizationContext</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/shared#AuthorizationResponse">AuthorizationResponse</a>

# Auth

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#AuthRequestParam">AuthRequestParam</a>

Methods:

- <code title="post /v1/auth/authorize">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#AuthService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#AuthAuthorizeParams">AuthAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/shared#AuthorizationResponse">AuthorizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/auth/status">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#AuthService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#AuthStatusParams">AuthStatusParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/shared#AuthorizationResponse">AuthorizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#HealthSchema">HealthSchema</a>

Methods:

- <code title="get /v1/health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#HealthSchema">HealthSchema</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ChatMessageParam">ChatMessageParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ChatRequestParam">ChatRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ChatMessage">ChatMessage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ChatResponse">ChatResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#Choice">Choice</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#Usage">Usage</a>

## Completions

Methods:

- <code title="post /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ChatResponse">ChatResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tools

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#AuthorizeToolRequestParam">AuthorizeToolRequestParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ExecuteToolRequestParam">ExecuteToolRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ExecuteToolResponse">ExecuteToolResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolDefinition">ToolDefinition</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolExecution">ToolExecution</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolExecutionAttempt">ToolExecutionAttempt</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ValueSchema">ValueSchema</a>

Methods:

- <code title="get /v1/tools">client.Tools.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolListParams">ToolListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolDefinition">ToolDefinition</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/tools/authorize">client.Tools.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolAuthorizeParams">ToolAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/shared#AuthorizationResponse">AuthorizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/tools/execute">client.Tools.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolService.Execute">Execute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolExecuteParams">ToolExecuteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ExecuteToolResponse">ExecuteToolResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tools/{name}">client.Tools.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolGetParams">ToolGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolDefinition">ToolDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Scheduled

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolScheduledGetResponse">ToolScheduledGetResponse</a>

Methods:

- <code title="get /v1/scheduled_tools">client.Tools.Scheduled.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolScheduledService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolScheduledListParams">ToolScheduledListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolExecution">ToolExecution</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/scheduled_tools/{id}">client.Tools.Scheduled.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolScheduledService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolScheduledGetResponse">ToolScheduledGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Formatted

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolFormattedListResponse">ToolFormattedListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolFormattedGetResponse">ToolFormattedGetResponse</a>

Methods:

- <code title="get /v1/formatted_tools">client.Tools.Formatted.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolFormattedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolFormattedListParams">ToolFormattedListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolFormattedListResponse">ToolFormattedListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/formatted_tools/{name}">client.Tools.Formatted.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolFormattedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolFormattedGetParams">ToolFormattedGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolFormattedGetResponse">ToolFormattedGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workers

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#CreateWorkerRequestParam">CreateWorkerRequestParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#UpdateWorkerRequestParam">UpdateWorkerRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerHealthResponse">WorkerHealthResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerResponse">WorkerResponse</a>

Methods:

- <code title="post /v1/workers">client.Workers.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerNewParams">WorkerNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerResponse">WorkerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/workers/{id}">client.Workers.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerUpdateParams">WorkerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerResponse">WorkerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workers">client.Workers.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerListParams">WorkerListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerResponse">WorkerResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/workers/{id}">client.Workers.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/workers/{id}">client.Workers.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerResponse">WorkerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workers/{id}/health">client.Workers.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerHealthResponse">WorkerHealthResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workers/{id}/tools">client.Workers.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerService.Tools">Tools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#WorkerToolsParams">WorkerToolsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go">arcadeengine</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/arcade-engine-go#ToolDefinition">ToolDefinition</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
