# docker-compose-task

*Информация о настройке и Volume-ах*

0) /go-ethereum. К сожалению поднимаю только одну ноду. И с синхронизацией light на MAINNET.
Сделал так для того что бы проверить свои написанные приложения. Доделать по заданию в срок не успеваю. Буду еще разбираться.

      Заускаю с параметрами --syncmode light --rpc --rpcaddr 0.0.0.0 --nousb

1) rabbitMq. Конфижится через env файл. Файл расположен по пути: configs/rabbitmq/.env.rmq
Так же маунтятся следующие Volume-мы
      - ./.docker/rabbitmq/etc/:/etc/rabbitmq/
      - ./.docker/rabbitmq/data/:/var/lib/rabbitmq/
      - ./.docker/rabbitmq/logs/:/var/log/rabbitmq/

      Наружу смотрит панелька. Проброшен порт :15672 

2) Postgre. Конфижится через env файл. Файл расположен по пути: configs/postgresql/.env.pg_file 
Так же маунтятся следующие Volume-мы
      - ./.docker/postgresql/data:/var/lib/postgresql/data

3) redis. Маунтятся следующие Volume-мы
      - ./.docker/redis:/data

4) "block-sender". Приложение которое берет последний номер блока из блокчейна и пеерсылает его в rabbitMq
Конфижится через переменные окружения:
      - "RABBIT_CONN_STR=amqp://guest:guest@10.1.1.4:5672"
      - "NODE_CONN_STR=http://10.1.1.15:8545" - строка для подключения к ноде 
      - "INTERVAL=5" - интервал опроса ноды для получения последнего блока

      Имеет апи для healthcheck (порт 1488)

5) "block-receiver". Приложение которое берет номер блока из блокчейна, записывает его в базу, и инкрементит счетчик в редисе
Конфижится через переменные окружения:
      - "RABBIT_CONN_STR=amqp://guest:guest@10.1.1.4:5672"
      - "REDIS_HOST=10.1.1.2:6379"

      Имеет апи для healthcheck (порт 1488)
