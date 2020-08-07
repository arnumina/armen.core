/*
#######
##       ___ _______ _  ___ ___       _______  _______
##      / _ `/ __/  ' \/ -_) _ \  _  / __/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package resources

import (
	"time"

	"github.com/arnumina/logger"
)

type (
	// Util AFAIRE.
	Util interface {
		LoggerPrefix(name, id string) string
		CloneLogger(logger *logger.Logger, prefix string) *logger.Logger
		FileExist(file string) (bool, error)
		UnixToTime(timestamp string) time.Time
		DecodeData(input, output interface{}) error
		EncryptString(text string) (string, error)
		DecryptString(text string) (string, error)
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
