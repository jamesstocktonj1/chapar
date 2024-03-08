package manager

import (
	"errors"

	"github.com/mirzakhany/chapar/internal/domain"
	"github.com/mirzakhany/chapar/internal/loader"
)

var ErrNotFound = errors.New("ErrNotFound")

type Manager struct {
	currentActiveEnv *domain.Environment

	environments map[string]*domain.Environment
	collections  map[string]*domain.Collection
	requests     map[string]*domain.Request
}

func New() *Manager {
	return &Manager{
		environments: make(map[string]*domain.Environment),
		collections:  make(map[string]*domain.Collection),
		requests:     make(map[string]*domain.Request),
	}
}

func (m *Manager) SetCurrentActiveEnv(env *domain.Environment) {
	m.currentActiveEnv = env
}

func (m *Manager) GetCurrentActiveEnv() *domain.Environment {
	return m.currentActiveEnv
}

func (m *Manager) AddEnvironment(env *domain.Environment) {
	m.environments[env.MetaData.ID] = env
}

func (m *Manager) SetEnvironments(envs []*domain.Environment) {
	for _, env := range envs {
		m.environments[env.MetaData.ID] = env
	}
}

func (m *Manager) GetEnvironment(id string) *domain.Environment {
	if env, ok := m.environments[id]; ok {
		return env
	}
	return nil
}

func (m *Manager) GetEnvironmentFromDisc(id string) (*domain.Environment, error) {
	envs, err := loader.ReadEnvironmentsData()
	if err != nil {
		return nil, err
	}

	for _, env := range envs {
		if env.MetaData.ID == id {
			return env, nil
		}
	}

	return nil, ErrNotFound
}

func (m *Manager) ReloadEnvironmentFromDisc(id string) {
	env, err := m.GetEnvironmentFromDisc(id)
	if err != nil {
		return
	}

	m.AddEnvironment(env)
}

func (m *Manager) GetEnvironments() map[string]*domain.Environment {
	return m.environments
}

func (m *Manager) AddCollection(collection *domain.Collection) {
	m.collections[collection.MetaData.ID] = collection
}

func (m *Manager) GetCollection(id string) *domain.Collection {
	if collection, ok := m.collections[id]; ok {
		return collection
	}
	return nil
}

func (m *Manager) GetCollections() map[string]*domain.Collection {
	return m.collections
}

func (m *Manager) AddRequest(request *domain.Request) {
	m.requests[request.MetaData.ID] = request
}

func (m *Manager) GetRequest(id string) *domain.Request {
	if request, ok := m.requests[id]; ok {
		return request
	}
	return nil
}

func (m *Manager) GetRequests() map[string]*domain.Request {
	return m.requests
}

func (m *Manager) DeleteEnvironment(id string) {
	delete(m.environments, id)
}

func (m *Manager) DeleteCollection(id string) {
	delete(m.collections, id)
}

func (m *Manager) DeleteRequest(id string) {
	delete(m.requests, id)
}

func (m *Manager) UpdateEnvironment(env *domain.Environment) {
	m.environments[env.MetaData.ID] = env
}

func (m *Manager) UpdateCollection(collection *domain.Collection) {
	m.collections[collection.MetaData.ID] = collection
}

func (m *Manager) UpdateRequest(request *domain.Request) {
	m.requests[request.MetaData.ID] = request
}

func (m *Manager) Clear() {
	m.environments = make(map[string]*domain.Environment)
	m.collections = make(map[string]*domain.Collection)
	m.requests = make(map[string]*domain.Request)
}

func (m *Manager) LoadData() error {
	envs, err := loader.ReadEnvironmentsData()
	if err != nil {
		return err
	}

	for _, env := range envs {
		m.AddEnvironment(env)
	}

	collections, err := loader.LoadCollections()
	if err != nil {
		return err
	}

	for _, collection := range collections {
		m.AddCollection(collection)
	}

	requests, err := loader.LoadRequests()
	if err != nil {
		return err
	}

	for _, request := range requests {
		m.AddRequest(request)
	}

	return nil
}
