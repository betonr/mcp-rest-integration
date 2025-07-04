# 🤖 Go API with MCP Integration via FastMCP

This project demonstrates how to make a simple Go-based REST API compatible with AI agents using the [Model Context Protocol (MCP)](https://openai.com/blog/mcp), through a lightweight bridge implemented in Python with [FastMCP](https://pypi.org/project/fastmcp/).

✅ Ideal for developers looking to integrate existing services with AI models like GPT-4o.


## 📦 Project Structure

```txt
.
├── api-go/                # Go-based REST API for managing customer orders
├── python-mcp-server/     # FastMCP Python server that exposes tools for the Go API
└── agent/                 # Python agent that interacts with the MCP server via OpenAI's SDK
```

## What You’ll Learn

- How MCP works and how it enables natural language interfaces for your APIs.
- How to expose a REST API (in Go) to an AI agent using a Python bridge.
- How to test the end-to-end flow: prompt → MCP → Go API → real action.


## ▶️ How to Run

### 1. Run the Go API

```bash
cd api-go
go run main.go
```

This starts a simple REST API at http://localhost:8080 with endpoints:

- POST /orders — Create an order
- GET /orders/{id} — Retrieve an order


### 2. Start the MCP Server (Python + FastMCP)

```bash
cd python-mcp-server
pip3 install fastmcp httpx
python main.py
```

This exposes the API as MCP tools via HTTP on port 9000.
- Make sure the Go API is running first.

### 3. Connect an AI Agent

Set your OpenAI key and run the agent:
```bash
cd agent
pip3 install openai openai-agents openai-agents-mcp
export OPENAI_API_KEY=your_key_here
python run_agent.py
```

Example prompt:
```bash
Create a new order for customer Carlos buying 3 black shirts.
```

✅ The agent will call the Go API using FastMCP and return the response


## 📌 Notes
- The Go API is unaware of MCP – it's a plain RESTful service.
- FastMCP serves as a JSON-RPC 2.0 compliant MCP layer.
- You can replace FastMCP with any MCP-compliant server once available in Go.
