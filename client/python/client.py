import socket
import json
import time
 

start = time.time()

for num in range(1,10000):

    s = socket.socket()

    host = "127.0.0.1"
    # port = 1231
    port = 1234

    s.connect((host, port))

    # data = "{\"method\": \"Registry.GetIP\", \"params\": [\"xxx\"], \"id\": 1}"
    # s.send(data)

    # jsondata = s.recv(1024)

    # s.close()

    # text = json.loads(jsondata)
    # result = text['result']

    # host = result.split(":")[0];
    # port = int(result.split(":")[1]);

    # s = socket.socket()

    # s.connect((host, port))

    data = "{\"method\": \"User.GetUser\", \"params\": [\"xxx\"], \"id\": 1}"
    s.send(data)


    # print(s.recv(1024))

    s.close()

end = time.time()


spend = end - start

print(10000/(spend))

