# rest-api-generator

Provide Go program to generate MOCK API. 

## Usage
  Download : `git clone https://github.com/mohit83k/rest-api-generator.git`
  
  Build : `go build main.go`
  
  Run : `./main`
  

## Configuration
  `server`  : Server Related Details
  
  `apiRoutes` : Each json object have a `url` and `file` name . When `url` is hit, it serves the specified file. 
  Example:
  ```
  {
    "url" : "/hello",
    "file" : "hello.html"
  }
  ```
  `mockDir` : Folder to contain the mock files. In our example, `hello.html` should be put there. If not configured, file will be searched for at absolute path. 
  
  
 ## Optional:
     
  Command line args:
    `config` : Provide location of configuration file. Deafult is `./config.json` 
    
    Help : `./main -h`
 
  
  
