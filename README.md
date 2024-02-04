# chippotto
Yet another chip8 emulator (eight is _otto_ in :it:) to experiment with Bazel and start learning Go.

## build and run
1. You need to have [Bazel](https://bazel.build/install) installed. There are multiple ways, if you're on linux one of the simplest is to download [bazelisk](https://github.com/bazelbuild/bazelisk/releases), rename it to bazel, make it executable and move it somewhere in `$PATH`.
```bash
# example for linux on amd64
curl -LO https://github.com/bazelbuild/bazelisk/releases/latest/download/bazelisk-linux-amd64
mv bazelisk-linux-amd64 bazel
chmod +x bazel
sudo mv bazelisk /usr/local/bin/
```
2. Use bazel to build/run the application:
* Use `bazel build //chippotto` to build the application.
* Use `bazel run //chippotto -- <arguments>` to run the program with some arguments. E.g.
```bash
bazel run //chippotto -- --help
```
