# Next-gen Kurento package for Go

[Kurento](https://kurento.openvidu.io/) is a WebRTC Media Server built on Gstreamer.

This package provides bindings to control Kurento over a WebSocket using Go.

## How this project works

Kurento uses various `MediaObject` types to control the routing and processing of media.
It uses [JSON-RPC 2.0](https://www.jsonrpc.org/specification) over a WebSocket to control the creation and manipulation of these objects.

> [Kurento Protocol Documentation](https://doc-kurento.readthedocs.io/en/latest/features/kurento_protocol.html)

Kurento publishes JSON files to be used with the [Kurento Module Creator](https://github.com/Kurento/kurento-module-creator/tree/6.18.0).
This generates Kurento's official Java and JavaScript clients.
The tool in the `build` folder of this repository uses the same files to generate the Go client in this repository.

## How to build a new version of this library

1. Make sure you have Kurento's `kms-core` and `kms-elements` projects checked out inside the `build` directory.
1. Run `make` - this will:
   1. Clean the output directories (`core` and `elements`) and some intermediate directories inside the `build` directory.
   1. Preprocess the `*.kmd.json` files to remove the non-spec newline characters inside strings
   1. Build the Go client

## Acknowledgements

- Thanks to the Kurento team for creating a way to auto-generate clients in other programming languages!
- Thanks goes to [@metal3d](https://github.com/metal3d) (Patrice Ferlet) for his original implementation of the Go client generator.
