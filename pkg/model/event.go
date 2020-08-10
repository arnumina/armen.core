/*
#######
##       ___ _______ _  ___ ___       _______  _______
##      / _ `/ __/  ' \/ -_) _ \  _  / __/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package model

import "time"

type (
	// Event AFAIRE.
	// AFINIR: data ?
	Event struct {
		Name     string
		Disabled bool
		After    *string
		Repeat   *string
	}

	// Instance AFAIRE.
	Instance struct {
		ID        string
		Host      string
		Port      int
		StartedAt time.Time
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
