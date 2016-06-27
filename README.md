# go-apidemo

Small piece of code demonstrating the separation of functionality in a web application:

+ HTML UI - static views
+ JS UI - dynamic content
+ API - access to dynamically loaded data

No presentation layer markup or JS should be contained within the API implementation

## Usage

````
    git clone https://github.com/sebcat/go-apidemo
    cd go-apidemo
    go run main.go
    open http://127.0.0.1:8080/ in web browser
````
