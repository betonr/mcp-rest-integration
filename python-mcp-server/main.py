from fastmcp import FastMCP
import httpx

API = "http://localhost:8080" # Go API URL

mcp = FastMCP("Orders-Go")

@mcp.tool(description="Create a new order")
async def create_order(customer: str, product: str, quantity: int):
    async with httpx.AsyncClient() as client:
        response = await client.post(f"{API}/orders", json={
            "customer": customer,
            "product": product,
            "quantity": quantity
        })
    return response.json()

@mcp.tool(description="Get an order by ID")
async def get_order(id: int):
    async with httpx.AsyncClient() as client:
        response = await client.get(f"{API}/orders/{id}")
    return response.json()

if __name__ == "__main__":
    mcp.run(transport="streamable-http", port=9000)