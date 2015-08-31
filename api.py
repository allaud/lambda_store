from bottle import route, run, template
from os import listdir, getpid

METHODS = {}
FUNC_LIST = []

for lang in listdir('src'):
    for func in listdir("src/" + lang):
        func_name, _ = func.split('.')
        name = lang + "/" + func_name

        FUNC_LIST.append(name)
        METHODS[name] = open("src/" + lang + "/" + func).read()


@route('/')
def index():
    return {'functions': len(FUNC_LIST)}

@route('/list')
def index():
    return {'functions': FUNC_LIST}

@route('/func/<lang>/<method>')
def name(lang, method):
    return {'func': METHODS[lang + '/' + method]}

with open('pid', 'w') as f:
        f.write(str(getpid()))

run(host='localhost', port=8122)
