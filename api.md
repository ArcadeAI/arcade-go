# Shared Response Types

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/shared#AuthorizationContext">AuthorizationContext</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/shared#AuthorizationResponse">AuthorizationResponse</a>

# Admin

## UserConnections

Response Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#UserConnectionResponse">UserConnectionResponse</a>

Methods:

- <code title="get /v1/admin/user_connections">client.Admin.UserConnections.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminUserConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminUserConnectionListParams">AdminUserConnectionListParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#UserConnectionResponse">UserConnectionResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/admin/user_connections/{id}">client.Admin.UserConnections.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminUserConnectionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## AuthProviders

Params Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthProviderCreateRequestParam">AuthProviderCreateRequestParam</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthProviderUpdateRequestParam">AuthProviderUpdateRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthProviderResponse">AuthProviderResponse</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminAuthProviderListResponse">AdminAuthProviderListResponse</a>

Methods:

- <code title="post /v1/admin/auth_providers">client.Admin.AuthProviders.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminAuthProviderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminAuthProviderNewParams">AdminAuthProviderNewParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthProviderResponse">AuthProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/admin/auth_providers">client.Admin.AuthProviders.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminAuthProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminAuthProviderListResponse">AdminAuthProviderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/admin/auth_providers/{id}">client.Admin.AuthProviders.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminAuthProviderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthProviderResponse">AuthProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/admin/auth_providers/{id}">client.Admin.AuthProviders.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminAuthProviderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthProviderResponse">AuthProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/admin/auth_providers/{id}">client.Admin.AuthProviders.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminAuthProviderService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminAuthProviderPatchParams">AdminAuthProviderPatchParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthProviderResponse">AuthProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Secrets

Response Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#SecretResponse">SecretResponse</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminSecretListResponse">AdminSecretListResponse</a>

Methods:

- <code title="get /v1/admin/secrets">client.Admin.Secrets.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminSecretService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminSecretListResponse">AdminSecretListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/admin/secrets/{secret_id}">client.Admin.Secrets.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AdminSecretService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, secretID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Auth

Params Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthRequestParam">AuthRequestParam</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ConfirmUserRequestParam">ConfirmUserRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ConfirmUserResponse">ConfirmUserResponse</a>

Methods:

- <code title="post /v1/auth/authorize">client.Auth.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthAuthorizeParams">AuthAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/shared#AuthorizationResponse">AuthorizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/auth/confirm_user">client.Auth.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthService.ConfirmUser">ConfirmUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthConfirmUserParams">AuthConfirmUserParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ConfirmUserResponse">ConfirmUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/auth/status">client.Auth.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthStatusParams">AuthStatusParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/shared#AuthorizationResponse">AuthorizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#HealthSchema">HealthSchema</a>

Methods:

- <code title="get /v1/health">client.Health.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#HealthSchema">HealthSchema</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

Params Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ChatMessageParam">ChatMessageParam</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ChatRequestParam">ChatRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ChatMessage">ChatMessage</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ChatResponse">ChatResponse</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#Choice">Choice</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#Usage">Usage</a>

## Completions

Methods:

- <code title="post /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ChatResponse">ChatResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tools

Params Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#AuthorizeToolRequestParam">AuthorizeToolRequestParam</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ExecuteToolRequestParam">ExecuteToolRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ExecuteToolResponse">ExecuteToolResponse</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolDefinition">ToolDefinition</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolExecution">ToolExecution</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolExecutionAttempt">ToolExecutionAttempt</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ValueSchema">ValueSchema</a>

Methods:

- <code title="get /v1/tools">client.Tools.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolListParams">ToolListParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolDefinition">ToolDefinition</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/tools/authorize">client.Tools.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolAuthorizeParams">ToolAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/shared#AuthorizationResponse">AuthorizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/tools/execute">client.Tools.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolService.Execute">Execute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolExecuteParams">ToolExecuteParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ExecuteToolResponse">ExecuteToolResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tools/{name}">client.Tools.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolGetParams">ToolGetParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolDefinition">ToolDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Scheduled

Response Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolScheduledGetResponse">ToolScheduledGetResponse</a>

Methods:

- <code title="get /v1/scheduled_tools">client.Tools.Scheduled.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolScheduledService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolScheduledListParams">ToolScheduledListParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolExecution">ToolExecution</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/scheduled_tools/{id}">client.Tools.Scheduled.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolScheduledService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolScheduledGetResponse">ToolScheduledGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Formatted

Response Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolFormattedListResponse">ToolFormattedListResponse</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolFormattedGetResponse">ToolFormattedGetResponse</a>

Methods:

- <code title="get /v1/formatted_tools">client.Tools.Formatted.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolFormattedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolFormattedListParams">ToolFormattedListParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolFormattedListResponse">ToolFormattedListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/formatted_tools/{name}">client.Tools.Formatted.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolFormattedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolFormattedGetParams">ToolFormattedGetParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolFormattedGetResponse">ToolFormattedGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workers

Params Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#CreateWorkerRequestParam">CreateWorkerRequestParam</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#UpdateWorkerRequestParam">UpdateWorkerRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerHealthResponse">WorkerHealthResponse</a>
- <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerResponse">WorkerResponse</a>

Methods:

- <code title="post /v1/workers">client.Workers.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerNewParams">WorkerNewParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerResponse">WorkerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/workers/{id}">client.Workers.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerUpdateParams">WorkerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerResponse">WorkerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workers">client.Workers.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerListParams">WorkerListParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerResponse">WorkerResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/workers/{id}">client.Workers.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/workers/{id}">client.Workers.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerResponse">WorkerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workers/{id}/health">client.Workers.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerHealthResponse">WorkerHealthResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workers/{id}/tools">client.Workers.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerService.Tools">Tools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#WorkerToolsParams">WorkerToolsParams</a>) (<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go">arcadego</a>.<a href="https://pkg.go.dev/github.com/ArcadeAI/arcade-go#ToolDefinition">ToolDefinition</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
