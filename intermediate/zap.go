package intermediate

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		log.Println("Error in intializing Zap logger")
	}

	// logger, err := zap.NewProduction()
	// if err != nil {
	// 	log.Println("Error in intializing Zap logger")
	// }

	// zap logger may contain buffered logs, so flush them before the program exits
	defer logger.Sync()

	logger.Info("This is an info message.")

	logger.Info("User logged in", zap.String("username", "John Doe"), zap.String("method", "GET"))
}
