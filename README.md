# Creamsensation - DI
## New()
Creates new dependency injection container.
```go
New(
  register deps...
)
```
## Register()
Dependency wrapper, which accept value reference like &testService{} and configurations.
```go
Register(&testService{}, Singleton())
Register(&testService{}, Name("test"), Singleton())
Register(&testService{}, Interface[testInterface](), Singleton())
```
### Dependency
Must have method Provide(Container), which returns singleton or new instance of dependency, where you can provide dependencies to it. <br>
```go
type testService struct {
  someRepository  *someRepository
  someService     someServiceInterface
}

func (s *testService) Provide(c di.Container) *testService {
  return &testService{
    someRepository: Provide[someRepository](c),
    someService:    Provide[someService](c),
  }
}
```
## Provide()
Dependency provider, which provide registered dependencies with type, name, or interface, must have type parameter, as argument you need to pass container.
```go
c := New(deps...)
Provide[testService](c)
Provide[testService](c, Name("test"))
Provide[testService](c, Interface[testInterface]())
```
## CreateProvider()
When you need some third-party things like db, cache, etc.. use with DI, you can do it with CreateProvider() wrapper.
```go
Register[sql.DB](CreateProvider(dbConnection), di.Singleton()),
```