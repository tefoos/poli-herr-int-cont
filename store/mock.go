package store

type MockStore struct {
	data map[string][]string
}

func NewMockStore() *MockStore {
	return &MockStore{data: make(map[string][]string)}
}

func (m *MockStore) Push(key, value string) error {
	m.data[key] = append([]string{value}, m.data[key]...)
	return nil
}

func (m *MockStore) Pop(key string) (string, error) {
	if len(m.data[key]) == 0 {
		return "", nil
	}
	value := m.data[key][0]
	m.data[key] = m.data[key][1:]
	return value, nil
}

func (m *MockStore) Enqueue(key, value string) error {
	m.data[key] = append(m.data[key], value)
	return nil
}

func (m *MockStore) Dequeue(key string) (string, error) {
	if len(m.data[key]) == 0 {
		return "", nil
	}
	value := m.data[key][0]
	m.data[key] = m.data[key][1:]
	return value, nil
}

func (m *MockStore) Append(key, value string) error {
	m.data[key] = append(m.data[key], value)
	return nil
}

func (m *MockStore) Remove(key string) (string, error) {
	if len(m.data[key]) == 0 {
		return "", nil
	}
	value := m.data[key][0]
	m.data[key] = m.data[key][1:]
	return value, nil
}

func (m *MockStore) GetState(key string) ([]string, error) {
	return m.data[key], nil
}
