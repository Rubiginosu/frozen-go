#!/usr/bin/env python
# coding=utf-8
# by Rubiginosu
# powered by Axoford12
# ------
# This is a function that processes socket infomation
# It will be run like a thread so that we can process more socket infomation
# ------

# Thread to process
# Conn is a param instance of socket-resource
def processThread (conn):
    conn.send('display|Hello,Welcome to FrozenGo Daemon|by Rubiginosu')

    # infinite loop
    while True:
        data = conn.recv(1024)
        # receive data from client
        data = data.split('|')
        # data will be split with string '|'
        if (data[0] != 'WithToken' or data[0] != 'BakdMsg'):
            conn.close()
            # Verify data with token if token not match
            # If token match faild , processor will consider this socekt 
            # client wth a hacker andd refuse this program.
            # Processor will close this socket
        # TODO 2017.0326 STOP HERE
