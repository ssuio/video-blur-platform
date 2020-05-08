# VYSIONEER assignment, A Video Processing Service

## Usage

1. Setup environment HOST_DATA_DIR, SESSION_KEY, WEB PORT in `docker-compose.yaml`. 
2. Setup docker sock `/var/run/docker.sock:/var/run/docker.sock` in `docker-compose.yaml`.
3. Setup video features.
Download `face-recognition.tar.gz` and build it `docker build -t face-blur .`.
3. `./build`
4. Build vysioneer-assignment-server `./build`
5. Build vysioneer-assignment-app `cd ./app` `./build`
6. Spin up service, `docker-compose up -d` in project root dir.

## Backup / Migrate

1. Copy `DATA_DIR`


