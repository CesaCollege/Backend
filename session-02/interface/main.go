package main

import "fmt"

type SaveDescriber interface {
	Save([]byte)
	Describe()
}

type Saver interface {
	Save([]byte)
}

type Describer interface {
	Describe()
}

type postgresDB struct {
	storage []byte
}

func NewPostgresDB() Saver {
	return &postgresDB{}
}

func (p *postgresDB) Save(d []byte) {
	// do some actual logic
	p.storage = append(p.storage, d...)
	fmt.Println("saved in postgres db!")
}

func (p *postgresDB) Describe() {
	fmt.Printf("PostgresDB: I have `%s` in my storage\n", string(p.storage))
}

func (p *postgresDB) SomeMethod() {
	// some actual logic
}

func NewMongoDB() Saver {
	return &mongoDB{}
}

type mongoDB struct {
	storage []byte
}

func (m *mongoDB) Save(d []byte) {
	// do some actual logic
	m.storage = append(m.storage, d...)
	fmt.Println("saved in mongo db!")
}

func (m *mongoDB) Describe() {
	fmt.Printf("MongoDB: I have `%s` in my storage\n", string(m.storage))
}

func NewRedisDB() Saver {
	return &redisDB{}
}

type redisDB struct {
	storage []byte
}

func (r *redisDB) Save(d []byte) {
	// do some actual logic
	r.storage = append(r.storage, d...)
	fmt.Println("saved in redis db!")
}

func (r *redisDB) Describe() {
	fmt.Printf("RedisDB: I have `%s` in my storage\n", string(r.storage))
}

func (r *redisDB) AnotherMethod() {
	// some actual logic
}

func main() {
	pg := NewPostgresDB()
	mongo := NewMongoDB()
	redis := NewRedisDB()
	storages := []Saver{pg, mongo, redis}

	// let's say we have some data
	data := []byte("some data")

	for _, s := range storages {
		s.Save(data)
	}

	for _, s := range storages {
		// casting or type assertion
		d, ok := s.(Describer)
		if ok {
			d.Describe()
		} else {
			fmt.Printf("failed to cast %s to Describer\n", s)
		}
	}

	for _, s := range storages {
		sd, ok := s.(SaveDescriber)
		if ok {
			// sd will have both Describe and Save methods available
		} else {
			fmt.Printf("failed to cast %s to SaveDescriber\n", sd)
		}
	}
}
