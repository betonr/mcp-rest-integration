from agents import Agent, Runner
from agents.mcp.server import MCPServerStreamableHttp
import asyncio

async def main():
    mcp_server = MCPServerStreamableHttp(
        params={"url": "http://localhost:9000/mcp/"}
    )
    await mcp_server.connect()

    agent = Agent(
        name="OrderBot",
        instructions="Use MCP tools to manage orders.",
        mcp_servers=[mcp_server],
        model="gpt-4o-mini",
    )

    result = await Runner.run(
        agent,
        "Create an order for customer Carlos buying 3 black shirts"
    )
    print(result.final_output)

    await mcp_server.cleanup()

if __name__ == "__main__":
    asyncio.run(main())