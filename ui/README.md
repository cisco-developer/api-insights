# API-Insights-UI

## How to run the Backend part that the UI part depends on it.
### 1. Enter the root folder of the git repo
### 2. Use the following command in a terminal
```
docker-compose up mysql
```
### 3. Use the following command in another terminal after step 2
```
docker-compose up backend
```
After that, the backed is running. and we can run UI part.

## How to run UI part
### 1. Enter /ui folder, use command in another terminal:
```
cd ui
```
### 2. Use the following commands (recommend to use node v18.10.0)
```
npm install
npm start
```
After that, UI code is ready for use and development.
If there are any porxy errors thrown by npm start, please try with node v16 or change proxy address in package.json from `"proxy": "http://localhost:8081/",` to `"proxy": "http://0.0.0.0:8081/",`
