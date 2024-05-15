# Go/Golang Clean Architecture

## Installation

```
$ git clone [the project name]

$ cd [the project name]

$ go get
```

## Run Development Mode

- make .env file in root directory
- ypu can make.env by rename .env.example
- In root directory, run this command
    ```
    $ go run main.go
    ```
- in requester (mis:postman) open localhost:3000/customers


## Run Production

```bash
docker build -t [the project name] .
docker run -d -p 3000:3000 [the project name]
```