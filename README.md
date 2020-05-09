# VYSIONEER assignment, A Video Processing Service

## Usage

1. Setup data folder,  <b>HOST_DATA_DIR</b> in `docker-compose.yaml`. 
2. Setup api domain, <b>API_HOST</b> in `docker-compose.yaml.
2. Setup video features.
Download `face-recognition.tar.gz` and build it `docker build -t face-blur .`.4. Build vysioneer-assignment-server `./build`
5. Install web packages `cd ./app`, `npm install` with Nodejs 10.16.3.
6. Build vysioneer-assignment-app `cd ./app` `./build`
7. Spin up service, `docker-compose up -d` in project root dir. The web app will run on `http://localhost:8080`, and the server will run on `http://localhost:9000`.

## Backup / Migrate

1. Copy `DATA_DIR`


