package ports

import (
	"encoding/json"
	"golang-project-template/internal/apigateway/app"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type IdRequest struct {
	ID int `json:"id"`
}

type HttpServer struct {
	service app.Service
}

func NewController(service app.Service) *HttpServer {
	return &HttpServer{
		service: service,
	}
}

func (c *HttpServer) CollectPostsHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := c.service.CollectPosts()
	if err != nil {
		http.Error(w, "error fetching data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)

}

// Create implements Service.
func (c *HttpServer) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var req app.NewPost

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	id, err := c.service.Create(&req)
	if err != nil {
		http.Error(w, "error creating a new post: "+err.Error(), http.StatusBadGateway)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

// Delete implements Service.
func (c *HttpServer) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var req IdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "error parsing request data: "+err.Error(), http.StatusBadGateway)
		return
	}
	msg, err := c.service.Delete(req.ID)
	if err != nil {
		http.Error(w, "error deleting the post: "+err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)

}

// GetById implements Service.
func (c *HttpServer) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intID, _ := strconv.Atoi(id)

	post, err := c.service.GetById(intID)
	if err != nil {
		http.Error(w, "error getting by id: "+err.Error(), http.StatusBadGateway)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

// Update implements Service.
func (c *HttpServer) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var data app.NewPost
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	post, err := c.service.Update(&data)
	if err != nil {
		http.Error(w, "error updating post object: "+err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
