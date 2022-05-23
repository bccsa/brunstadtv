// This package accepts messages from PUB-SUB and processes them based on certain rules
// The messages it accepts should be in the https://cloudevents.io/ format
package main

import (
	"context"

	awsSDKConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/bcc-code/brunstadtv/backend/cmd/jobs/server"
	"github.com/bcc-code/brunstadtv/backend/directus"
	"github.com/bcc-code/brunstadtv/backend/utils"
	"github.com/bcc-code/mediabank-bridge/log"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()

	log.ConfigureGlobalLogger(zerolog.DebugLevel)
	log.L.Debug().Msg("Seting up tracing!")

	// Here you can get a tracedHttpClient if useful anywhere
	tracedHTTPClient := utils.MustSetupTracing()
	ctx, initTrace := trace.StartSpan(ctx, "init")

	config := getEnvConfig()

	serverConfig := server.ConfigData{
		IngestBucket:       config.AWS.IngestBucket,
		StorageBucket:      config.AWS.StorageBucket,
		PackagingGroupID:   config.AWS.PackagingGroupARN,
		MediapackageRole:   config.AWS.MediapackageRoleARN,
		MediapackageSource: config.AWS.MediapackageSourceARN,
	}

	awsConfig, err := awsSDKConfig.LoadDefaultConfig(ctx)
	if err != nil {
		// TODO: Better messages
		panic(err)
	}

	awsConfig.Region = config.AWS.Region

	s3Client := s3.NewFromConfig(awsConfig)
	mediaPackageVOD := mediapackagevod.NewFromConfig(awsConfig)

	debugDirectus := false
	directusClient := directus.New(config.Directus.BaseURL, config.Directus.Key, debugDirectus, tracedHTTPClient)

	log.L.Debug().Msg("Set up HTTP server")
	router := gin.Default()

	handlers := server.NewServer(server.ExternalServices{
		S3Client:        s3Client,
		MediaPackageVOD: mediaPackageVOD,
		DirectusClient:  directusClient,
	}, serverConfig)

	apiGroup := router.Group("api")
	{
		apiGroup.POST("message", handlers.ProcessMessage)
	}

	initTrace.End()
	router.Run(":" + config.Port)
}
