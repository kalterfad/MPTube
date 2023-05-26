```shell
sudo apt install libnotify-bin
```

```shell
go build -o mptube main.go
```

```shell
sudo mv mptube /usr/local/bin/
```

```shell
sudo mkdir /usr/share/icons/MPTube/
```

```shell
sudo cp icon.svg /usr/share/icons/MPTube/
```

---
```shell
sudo cp mptube.desktop /usr/share/applications/
```

### Или

Но тут нужно будет создать файл env.txt и скопировать туда все переменные окружения, иначе MPV не работает.

```shell
sudo cp mptube.service /etc/systemd/system/
```

```shell
sudo systemctl daemon-reload
```

```shell
sudo systemctl enable mptube.service
```

```shell
sudo systemctl start mptube.service
```