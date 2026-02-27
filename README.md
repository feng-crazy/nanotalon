# nanotalon

An ultra-minimalist, security-focused replacement for OpenClaw that is constructed entirely from the ground up using the Go programming language (Golang). Designed as a personal AI assistant that supports multiple communication channels and integrates with various LLM providers.

## Table of Contents
- [Features](#features)
- [Architecture](#architecture)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Supported Channels](#supported-channels)
- [LLM Providers](#llm-providers)
- [Docker Deployment](#docker-deployment)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Multi-Channel Support**: Connect via WhatsApp, Telegram, Discord, Slack, Feishu, DingTalk, QQ, Email, and more
- **Multiple LLM Providers**: Integrated with Anthropic, OpenAI, OpenRouter, Groq, and other providers
- **Secure by Design**: Minimalist architecture with security-focused practices
- **Flexible Tools**: Built-in tools for file operations, web search, and system interaction
- **Scheduling System**: Built-in cron-like functionality for automated tasks
- **Persistent Memory**: Long-term and session-based memory management
- **Modular Architecture**: Easy to extend with custom tools and channels

## Architecture

The project follows a modular architecture with these core components:

- **Agent**: Core AI processing engine that handles user requests and LLM interactions
- **Channels**: Communication platform integrations (Telegram, Discord, Slack, etc.)
- **Providers**: LLM provider integrations (Anthropic, OpenAI, etc.)
- **Tools**: Built-in tools for file operations, web search, execution, etc.
- **Configuration**: Centralized configuration system
- **Memory**: Two-layer memory system (long-term and session-based)
- **Scheduler**: Cron-like job scheduling system
- **Commands**: CLI interface built with Cobra

## Installation

### Prerequisites
- Go 1.24+
- Git
- An API key from your preferred LLM provider (Anthropic, OpenAI, OpenRouter, etc.)

### Build from Source
```bash
# Clone the repository
git clone https://github.com/yourusername/nanotalon.git
cd nanotalon

# Install dependencies
go mod tidy

# Build the application
./build.sh

# Or build directly
go build -o bin/nanotalon ./cmd/main.go
```

## Configuration

### Quick Setup
Run the onboard command to initialize your configuration:
```bash
./bin/nanotalon onboard
```

### Manual Configuration
Create a configuration file at `~/.nanotalon/config.yaml`:

```yaml
agents:
  defaults:
    workspace: "~/nanotalon-workspace"
    model: "anthropic/claude-3-sonnet-20240229"
    max_tokens: 8192
    temperature: 0.1
    max_tool_iterations: 40
    memory_window: 100

channels:
  send_progress: true
  send_tool_hints: false

  # Telegram configuration
  telegram:
    enabled: false
    token: "your-telegram-bot-token"
    allow_from:
      - "user123"

  # Discord configuration
  discord:
    enabled: false
    token: "your-discord-bot-token"
    allow_from:
      - "user456"

  # Slack configuration
  slack:
    enabled: false
    bot_token: "xoxb-your-bot-token"
    app_token: "xapp-your-app-token"
    allow_from:
      - "U1234567890"

  # WhatsApp configuration
  whatsapp:
    enabled: false
    allow_from:
      - "+1234567890"

  # Feishu configuration
  feishu:
    enabled: false
    app_id: "your-feishu-app-id"
    app_secret: "your-feishu-app-secret"
    encrypt_key: ""
    verification_token: ""
    allow_from: []

  # Mochat configuration
  mochat:
    enabled: false
    base_url: "https://api.mochat.example.com"
    claw_token: "your-mochat-token"
    allow_from: []

providers:
  openrouter:
    api_key: "your-openrouter-api-key"
    api_base: ""
  anthropic:
    api_key: "your-anthropic-api-key"
    api_base: ""
  openai:
    api_key: "your-openai-api-key"
    api_base: ""

gateway:
  host: "0.0.0.0"
  port: 18790
  heartbeat:
    enabled: true
    interval_s: 1800

tools:
  web:
    search:
      api_key: "your-search-api-key"
      max_results: 5
  exec:
    timeout: 60
  restrict_to_workspace: false
  mcp_servers: {}
```

## Usage

### CLI Commands

```bash
# Show help
./bin/nanotalon --help

# Interact with the AI agent
./bin/nanotalon agent

# Send a single message to the agent
./bin/nanotalon agent -m "Hello, how can you help me?"

# Check channel status
./bin/nanotalon channels status

# Manage cron jobs
./bin/nanotalon cron --help

# Check system status
./bin/nanotalon status
```

### Interactive Mode
Start an interactive session with the AI:
```bash
./bin/nanotalon agent
```

### Channel Management
Enable/disable specific channels in your configuration and check their status:
```bash
./bin/nanotalon channels status
```

## Supported Channels

| Channel | Status | Configuration Required |
|---------|--------|----------------------|
| Telegram | ✅ Working | Bot Token |
| Discord | ✅ Working | Bot Token |
| Slack | ✅ Working | Bot Token + App Token |
| WhatsApp | ✅ Working | API Key |
| Feishu | ✅ Working | App ID + Secret |
| DingTalk | ✅ Working | Client ID + Secret |
| QQ | ✅ Working | App ID + Secret |
| Email | ✅ Working | IMAP/SMTP Credentials |
| Mochat | ✅ Working | Base URL + Token |

## LLM Providers

| Provider | Models | Configuration Required |
|----------|--------|----------------------|
| Anthropic | Claude series | API Key |
| OpenAI | GPT series | API Key |
| OpenRouter | Various | API Key |
| Groq | LLaMA, Mixtral | API Key |
| DeepSeek | DeepSeek series | API Key |
| ZhiPu | GLM series | API Key |
| DashScope | Qwen series | API Key |
| Gemini | Gemini series | API Key |
| Moonshot | Kimi series | API Key |

## Docker Deployment

### Building the Docker Image

```bash
# Build the image
docker build -t nanotalon .

# Run with custom configuration
docker run -d \
  --name nanotalon \
  -p 18790:18790 \
  -v ~/.nanotalon:/root/.nanotalon \
  -v ~/nanotalon-workspace:/workspace \
  nanotalon
```

### Using Docker Compose

Create a `docker-compose.yml`:

```yaml
version: '3.8'
services:
  nanotalon:
    build: .
    ports:
      - "18790:18790"
    volumes:
      - ~/.nanotalon:/root/.nanotalon
      - ~/nanotalon-workspace:/workspace
    environment:
      - WORKSPACE_DIR=/workspace
    restart: unless-stopped
```

Run with:
```bash
docker-compose up -d
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests for new functionality
5. Run tests (`go test ./...`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Security

The project is designed with security in mind:
- Minimalist architecture reduces attack surface
- Channel access can be restricted via allow lists
- Workspace restrictions can limit file system access
- Input validation and sanitization

For security issues, please contact the maintainers privately.

---

**Note**: This is a security-focused replacement for OpenClaw with enhanced privacy and control features.