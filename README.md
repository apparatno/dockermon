# Docker monitor

Monitors Docker containers and logs status to Firebase.

## Compiling and testing

To compile for your local system, run `make`.

To build binaries for distribution, run `make build`.
Currently it will build 64-bit binaries for 
Linux, macOS, and ARM (Raspberry Pi).

To test, run `make test`.

To clean up binaries, run `make clean`.

## Running

Create a service account and keep the JSON file on disk somewhere.
Set `GOOGLE_APPLICATION_CREDENTIALS` to the path of the file.

`dockermon` will write to a collection called `monitoring`.

Run `dockermon` to check all running containers on the host and
write the result to Firebase.

Use the `-dry-run` flag to avoid writing to Firebase.
The service account described above is not necessary when using this flag.
