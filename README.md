# Vektor
Vektor is a simple (somewhat fast) vector database written in Go.

## Installation
### From source
```bash
$ git clone https://github.com/bpradana/vektor.git
$ cd vektor
$ make install
$ make build
$ ./vektor
```
### Running with Make
```bash
$ git clone https://github.com/bpradana/vektor.git
$ cd vektor
$ make install
$ make start
```
### Running with Docker
```bash
$ git clone https://github.com/bpradana/vektor.git
$ cd vektor
$ docker build -t vektor:latest .
$ docker run -p 8080:8080 vektor
```

## Query
### Create a vector
```sql
CREATE <vector_name> <vector_value>
CREATE vector1 [[1,2,3],[4,5,6]]
```
### Read a vector
```sql
READ <vector_name>
READ vector1
```
### Update a vector
```sql
UPDATE <vector_name> <vector_value>
UPDATE vector1 [[1,2,3],[4,5,6]]
```
### Search a vector
```sql
SEARCH <COSINE|EUCLIDEAN|EUCLIDEAN_L2> <threshold> <vector_value>
SEARCH COSINE 0.5 [1,2,3]
```
### Delete a vector
```sql
DELETE <vector_name>
DELETE vector1
```

## Connecting to Vektor
### Using Python
```python
class Client:
    def __init__(self, host, port):
        self.host = host
        self.port = port

    def connect(self):
        self.sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.sock.connect((self.host, self.port))

    def close(self):
        self.sock.close()

    def send_request(self, request):
        self.sock.sendall(request.encode())

        response = self.sock.recv(1024).decode()
        return response

if __name__ == "__main__":
    client = Client("localhost", 8080)
    client.connect()

    request = "CREATE vector1 [[1,2,3],[4,5,6]]"
    response = client.send_request(request)
    print(response)
```
