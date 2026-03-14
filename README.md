# zapret-tray
простая утилита для включения/выключения zapret одной кнопкой в системном трее.

### build
```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o zapret-tray
```
### usage
`.\zapret-tray` при установке, далее можно закинуть в автозапуск через контекстное меню
### credits
https://github.com/energye/systray
https://github.com/Snowy-Fluffy/zapret.installer


`прогу я делал чисто для себя, поэтому её работа тестировалась только на арче`
