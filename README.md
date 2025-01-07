# Goat

Remote deployment from locally. Easy to config and extends.

### Usage

Create a file `goat.yaml` in root folder. Here is example for `goat.yaml`.

```yaml
servers:
  production:
    user: __USER__
    host: __HOST__
    port: __PORT__
    password: __PASSWORD__
    exec:
      steps:
        - name: docker pull storyui/web:latest
          run: docker pull registry.ninety4.tech/storyui/web:latest
          delay: 2000
        - name: docker pull storyui/api:latest
          run: docker pull registry.ninety4.tech/storyui/api:latest
```

### Exec

```sh
goat exec production
```