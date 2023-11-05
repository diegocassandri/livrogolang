# Mocking Techniques

## When I need a mock?

### I need to mock a function

#### Higher Order Functions

- Good
- Clear and proximal to function under test
- Stateless
- Not so good
  - Parameter list can get ugly (using a type for the function can help)
  - Think of dependency graph

#### Monkey Patching

- Good
  - dont need to modify function signature
- not so good
  - Allergic to parallelism (stateful)
  - Have to make variable public (if testing from _test package)

### I need to mock a method on a type

#### Interface Substituition

- Accept interfaces, return structs
- We used io.ReadCloser, but dont be afraid to define you own interface

### I have a large interface, I need to mock a small set of its methods

#### Embedding

- Embedding can be a hole in the system

### I need to mock an HTTP call

#### net/http/httptest

- We made one test server per subtest, but you cold make one server with many routers, just like a normal server
