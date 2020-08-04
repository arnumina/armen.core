/*
#######
##       ___ _______ _  ___ ___       _______  _______
##      / _ `/ __/  ' \/ -_) _ \  _  / __/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package resources

import "github.com/gorilla/mux"

type (
	// Server AFAIRE.
	Server interface {
		Port() int
		NewRouter(prefix string) *mux.Router
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
