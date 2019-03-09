# redis-desktop

# development
- build front

    ```
       > cd front
       > yarn install
       > yarn serve
    ```

- start

    ```
       > go run . -env=dev
    ```

# build


    ```
        > go run -tags generate gen.go
        
        > ./build-macos.sh
    ```