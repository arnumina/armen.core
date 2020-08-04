/*
#######
##       ___ _______ _  ___ ___       _______  _______
##      / _ `/ __/  ' \/ -_) _ \  _  / __/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package resources

import "time"

type (
	// Util AFAIRE.
	Util interface {
		LoggerPrefix(name, id string) string
		FileExist(file string) (bool, error)
		UnixToTime(timestamp string) time.Time
		EncryptString(text string) (string, error)
		DecryptString(text string) (string, error)
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
