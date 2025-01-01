# Key features of this handler

- Receives JSON data from STDIN (Consul sends this automatically)

- Logs changes to /var/log/consul-watch.log

- Includes key name, modification index, and value in the logs

- Handles JSON parsing errors gracefully

- The main differences from a typical Go program are:

- It reads from STDIN instead of making API calls

- It processes a single batch of changes and exits (Consul will run it again when new changes occur)

- It's designed to be executed by Consul's watch system rather than running continuously
