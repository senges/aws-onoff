# onoff

Start / Stop AWS EC2 instances.

## Howto

```
# Start single instance
onoff start <name>

# Stop single instance
onoff stop <name>

# Start all instances
onoff start all

# Stop all instances
onoff stop all
```

Instance `<name>` corresponds to `[Instance.<name>]` in config file.

## Build

If you trust binaries (meh), last version of the tool is served in `./releases` folder for both linux and windows x64.

Otherwise, take a look at [BUILD.md](./BUIDL.md) file.
