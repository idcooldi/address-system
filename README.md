# address-system

## ФЕДЕРАЛЬНАЯ ИНФОРМАЦИОННАЯ АДРЕСНАЯ СИСТЕМА

Данная система позволяет с ресурса "ФЕДЕРАЛЬНАЯ НАЛОГОВАЯ СЛУЖБА" распарсивать адреса в базу данных для дальнейшей работы.

## Инструкция 
Шаг 1. https://fias.nalog.ru/Updates.aspx скачиваем архив с сайта и распаковываем. 

Шаг 2. В архиве несколько XML файлов выбираем нужный и переименовываем в example.xml

Шаг 3. Запускаем ```doceker-compose up -d``` (если нет [docker](https://docs.docker.com/install/#next-release) и [docker-compose](https://docs.docker.com/compose/install/) устанавливаем)

Шаг 4. http://localhost:8080/adminer.php  
```Settings
		Сервер:db
		Имя пользователя: root
		Пароль: secret
```
