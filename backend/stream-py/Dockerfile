FROM ghcr.io/astral-sh/uv:python3.12-bookworm-slim
WORKDIR /app

COPY . .

RUN uv venv 
RUN uv sync 

ENV PATH="/app/.venv/bin:$PATH"


EXPOSE 8000
CMD ["fastapi", "run","main.py"]
