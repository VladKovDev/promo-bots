package app

import (
	"context"
	"fmt"
	"os"

	"github.com/VladKovDev/promo-bot/internal/config"
)

func Run(ctx context.Context) error {

	configPath := os.Getenv("PROMO_BOTS_CONFIG_PATH")
	cfg, err := InitConfig(configPath, ctx)
	if err != nil {
		return fmt.Errorf("failed to init config: %w", err)
	}
	_ = cfg
	fmt.Println("Successful config run")
	return nil
}

func InitConfig(configPath string, ctx context.Context) (*config.Config, error) {
	cfg, err := config.Load(configPath, ctx)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
