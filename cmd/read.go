package cmd

import (
	"bufio"
	"bytes"
	"os"
	"quake-log/internal/logreader"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a log file",
	Long:  `Read log file `,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: reader itself can be a loader, so we can
		// use the same command to read from file or other sources
		// like HTTP.
		source, err := cmd.Flags().GetString("source")
		if err != nil {
			log.Error().Err(err).Msg("error while reading flag source")
			return
		}
		// read file
		fl, err := os.ReadFile(source)
		if err != nil {
			log.Error().Err(err).Msg("error while reading file")
			return
		}

		// run our app
		engine := logreader.NewDefaultEngine()
		report, err := engine.Process(bufio.NewReader(bytes.NewBuffer(fl)))
		if err != nil {
			log.Error().Err(err).Msg("error while processing file")
			return
		}

		// show our result on console
		log.Info().Any("report", report).Msg("execution result")
	},
}

func init() {
	readCmd.Flags().String("source", "", "Source log file to read from")
}
