package common

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

// PrintAsYAML prints a generic object to console using YAML codec.
func PrintAsYAML(comment, obj interface{}) {
	objRaw, err := yaml.Marshal(obj)
	if err != nil {
		log.Error().Err(err).Msgf("Marshalling (%T) to YAML failed", obj)
		return
	}

	log.Info().Msgf("%s\n%s", comment, string(objRaw))
}
