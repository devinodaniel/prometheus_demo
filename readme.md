## Install Docker

Install or update Docker.

```bash
docker -v
20.10.8, build 3967b7d
```

## Enable Docker Metrics

Enable the docker daemon as a Prometheus target. This allows Prometheus to use service discovery to find containers to monitor.

- **Docker for Mac** - Click the Docker icon in the toolbar, select **Preferences**, then select **Daemon**. Click Advanced.
- **Linux** - edit `/etc/docker/daemon.json`

Add the following:

> Note: You may just need to switch `experimental` from `false` -> `true` and add `"metrics-addr" : "0.0.0.0:9323"` as it's own line. Docker configurations vary per user.
 
```json
{
  "metrics-addr" : "0.0.0.0:9323",
  "experimental" : true
}
```

## Add Slack webhook

Create a file called `slack_api_url.secret` in the root of this repo. (Files with `.secret` extensions are .gitignore'd so that you do not accidentally commit them). Add the URL of your Slack webhook to this file.

*slack_api_url.secret:*

```bash
https://hooks.slack.com/services/TXXXXXXXM/BXXXXXXXXN/CXXXXXXYYYYYYZZZZZZZT
```

## Launch Prometheus

```bash
docker compose --profile monitor up -d
```

After a few seconds the `WebserverDown`  and `GrafanaDown` alert will fire. You should receive a Slack notification.

## Start the webserver

```bash
docker compose --profile web up -d
```

After a few seconds the `WebserverDown` alert will resolve.

## Start Grafana

```bash
docker compose --profile grafana up -d
```

After a few seconds the `GrafanaDown` alert will resolve.

## Tweak and retry

Tweak the configuration files and alerting rules. Then, rebuild all the containers:

```bash
docker compose --profile demo down
docker compose --profile demo up
```

## Add dashboard to Grafana

Open grafana at `http://localhost:3000`.

Put in `admin` / `admin` as default user and password. You can *Skip* the new password.

From the left sidebar, select *Configuration*, then the blue *Add data source* button.

*Select* the Prometheus data source.

For *URL* enter `prometheus:9090`, leave access as `Server`.

Scroll to the bottom and select *Save & test*. You should see a green success notification.

Click the *+* (plus) sign in the left menu bar to create a dashboard.

Select *Import* from the submenu.

Select *Upload JSON file*, then choose `demo_dash.json` from this repo.

> Important: when you remove the Grafana container you will lose any customizations to dashboards. ALL WORK WILL BE LOST. To keep the Grafana dashboards you create you'll need to mount a volume.

## Clean up

Stop and remove all containers and downloaded images.

```bash
docker compose --profile demo down --rmi all
```
