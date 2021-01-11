## Kubernetes Test Go App

### Usage

Before running the container the image must be built by running

```docker build -t kubertest:<tag> .```


Once the image has been built, any way of running a docker container can be used for deployment. A common way to test out locally is running it ad-hoc with docker. This can be achieved with the following command

```docker run -p 8080:8080 -e APP_ID=<app-id> -e BASE_URL=https://openexchangerates.org/api -e REDIS_ADDRESS=<redis-host> -e REDIS_PORT=<redis-port> kubertest:<tag>```

Keep in mind that you may need to spin up a network where a redis server/cluster lives and run the app on that network. Another alternative is running a redis server locally and using the bridge network, in which case you may need to check your local IP (via ifconfig).

### Required environment variables

* BASE_URL: the base URL for the openexchangerates API.
* APP_ID: the ID for the calling app for openexchangerates API.
* REDIS_ADDRESS: self explanatory.
* REDIS_PORT: self explanatory.
