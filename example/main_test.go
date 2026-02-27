package main

import (
	"os"
	"testing"

	"nanotalon/agent"
	"nanotalon/config"
)

func TestMainFunction(t *testing.T) {
	// Test that the config loading works without panicking
	cfg, err := config.LoadConfig()
	if err != nil {
		// This is expected if no config file exists, just continue with defaults
		t.Logf("Config loading failed (expected): %v", err)
		// Create minimal config for testing
		cfg = &config.Config{
			Agents: config.AgentsConfig{
				Defaults: config.AgentDefaults{
					Model: "anthropic/claude-3-haiku-20240307", // Use a free/fast model for testing
				},
			},
			Providers: config.ProvidersConfig{
				Anthropic: config.ProviderConfig{
					APIKey: os.Getenv("ANTHROPIC_API_KEY"), // Will be empty if not set
				},
			},
		}
	}

	// Try to create an agent loop (this might fail if no API key is configured)
	agentLoop, err := agent.NewAgentLoop(cfg)
	if err != nil {
		t.Logf("Agent creation failed (expected if no API key): %v", err)
		// This is expected when no API key is configured
		return
	}

	// If agent was created, at least make sure it's not nil
	if agentLoop == nil {
		t.Error("Agent loop is nil")
	}
}