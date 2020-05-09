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

## Production Usage

### Register

1. New user can register a new account in  `https://web.ezsofa.com/entry` [Entry](https://web.ezsofa.com/entry).

### Login

1. New user can login by registered account in  `https://web.ezsofa.com/entry` [Entry](https://web.ezsofa.com/entry).

### Transfer Video for Face-blur Effect

1. After login, you will be redirect to dashboard.
2. Click left site navigate bar to <b>Transfer Page </b>.
3. Follow step1 to step3, you will complete a video uploaded and create a task to backend server.

### VideoList

1. Click left site navigate bar to <b>Video List Page </b>. 
2. The video just uploaded will get into <b>pending</b> state.
3. Video transfer task will be schedule to execute and turn into <b>processing</b> state.
4. Refresh video list to check the staus of the videos.
5. After processing the video state will turn into <b>doone</b>.
6. Switch public column to make video shareable.
7. Expand the row by clicking right side icon can see sharelink copy button.
8. Use the link to download video.
7. Tickle the row and tap delete to delete a video.





