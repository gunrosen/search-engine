# search-engine
Index data by using Elasticsearch
### 0. Configure .env
```
cp .env.example .env
# Configure env
vim .env
```
### 1. Start docker services

In development env, we use Kibana to get support from UI
```
 docker compose -f docker-compose.development.yml -f docker-compose.yml up -d --build
```
In production
```
 docker compose up -d --build 
```


### 2. Init Elasticsearch mapping type
```
docker exec -ti search /app/search init
```
And data will be ingested into Elasticsearch periodically by `CRON_JOB` .env