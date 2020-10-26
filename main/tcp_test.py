import socket
import time
import json

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(('127.0.0.1', 8080))    

file = open("ip2.log")
for ip in file:
	print(ip)
	sock.send(str.strip(ip)) 
	# ipData = json.loads(sock.recv(1024))
	# print(ipData)
	print(sock.recv(1024))  
	# time.sleep(.1)

file.close()

sock.close()

print("done!")