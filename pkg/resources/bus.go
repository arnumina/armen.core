/*
#######
##       ___ _______ _  ___ ___       _______  _______
##      / _ `/ __/  ' \/ -_) _ \  _  / __/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package resources

import "github.com/arnumina/armen.core/pkg/message"

type (
	// Bus AFAIRE.
	Bus interface {
		AddPublisher(name string, chCapacity, consumer int) chan<- *message.Message
		Subscribe(cb func(*message.Message), reList ...string) error
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
