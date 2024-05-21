package handler

import (
	"context"
	"log"
	"net/http"
	"port-service/internal/core/domain"
)

// PortService is a port service
type PortService interface {
	CreateOrUpdatePort(ctx context.Context, port *domain.Port) error
}

// HttpServer is HTTP server for ports
type HttpServer struct {
	service PortService
}

// NewHttpServer creates a new HTTP server for ports
func NewHttpServer(service PortService) HttpServer {
	return HttpServer{
		service: service,
	}
}

func (h HttpServer) FetchPorts(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching ports")

	portChan := make(chan Port)
	errChan := make(chan error)
	doneChan := make(chan struct{})

	go func() {
		err := fetchPortsFromJSON(r.Context(), r.Body, portChan)
		if err != nil {
			errChan <- err
		} else {
			doneChan <- struct{}{}
		}
	}()
	for {
		select {
		case <-r.Context().Done():
			log.Printf("request cancelled")
			return
		case <-doneChan:
			log.Printf("finished reading ports")
			RespondOK(w, r)
			return
		case err := <-errChan:
			log.Printf("error while parsing port json: %+v", err)
			BadRequest(err, w, r)
			return
		case port := <-portChan:
			log.Printf("fetch port: %+v", port)
			p, err := portHttpToDomain(&port)
			if err != nil {
				BadRequest(err, w, r)
				return
			}
			if err := h.service.CreateOrUpdatePort(r.Context(), p); err != nil {
				log.Printf("error while saving: %+v", err)
				BadRequest(err, w, r)
				return
			}
		}
	}
}

func portHttpToDomain(p *Port) (*domain.Port, error) {
	return domain.NewPort(
		p.Id,
		p.Name,
		p.Code,
		p.City,
		p.Country,
		append([]string(nil), p.Alias...),
		append([]string(nil), p.Regions...),
		append([]float64(nil), p.Coordinates...),
		p.Province,
		p.Timezone,
		append([]string(nil), p.Unlocs...),
	)
}
