package main

import (
	"context"
	"encoding/base64"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})).With("module", "aws-kms")

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	client := kms.NewFromConfig(cfg)

	keyId := "alias/local-kms-key"
	plaintext := "Hello, world!"
	log.Info("Plaintext: " + plaintext)

	// Encrypt
	encryptOutput, err := client.Encrypt(ctx, &kms.EncryptInput{
		KeyId:     aws.String(keyId),
		Plaintext: []byte(plaintext),
	})
	if err != nil {
		log.Error("Encrypt failed", "err", err.Error())
		os.Exit(1)
		return
	}
	encrypted := base64.StdEncoding.EncodeToString(encryptOutput.CiphertextBlob)
	log.Info("Encrypted: " + encrypted)

	// Decrypt
	decryptOutput, err := client.Decrypt(ctx, &kms.DecryptInput{
		CiphertextBlob: encryptOutput.CiphertextBlob,
	})
	if err != nil {
		log.Error("Decrypt failed", "err", err.Error())
		os.Exit(1)
		return
	}
	decrypted := string(decryptOutput.Plaintext)
	log.Info("Decrypted: " + decrypted)
}
