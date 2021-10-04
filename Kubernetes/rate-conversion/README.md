## Multi Stage App

This folder contains a sample application for orchestration which consists of a Go backend which connects to a public API via internet, and also to a redis cluster, which is assumed to live on some private network.

To run locally one may use docker compose, for which a docker-compose.yml file is provided. Use with:

```BASE_URL=<base-url> APP_ID=<app-id> docker-compose up --abort-on-container-exit```

Make sure to have an available 8080 port on your host machine, or change the binding to something available on the docker-compose.yml file.

For a clearer explanation of the required environment variables, see [app repo](app).
