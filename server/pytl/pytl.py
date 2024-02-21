import os

from bs4 import BeautifulSoup

insert = ""

DEFAULT = f'''
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Bootstrap demo</title>
    <link href="https://cdn.jsdelivr.net/npm/bootswatch@5.3.2/dist/cerulean/bootstrap.min.css" rel="stylesheet">
    <link rel="stylesheet" href="style/embed.css">
  </head>
  <body>
    {insert}
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL" crossorigin="anonymous"></script>
    <script src="js/embed.js"></script>
  </body>
</html>
'''

file = os.chdir('../public/home/')
print(file)

if __name__ == '__main__':
    print(file)