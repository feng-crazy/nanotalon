package channels

import (
	"fmt"
	"nanotalon/config"
)

// Channel represents a chat platform channel
type Channel interface {
	// Start starts the channel
	Start() error

	// Stop stops the channel
	Stop() error

	// Name returns the channel name
	Name() string

	// Send sends a message to the channel
	Send(chatID, message string) error
}

// Manager manages multiple channels
type Manager struct {
	channels map[string]Channel
	config   *config.Config
}

// NewManager creates a new channel manager
func NewManager(cfg *config.Config) *Manager {
	manager := &Manager{
		channels: make(map[string]Channel),
		config:   cfg,
	}

	// Initialize configured channels
	manager.initChannels()

	return manager
}

// initChannels initializes channels based on configuration
func (cm *Manager) initChannels() {
	// Initialize Telegram if enabled
	if cm.config.Channels.Telegram.Enabled {
		telegram := NewTelegramChannel(cm.config.Channels.Telegram.Token, cm.config.Channels.Telegram.AllowFrom)
		cm.Register(telegram)
	}

	// Initialize Discord if enabled
	if cm.config.Channels.Discord.Enabled {
		discord := NewDiscordChannel(cm.config.Channels.Discord.Token, cm.config.Channels.Discord.AllowFrom)
		cm.Register(discord)
	}

	// Initialize Slack if enabled
	if cm.config.Channels.Slack.Enabled {
		slack := NewSlackChannel(cm.config.Channels.Slack.BotToken, cm.config.Channels.Slack.AppToken, cm.config.Channels.Slack.AllowFrom)
		cm.Register(slack)
	}

	// Initialize Feishu if enabled
	if cm.config.Channels.Feishu.Enabled {
		feishu := NewFeishuChannel(
			cm.config.Channels.Feishu.AppID,
			cm.config.Channels.Feishu.AppSecret,
			cm.config.Channels.Feishu.EncryptKey,
			cm.config.Channels.Feishu.Verification,
			cm.config.Channels.Feishu.AllowFrom,
		)
		cm.Register(feishu)
	}

	// Initialize Mochat if enabled
	if cm.config.Channels.Mochat.Enabled {
		mochat := NewMochatChannel(
			cm.config.Channels.Mochat.BaseURL,
			cm.config.Channels.Mochat.ClawToken,
			cm.config.Channels.Mochat.AllowFrom,
		)
		cm.Register(mochat)
	}

	// Initialize DingTalk if enabled
	if cm.config.Channels.DingTalk.Enabled {
		dingtalk := NewDingTalkChannel(
			cm.config.Channels.DingTalk.ClientID,
			cm.config.Channels.DingTalk.Secret,
			cm.config.Channels.DingTalk.AllowFrom,
		)
		cm.Register(dingtalk)
	}

	// Initialize Email if enabled
	if cm.config.Channels.Email.Enabled {
		email := NewEmailChannel(&cm.config.Channels.Email)
		cm.Register(email)
	}

	// Initialize QQ if enabled
	if cm.config.Channels.QQ.Enabled {
		qq := NewQQChannel(
			cm.config.Channels.QQ.AppID,
			cm.config.Channels.QQ.Secret,
			cm.config.Channels.QQ.AllowFrom,
		)
		cm.Register(qq)
	}

	// Initialize WhatsApp if enabled
	if cm.config.Channels.WhatsApp.Enabled {
		waConfig := &WhatsAppConfig{
			Enabled:   cm.config.Channels.WhatsApp.Enabled,
			AllowFrom: cm.config.Channels.WhatsApp.AllowFrom,
		}
		whatsapp := NewWhatsAppChannel(waConfig)
		cm.Register(whatsapp)
	}
}

// Register registers a channel
func (cm *Manager) Register(channel Channel) {
	cm.channels[channel.Name()] = channel
}

// Get returns a channel by name
func (cm *Manager) Get(name string) (Channel, bool) {
	channel, exists := cm.channels[name]
	return channel, exists
}

// StartAll starts all registered channels
func (cm *Manager) StartAll() error {
	for name, channel := range cm.channels {
		if err := channel.Start(); err != nil {
			return fmt.Errorf("failed to start channel %s: %w", name, err)
		}
	}
	return nil
}

// StopAll stops all registered channels
func (cm *Manager) StopAll() error {
	var lastErr error
	for name, channel := range cm.channels {
		if err := channel.Stop(); err != nil {
			lastErr = fmt.Errorf("failed to stop channel %s: %w", name, err)
		}
	}
	return lastErr
}

// SendToChannel sends a message to a specific channel
func (cm *Manager) SendToChannel(channelName, chatID, message string) error {
	channel, exists := cm.Get(channelName)
	if !exists {
		return fmt.Errorf("channel %s not found", channelName)
	}

	return channel.Send(chatID, message)
}

// GetEnabledChannels returns a list of enabled channel names
func (cm *Manager) GetEnabledChannels() []string {
	var enabled []string
	for name := range cm.channels {
		enabled = append(enabled, name)
	}
	return enabled
}