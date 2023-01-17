# Auckland Council API (unofficial)

## Docker image

This fork of the orignial repository by [rusq](https://github.com/rusq/aklapi) provides all the files necessary to self-host the API in a docker container. It's as simple as creating the following `docker-compose.yml`:

```yaml
version: '3.4'

services:
  aklapi:
    image: histefanhere/aklapi
    restart: unless-stopped
    environment:
     - NO_RUBBISH_CACHE= #optional
    ports:
      - 8080:8080
```

And running the command:

```
docker-compose up -d
```

Alternatively, you can just run the following command:

```
docker run -p 8080:8080 histefanhere/aklapi
```

Now just navigate to `http://<your server ip>:8080` and you should see the API up and running! 🎉 Refer to the next section for how to use the API and its features.

## API specification

Full list of available endpoints, for detailed description see below.

| Name | Endpoint | Parameters | Comments |
|:----|:----|:----|:-----|
|Address|`/api/v1/addr`|`addr`: partial address|Address Query|
|Rubbish and Recycling short|`/api/v1/rr`|`addr`: partial address|Rubbish and Recycling, short format|
|Rubbish and Recycling|`/api/v1/rrext`|`addr`: partial address|Rubbish and Recycling|

### Address search

* `/api/v1/addr`, parameter: `addr` - 

### Rubbish and Recycling

Two endpoints so far, both accepting `addr` parameter.

* `/api/v1/rr/` - rubbish and recycling, returns the JSON of the following format:

      {
          "rubbish": "2020-02-25",
          "recycle": "2020-02-25",
          "address": "Britomart, CBD"
      }

* `/api/v1/rrext/` - extended rubbish and recycling.  Returns the JSON in the following format:

      {
          "Collections": [
              {
                  "Day": "Monday 24 January",
                  "Date": "2020-01-24T00:00:00+13:00",
                  "Rubbish": true,
                  "Recycle": true
              },
              {
                  "Day": "Monday 31 January",
                  "Date": "2020-01-31T00:00:00+13:00",
                  "Rubbish": true,
                  "Recycle": false
              }
          ],
          "Address": {
              "ACRateAccountKey": "12342478585",
              "Address": "500 Queen Street, Auckland Central",
              "Suggestion": "500 Queen Street, Auckland Central"
          }
      }

Example:

```sh
$ curl --location --request GET 'https://<server>/api/v1/rr/?addr=500%20Queen%20Street'
{"rubbish":"2020-02-24","recycle":"2020-02-24","address":"500 Queen Street, Auckland Central"}
```

### Integrating with Home Assistant

Add the following to your `configuration.yaml`:

```yaml
sensor:
  - platform: rest
    resource: https://your_server/api/v1/rr/?addr=500%20Mystreet
    name: Recycle
    value_template: '{{ value_json.recycle }}'
```