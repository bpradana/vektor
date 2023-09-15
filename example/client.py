import socket
import random


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


def generate_key():
    alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    number = "0123456789"
    return f"{''.join(random.sample(alphabet, 3))}{''.join(random.sample(number, 4))}"


def generate_vectors(num_vectors=10, dimension=2662):
    return str(
        [[random.random() for _ in range(dimension)] for _ in range(num_vectors)]
    ).replace(" ", "")


if __name__ == "__main__":
    client = Client("localhost", 8080)
    client.connect()

    data_num = 10
    dimension = 1000
    num_vectors = 10
    option = "EUCLIDEAN_L2"  # "EUCLIDEAN", "COSINE", "EUCLIDEAN_L2"
    threshold = 0.9
    keys = [generate_key() for _ in range(data_num)]
    vectors = [
        generate_vectors(num_vectors=num_vectors, dimension=dimension)
        for _ in range(data_num)
    ]

    for key, vector in zip(keys, vectors):
        request = f"CREATE {key} {vector};"
        response = client.send_request(request)
        print(response)

    for key in keys:
        request = f"READ {key};"
        response = client.send_request(request)
        print(response)

    for key, vector in zip(keys, vectors):
        request = f"UPDATE {key} {vector};"
        response = client.send_request(request)
        print(response)

    for _ in range(data_num):
        request = f"SEARCH {option} {str(threshold)} {str([random.random() for _ in range(dimension)]).replace(' ', '')};"
        response = client.send_request(request)
        print(response)

    for key in keys:
        request = f"DELETE {key};"
        response = client.send_request(request)
        print(response)
