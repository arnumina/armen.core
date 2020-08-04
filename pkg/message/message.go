/*
#######
##       ___ _______ _  ___ ___       _______  _______
##      / _ `/ __/  ' \/ -_) _ \  _  / __/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package message

import "github.com/arnumina/uuid"

type (
	// Message AFAIRE.
	Message struct {
		ID        string
		Topic     string
		Data      interface{}
		Publisher string
	}
)

// New AFAIRE.
func New(topic string, data interface{}) *Message {
	return &Message{
		ID:    uuid.New(),
		Topic: topic,
		Data:  data,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
