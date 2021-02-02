# #Руководство по использованию fishingServer
## Библиотеки
* gopkg.in/yaml.v3
## Конфигурационный файл
Конфигурационный файл указывается как первый аргумент командной строки.
Файл должен быть в формате .yaml.
```
serverParams:
  externalServer: http://external.com
  serverPort: 4444
  pathToPic: public/images/picture.jpg
  pathToHTML: public/html/test.html
```

* pathToPic - путь до картинки, используемой в почтовом сообщении
* pathToHTML - путь до статической страницы
