iot-cloud-cli is a command line interface that allows to exploit the features of Arduino IoT Cloud. As of now, it is possible to provision a device.

### Requirements

This is all you need to use iot-cloud-cli for device **provisioning**:
 * A client ID and a secret ID, retrievable from the [cloud](https://create.arduino.cc/iot/integrations) by creating a new API key
 * The folder containing the precompiled provisioning firmwares (`binaries`) needs to be in the same location you run the command from

## Set a configuration

iot-cloud-cli needs to be configured before being used. In particular a client ID and the corresponding secret ID should be set.
You can retrieve them from the [cloud](https://create.arduino.cc/iot/integrations) by creating a new API key.

Once you have the IDs, call this command with your parameters:

`$ iot-cloud-cli config -c <clientID> -s <secretID>`

A file named `config.yaml` will be created in the Current Working Directory containing the login credentials.
Example

```yaml
client: 00112233445566778899aabbccddeeff
secret: 00112233445566778899aabbccddeeffffeeddccbbaa99887766554433221100
```

## Device provisioning

When provisioning a device, you can optionally specify the port to which the device is connected to and its fqbn. If they are not given, then the first device found will be provisioned.

Use this command to provision a device:

`$ iot-cloud-cli device create --name <deviceName> --port <port> --fqbn <deviceFqbn>`