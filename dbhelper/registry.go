package dbhelper

type Registry struct {
	engine   *Engine
	entities map[string]*Entity
	models   map[string]IModel
}

func NewRegistry(engine *Engine) *Registry {
	r := new(Registry)
	r.engine = engine
	r.entities = make(map[string]*Entity)
	r.models = make(map[string]IModel)
	return r
}

func (r *Registry) Query(modelName string) *Query {
	var model IModel
	model, ok := r.models[modelName]
	if !ok {
		model = NewModel(modelName)
		model.SetRegistry(r)
	}
	return NewQuery(model)
}
