# Common Utils

A shared utility library for microservices in the wkstudio project.

## Features

- **Context Utilities**: Extract user ID from JWT context
- **Common Error Handling**: Shared error definitions
- **Validation Utilities**: Common validation functions

## Usage

### Get User ID from Context

```go
import "common-utils/commonutils"

// In your logic function
userID, err := commonutils.GetUserIDFromContext(ctx)
if err != nil {
    return nil, err
}
```

## Installation

In your microservice's `go.mod`, add:

```go
require common-utils v0.0.0

replace common-utils => ../common-utils
```

Then run:
```bash
go mod tidy
```

## Available Functions

### `GetUserIDFromContext(ctx context.Context) (int64, error)`

Extracts user ID from JWT context. Handles multiple type conversions:
- `float64` (from JSON unmarshaling)
- `json.Number`
- `int64`
- `int`
- `int32`

Returns:
- `userID`: The user ID as int64
- `error`: "unauthorized" if user_id not found, "invalid user ID" if type conversion fails

## License

MIT
