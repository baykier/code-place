import requests 

def get(url = ''):
    print(url)
    print('init url')
    r = requests.get('https://www.4455vs.com/tupian/list-%E4%BA%9A%E6%B4%B2%E5%9B%BE%E7%89%87.html')
   
    if(r.status_code == 200):
        print('success request is ' + str(r.status_code))
get(url = 'baidu.com')