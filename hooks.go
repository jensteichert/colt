package colt

type HookType string
const (
	BeforeInsert HookType = "BeforeInsert"
)

type Hook[T Document] struct {
	Type HookType
	Exec func (doc T) error
}

func (repo *Collection[T]) AddHook(t HookType, exec func (doc T) error) {
	repo.hooks = append(repo.hooks, Hook[T]{Type: t, Exec: exec})
}

func (repo *Collection[T]) hooksOfType(t HookType) []Hook[T] {
	filtered := []Hook[T]{}

	for _, h := range repo.hooks {
		if h.Type == t {
			filtered =append(filtered, h)
		}
	}

	return filtered
}

func (repo *Collection[T]) execHooks(t HookType, doc T) error {
	hooks := repo.hooksOfType(t)

	for _, h := range hooks {
		err := h.Exec(doc)
		if err != nil {
			return err
		}
	}

	return nil
}