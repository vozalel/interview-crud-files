# interview-crud-files

Тестовое задание 

Есть две сущности — пользователи и источники данных 
(пусть, без ограничения общности, это будут csv-файлы). 
У каждого пользователя для каждого файла может быть доступ на чтение, 
на чтение/изменение или не быть доступа вообще. 
Так же у части пользователей может быть право на создание новых источников. 

Задачи:

1) Реализовать схему хранения этой информации (пользователи + файлы + права доступа).

2) Реализовать часть АПИ, которое позволяет:
— создать новый файл.
— прочитать существующий файл.
— изменить существующий файл.
— удалить существующий файл.

Можно считать что файлы лежат на сервере в заранее известной директории.
Можно считать, что авторизация уже реализована, и у нас уже есть необходимая информация (например, id пользователя)
