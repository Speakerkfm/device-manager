## Device manager
### Установка

1. Склонировать проект
    ```sh
    go get github.com/Speakerkfm/device-manager
    ```
    
2. Установить зависимости
    ```sh
    cd $GOPATH/src/github.com/Speakerkfm/device-manager
    brew install dep
    dep ensure
    ```
    
3. Настроить конфиг
    ```sh
    cp .env.dist .env
    ```
    
4. Запустить тесты
    ```sh
    make test
    ```
    
5. Запустить API локально
    ```sh
    make start
    ```
    
### Документация 
    make doc
 
 ### Запуск в Docker 
     docker run -d -p 8080:8080 speakerkfm/device-manager
