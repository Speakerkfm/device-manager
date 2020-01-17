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
    
3. Запустить тесты
    ```sh
    make test
    ```
    
4. Запустить API локально
    ```sh
    make start
    ```
    
### Документация 
    ```sh
    make doc
    ```
 
 ### Запуск в Docker 
    ```sh
     docker run -d -p 8080:8080 --network=host speakerkfm/device-service
     ```
     