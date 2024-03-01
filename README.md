## LRU CACHE

### Description

Least Recently Used cache is referred to as LRU cache. Computer systems employ this kind of cache to hold data in memory so that subsequent requests for that data can be fulfilled more quickly. An LRU cache's main function is to monitor the order in which data is accessed and, upon reaching capacity, remove the least recently used items first to create space for newly added data.

### Steps to clone

- `git clone git@github.com:ganeshardlkar/apica_assignment.git`
- Navigate to `api/` and run `go mod tidy`
- Navigate to `client/` and run `npm install`
- Start the respective servers using `go run main.go` and `npm start`

### Screenshot/Recording

<video controls src="20240301-0506-49.4488199.mp4" title="Title"></video>

### API endpoints exposed

- `http://localhost:8080/get?key=${key}`
- `http://localhost:8080/set?key=${key}&value=${value}`

### Additional Info

- Maximum capacity of cache - 1024
- Expiration time - 10 seconds
