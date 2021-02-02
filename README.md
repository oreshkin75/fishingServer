# #Руководство по использованию fishingServer
## Библиотеки
* gopkg.in/yaml.v3
## Конфигурационный файл
Конфигурационный файл указывается как первый аргумент командной строки.
Файл должен быть в формате .yaml.
```
serverParams:
  externalServer: 
  serverPort:
  pathToPic:
  pathToHTML:
```

* pathToPic - путь до картинки, используемой для в почтовом сообщении
* pathToHTML - путь до статической страницы
