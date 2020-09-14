# go-app-grpc-chat-backend

Backend for an example chat application using the `go-app` package and gRPC over WebSockets (like gRPC-Web). See [pojntfx/go-app-grpc-chat-frontend-web](https://github.com/pojntfx/go-app-grpc-chat-frontend-web) for the example frontend.

![Go CI](https://github.com/pojntfx/go-app-grpc-chat-backend/workflows/Go%20CI/badge.svg)

## Overview

`go-app-grpc-chat-backend` is a short example showing the use of the [go-app](https://github.com/maxence-charriere/go-app) package and [gRPC](https://grpc.io/). This is the backend, which echoes messages back to the frontend and add's the server's timestamp to them.

## Installation

### Prebuilt Binaries

Linux, macOS and Windows binaries are available on [GitHub Releases](https://github.com/pojntfx/go-app-grpc-chat-backend/releases).

### Go Package

A Go package [is available](https://pkg.go.dev/mod/github.com/pojntfx/go-app-grpc-chat-backend).

## Usage

```bash
% go-app-grpc-chat-backend -help
Usage of go-app-grpc-chat-backend:
  -laddr string
        Listen address. (default "0.0.0.0:1925")
  -wsladdr string
        Listen address (for the WebSocket proxy). (default "0.0.0.0:10000")
```

## License

go-app-grpc-chat-backend (c) 2020 Felicitas Pojtinger

SPDX-License-Identifier: AGPL-3.0
