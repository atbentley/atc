package resourceserver

import (
	"encoding/json"
	"net/http"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/api/present"
	"github.com/concourse/atc/auth"
	"github.com/concourse/atc/db"
)

func (s *Server) GetResource(pipelineDB db.PipelineDB) http.Handler {
	logger := s.logger.Session("get-resource")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config, _, found, err := pipelineDB.GetConfig()
		if err != nil {
			logger.Error("failed-to-get-config", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !found {
			logger.Info("config-not-found")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resourceName := r.FormValue(":resource_name")
		teamName := r.FormValue(":team_name")

		dbResource, found, err := pipelineDB.GetResource(resourceName)
		if err != nil {
			logger.Error("failed-to-get-resource", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !found {
			logger.Debug("resource-not-found", lager.Data{"resource": resourceName})
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resource := present.Resource(
			dbResource,
			config.Groups,
			auth.IsAuthenticated(r),
			teamName,
		)

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(resource)
	})
}
