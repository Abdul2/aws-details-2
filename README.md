[![Build Status](http://138.197.88.141/api/badges/Abdul2/awsdetails/status.svg)](http://138.197.88.141/Abdul2/awsdetails)



---

### intro

this projects aims to do 3 things

    1 develop a GO application to report on aws resources

    2 build a docker container to host the Go application

    3 stand-up a test infra to auatomate the build and publishing

### files

    awsdetails.go

    Go application that queries aws account and write output into a json file.

    docker-compose.yml

    Build integration services to link github/drone/quay (this is experimental and will use official work platform

    .drone.yml

    drone config to make use of above integration

    secrets.yml

    this only shows the structure. the file will be parsed by drone locally to injuct encrypted secrets into drone publishing into quay.


