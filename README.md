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

| Name                   | Endpoint              | Parameters              | Comments                                        |
|:-----------------------|:----------------------|:------------------------|:------------------------------------------------|
| Address search         | `/api/v2/address`     | `addr`: partial address | Lookup full address and address key             |
| Waste collection dates | `/api/v2/dates`       | `addr`: partial address | Next dates for each waste type collection       |
| All collections        | `/api/v2/collections` | `addr`: partial address | More detailed information about collection days |

### Address search

* `/api/v2/address` endpoint for looking up the full address of a partial address. Also returns the address key used by Auckland Council. parameter: `addr` - a partial address. Returns JSON of the following format:

    [
        {
            "ACRateAccountKey": "12344314126",
            "Address": "14 Glenfell Place, Epsom",
            "Suggestion": "14 Glenfell Place, Epsom"
        }
    ]

### Collection Days

A collection day is any day when any waste type is collected from an address. Different waste types can be collected on different days. There are two endpoints for accessing information on collections days, both accepting the partial address `addr` parameter.

* `/api/v2/dates/` - Returns the date that each waste type will be collected of rubbish, recycling, and food scraps. returns JSON of the following format:

    {
        "rubbish": "2023-05-09",
        "recycle": "2023-05-09",
        "foodscraps": "0001-01-01",
        "address": "14 Glenfell Place, Epsom"
    }

* `/api/v2/collections/` - Provides more detailed information of the collection days and which waste types will be collected on which day. Also provides the address information from the `address` endpoint. Returns JSON in the following format:

    {
        "Collections": [
            {
                "Day": "Tuesday 9 May",
                "Date": "2023-05-09T00:00:00+12:00",
                "Rubbish": true,
                "Recycle": true,
                "FoodScraps": false
            },
            {
                "Day": "Tuesday 16 May",
                "Date": "2023-05-16T00:00:00+12:00",
                "Rubbish": true,
                "Recycle": false,
                "FoodScraps": false
            }
        ],
        "Address": {
            "ACRateAccountKey": "12344314126",
            "Address": "14 Glenfell Place, Epsom",
            "Suggestion": "14 Glenfell Place, Epsom"
        }
    }

## Usage

Example:

```sh
$ curl --location --request GET 'https://<server>/api/v2/dates/?addr=14%20Glenfell%20Place'
{"rubbish":"2023-05-09","recycle":"2023-05-09","foodscraps":"0001-01-01","address":"14 Glenfell Place, Epsom"}
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
