/*
#######
##       ___ _______ _  ___ ___       _______  _______
##      / _ `/ __/  ' \/ -_) _ \  _  / __/ _ \/ __/ -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package resources

import "github.com/arnumina/armen.core/pkg/jw"

type (
	// Model AFAIRE.
	Model interface {
		PluginConfig(plugin string, config interface{}) error
		InsertJob(job *jw.Job) error
		MaybeInsertJob(job *jw.Job) (bool, error)
		InsertWorkflow(wf *jw.Workflow) error
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
