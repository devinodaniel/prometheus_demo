## Install Docker

Install or update Docker.

```bash
docker -v
20.10.8, build 3967b7d
```

## Enable Docker Metrics 

Enable the docker daemon as a Prometheus target. This allows Promethues to use service discovery to find srvices to monitor.

- **Docker for Mac** - Click the Docker icon in the toolbar, select **Preferences**, then select **Daemon**. Click Advanced.
- **Linux** - edit `/etc/docker/daemon.json`

If the file is blank, paste the following:

```json
{
  "metrics-addr" : "0.0.0.0:9323",
  "experimental" : true
}
```

## Start webservers

You'll need something to monitor. Let's start some simple webservers.

```bash
docker compose --profile web up
```

## Add Slack webhook to `alertmanager.yml`

Under `global`.`slack_api_url` add a valid slack webhook URL.

Under `slack_configs`.`channel` update the channel name to be a valid channel.

If you'd lke some fun reading on why Alertmanager doesn't support env vars in config files and some alternatives on how to manage secrets (like Slack webhooks) check out this [Alertmanager issue](https://github.com/prometheus/alertmanager/issues/504).

