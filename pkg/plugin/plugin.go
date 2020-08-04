/*
#######
##       ___ _______ _  ___ ___       _______  _______
##      / _ `/ __/  ' \/ -_) _ \  _  / __/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package plugin

import (
	"time"

	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/logger"
)

type (
	// Plugin AFAIRE.
	Plugin interface {
		Name() string
		Version() string
		BuiltAt() time.Time
		Build() error
		RunJob(job *jw.Job, logger *logger.Logger) *jw.Result
		Close()
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
