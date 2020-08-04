/*
#######
##       ___ _______ _  ___ ___       _______  _______
##      / _ `/ __/  ' \/ -_) _ \  _  / __/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package resources

import "github.com/arnumina/logger"

type (
	// Container AFAIRE.
	Container interface {
		Util() Util
		Application() Application
		Logger() *logger.Logger
		Bus() Bus
		Leader() Leader
		Model() Model
		Server() Server
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
