# Consul Watch Handler

## Key Features

- **JSON Input**: Receives JSON data from STDIN, automatically sent by Consul.
- **Logging**: Logs changes to `/var/log/consul-watch.log`.
- **Detailed Logs**: Includes key name, modification index, and value in the logs.
- **Error Handling**: Gracefully handles JSON parsing errors.

## Differences from Typical Go Programs

- **Input Source**: Reads from STDIN instead of making API calls.
- **Execution Model**: Processes a single batch of changes and exits. Consul will re-execute it when new changes occur.
- **Purpose**: Designed to be executed by Consul's watch system rather than running continuously
