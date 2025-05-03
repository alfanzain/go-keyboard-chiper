package core

import (
	"context"
	"strings"

	"github.com/alfanzain/keyboard-chiper/internal/core/helper"
)

type Service interface {
	HandleDecode(ctx context.Context, input string) (output string, err error)
	// HandleEncode(ctx context.Context, input string) (output string, err error) // Future feat
	// HandleSave(ctx context.Context, decodeStr string, encodeStr string) (err error) // Future feat
}

type ServiceConfig struct {
	// Storage Storage `validate:"nonnil"`
}

func NewService(cfg ServiceConfig) (Service, error) {
	return &service{
		ServiceConfig: cfg,
	}, nil
}

type service struct {
	ServiceConfig
}

func (s *service) HandleDecode(ctx context.Context, input string) (output string, err error) {
	if len(input) == 0 {
		return "", nil
	}

	var decoded strings.Builder
	for i := range len(input) {
		var letter byte = input[i]

		decoded.WriteByte(helper.Shifter(letter, helper.Left))
	}

	return decoded.String(), nil
}
