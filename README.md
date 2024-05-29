# SendPriorityMessageToAssets

This activity sends a priority message to a specified asset.

## Configuration

### Input:

| Name              | Type   | Description                                                                                             |
| :---------------- | :----- | :------------------------------------------------------------------------------------------------------ |
| IP                | string | IP address followed by port. e.g. "32.41.13.112:322"                                                    |
| Username          | string | Username used to log into Sofia. e.g. "adminUser"                                                       |
| Password          | string | Password used to log into Sofia. e.g. "adminPassword"                                                   |
| StaffIdList       | string | Takes a comma delimited list, a single id, or staff object. e.g. "234, 1234, 15" OR "1023", Or StaffObj |
| Message           | string | Message to be sent to asset(s). e.g. "Report to nearest checkpoint"                                     |
| MessageOptions    | string | Choose one of the tones listed for message event                                                        |
| MessageExpiration | int    | Message to be sent to asset(s). e.g. 10 or 0 for no expiration                                          |

### Output:

| Name   | Type   | Description                       |
| :----- | :----- | :-------------------------------- |
| Status | string | True if message sent successfully |
