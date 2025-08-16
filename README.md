# go-session-manager

> [!CAUTION]
> This project is mainly used in my private projects and might not be tested fully.
> It is not meant to be used in productive scenarios (yet).
> Use at your own risk!

Session management library for Go, providing an easy-to-use interface for starting, retrieving, and ending sessions with pluggable storage backends.

## Features
- In-memory session store implementation
- Simple API for session lifecycle management
- Context-based session handling

## Installation

```bash
go get -u github.com/Neatoro/go-session-manager
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/Neatoro/go-session-manager"
)

func main() {
	manager := gosessionmanager.NewInMemorySessionManager()
	ctx := context.Background()

	// Start a new session
	ctx, err := manager.StartSession(ctx)
	if err != nil {
		panic(err)
	}

	// End the session
	err = manager.EndSession(ctx)
	if err != nil {
		panic(err)
	}
}
```
