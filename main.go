// A Subako function package: an MCP server whose tools an agent can call.
package main

import (
	"context"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	subakofn "github.com/subako-ai/subako-function-sdk-go"
)

type GreetInput struct {
	Name string `json:"name" jsonschema:"who to greet"`
}

type Greeting struct {
	Message string `json:"message"`
}

func greet(ctx context.Context, req *mcp.CallToolRequest, in GreetInput) (*mcp.CallToolResult, Greeting, error) {
	if in.Name == "" {
		return nil, Greeting{}, errors.New("name must not be empty")
	}
	return nil, Greeting{Message: "Hello, " + in.Name + "!"}, nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "hello", Version: "0.1.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "Greet someone by name."}, greet)
	subakofn.Serve(server)
}
