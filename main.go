package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loglevel string
var Logger *zap.Logger = logger()
var logFile string = "logs.log"

// Sets up Cobra command line
var rootCmd = &cobra.Command{
	Use:   "validate [flags] [src]",
	Short: "A command-line tool for validating",
	Run: func(cmd *cobra.Command, args []string) {

		// Set loglevel for logger
		switch loglevel {
		case "INFO":
			Logger = Logger.WithOptions(zap.IncreaseLevel(zapcore.InfoLevel))
		case "DEBUG":
			Logger = Logger.WithOptions(zap.IncreaseLevel(zapcore.DebugLevel))
		case "ERROR":
			Logger = Logger.WithOptions(zap.IncreaseLevel(zapcore.ErrorLevel))
		default:
			Logger = Logger.WithOptions(zap.IncreaseLevel(zapcore.InfoLevel))
		}
	},
}

// ValidateLogLevel validates the log level
func ValidateLoglevel() error {
	switch loglevel {
	case "INFO", "DEBUG", "ERROR":
		return nil
	default:
		return errors.New("invalid log level. Allowed values are INFO, DEBUG, or ERROR")
	}
}

// ApplyExitOnHelp exits out of program if --help is flag
func ApplyExitOnHelp(c *cobra.Command, exitCode int) {
	helpFunc := c.HelpFunc()
	c.SetHelpFunc(func(c *cobra.Command, s []string) {
		helpFunc(c, s)
		os.Exit(exitCode)
	})
}

// logger creates logger with output of info and debug to file and error to stdout
func logger() *zap.Logger {
	pe := zap.NewDevelopmentEncoderConfig()

	fileEncoder := zapcore.NewJSONEncoder(pe)

	pe.EncodeTime = zapcore.ISO8601TimeEncoder // The encoder can be customized for each output

	// Create file core
	file, err := os.Create(logFile)
	if err != nil {
		panic(err)
	}

	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zap.DebugLevel)

	// Create a logger with two cores
	logger := zap.New(zapcore.NewTee(fileCore), zap.AddCaller())

	return logger
}

// init initates flags
func init() {
	rootCmd.Flags().StringVarP(&loglevel, "loglevel", "", "INFO", "Log level (INFO, DEBUG, ERROR)")
}

func main() {
	ApplyExitOnHelp(rootCmd, 0)

	if err := rootCmd.Execute(); err != nil {
		Logger.Error("Error setting command line",
			zap.Error(err))
		fmt.Println("There was an error setting the command line")
		os.Exit(1)
	}
	fmt.Println("Hello world")

}
