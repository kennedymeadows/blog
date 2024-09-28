from django.http import HttpResponse

def index(request):
    html = f'''
    <html>
        <head>
            <title>Simon Lewis</title>
        </head>
        <body>
            <h1>Welcome to the blog!</h1>
        </body>
    </html>
    '''
    return HttpResponse(html)