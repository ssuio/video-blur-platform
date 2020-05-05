# VYSIONEER assignment, A Video Processing Service

## Usage

1. Setup environment DATA_DIR, SESSION_KEY, WEB PORT in `docker-compose.yaml`. 
2. Setup video features.
Download `face-recognition.tar.gz` and build it `docker build -t face-blur .`.
3. Spin up service, `docker-compose up -d`

### Preparaging

In docker-compose.yaml, make sure docker sock is mapped.

```
    /var/run/docker.sock:/var/run/docker.sock

```

## Backup / Migrate

1. Copy `DATA_DIR`

## Develop

### Build docker image

1. cd `./app`
2. Pack web app `npm run release`
3. cd ../server
4. Build docker image `docker build -t vysioneer-assignment --no-cache .`


