This is a Weather Information System that give weather information about few famous cities of India.

Pre-requisites:
1. Golang 1.12
2. GORM `go get -u github.com/jinzhu/gorm`

3. API token for `http://api.openweathermap.org`

Steps to run:
1. Signup on `http://api.openweathermap.org`, get the API token and set the environment variable `APIKEY`

`export APIKEY=<API-TOKEN>` 
2. Simply run `run.sh` script. After that system is ready to handle requests.

Description:
1. This application runs as a local http sever at 8080 port and continously gathers data in the background from  `http://api.openweathermap.org` in goroutines.
2. It saves data in `mysql` container running on the localhost and the container port `3306` is mapped to port `3306`of localhost.
3. Weather information for a particular city can be queried in 3 ways

    a. with `city name`

    b. with `latitute` and `longitude`

    c. with `zip-code`
    
   So the applicatin exposes 3 endpoints for the user:
   
   a. http://localhost:8080/weather/?city=`cityname`
   
   b. http://localhost:8080/weather/?lat=`latitude`&lon=`longitude`
   
   c. http://localhost:8080/weather/?zip=`zip-code`
   
Note: There is a `.http` file in this project to make sample API calls to sever