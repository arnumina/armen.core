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
	// Application AFAIRE.
	Application interface {
		ID() string
		Name() string
		Version() string
		BuiltAt() time.Time
		StartedAt() time.Time
		FQDN() string
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
