from fastapi import FastAPI

app = FastAPI()

@app.get("/api/v1/health")
async def health():
    return {"status": "ok"}


# if __name__ == "__main__":
#     main()
