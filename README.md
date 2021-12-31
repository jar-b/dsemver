# dsemver

semver in a scratch docker container

## Docker Usage

```sh
docker run jar-b/semver 0.1.1
# 0.1.2
docker run jar-b/semver -i minor 0.1.1
# 0.2.0
docker run jar-b/semver -i major 0.1.1
# 1.0.0
```

## Motivation

Mimicing the version bumping logic from the [alpine/semver](https://hub.docker.com/r/alpine/semver) image (built on `node:alpine`), but in a scratch container with Go.

## References

* [alpine/semver](https://hub.docker.com/r/alpine/semver)
* [semver](https://github.com/blang/semver) (Go implementation)

