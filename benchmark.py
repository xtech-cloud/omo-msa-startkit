import json
import time
import urllib.request
from threading import Thread

url = 'http://127.0.0.1:8080/msa/startkit/Echo/Call'
params = {
        'msg':'hello',
        }
options = {
        'content-type':'application/json;charset=UTF-8'
        }

def call(_id, _n):
    for i in range(0, _n):
        data = json.dumps(params)
        data = bytes(data, 'utf8')
        request = urllib.request.Request(url, data, options)
        try:
            reply = urllib.request.urlopen(request).read().decode('utf-8')
            #print(reply)
        except Exception as e:
            print(e)
    print('thread %d finish'%(_id))

# 创建100个线程
for i in range(0,100):
    # 每个线程10次访问
    t = Thread(target=call, args=(i, 10,))
    t.start()
