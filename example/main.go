package main

import (
	"fmt"
	"log"

	"nanotalon/agent"
	"nanotalon/config"
)

func main() {
	fmt.Println("_nanotalon: Ultra-Lightweight Personal AI Assistant_")

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		// If config doesn't exist, use defaults
		log.Printf("Could not load config: %v, using defaults", err)
		cfg = &config.Config{
			Agents: config.AgentsConfig{
				Defaults: config.AgentDefaults{
					Model: "anthropic/claude-opus-4-5",
				},
			},
			Providers: config.ProvidersConfig{
				OpenRouter: config.ProviderConfig{
					APIKey: "", // User needs to set this
				},
			},
		}
	}

	// Check if API key is configured
	if cfg.Providers.GetAPIKey("") == "" {
		fmt.Println("⚠️  Warning: No API key configured. Please set up your config file at ~/.nanotalon/config.yaml")
		fmt.Println("Get an API key at: https://openrouter.ai/keys")
		fmt.Println("Then run: nanotalon onboard")
		return
	}

	// Create agent
	agentLoop, err := agent.NewAgentLoop(cfg)
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	// Simple demo
	fmt.Println("\n🚀 nanotalon is ready!")
	fmt.Println("Try running: nanotalon agent -m \"Hello!\"")
	fmt.Println("Or: nanotalon agent # for interactive mode")

	// Use agentLoop to show it's referenced
	_ = agentLoop
}
