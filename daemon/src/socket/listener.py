#!/usr/bin/env python
# coding=utf-8
import sys
import socket
# by Rubiginosu
# powered by Axoford12
# import socket lib
import thread
from processor import processThread
HOST = ''
PORT = 1212
try:
# create socket with AFINET STREAM (TCP)
    s = socket.socket(socket.AF_INET,socket.SOCK_STREAM);
except socket.error ,msg:
    print 'Failed to create socket,withh Error code:'+ str(msg[0]) + 'Error Message :' + msg[1]
    sys.exit()
print 'Socket created'
try:
    s.bind((HOST,PORT))
except socket.error, msg:
    print 'Bind port error Error Code : ' + str(msg[0]) + ' Message ' + msg[1]
    sys.exit()
print 'Socket bind complete'

s.listen(20) # Default listen to 20 client
while True:
    conn, addr = s.accept()
    print 'Client connected with ' + addr[0] + ':' + str(addr[1])
    thread.start_new_thread(processThread,(conn,))
s.close()
